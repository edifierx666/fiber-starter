package middleware

import (
  "encoding/json"
  "fmt"
  "time"

  "ats3fx/internal/g"
  "github.com/gofiber/fiber/v2"
)

// LogLayout 日志layout
type LogLayout struct {
  Time      time.Time
  Method    string
  Metadata  map[string]interface{} // 存储自定义原数据
  Path      string                 // 访问路径
  Query     string                 // 携带query
  Body      string                 // 携带body数据
  IP        string                 // ip地址
  UserAgent string                 // 代理
  Error     string                 // 错误
  Cost      time.Duration          // 花费时间
  Source    string                 // 来源
}

type Logger struct {
  // Filter 用户自定义过滤
  Filter func(c *fiber.Ctx) bool
  // FilterKeyword 关键字过滤(key)
  FilterKeyword func(layout *LogLayout) bool
  // AuthProcess 鉴权处理
  AuthProcess func(c *fiber.Ctx, layout *LogLayout)
  // 日志处理
  Print func(LogLayout)
  // Source 服务唯一标识
  Source string
}

func (l Logger) SetLoggerMiddleware() fiber.Handler {
  return func(c *fiber.Ctx) error {
    start := time.Now()
    path := string(c.Request().URI().Path())
    query := string(c.Request().URI().QueryString())
    var body []byte
    if l.Filter != nil && !l.Filter(c) {
      // body, _ = c.GetRawData()
      // 将原body塞回去
      // c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
    }
    err := c.Next()
    cost := time.Since(start)
    layout := LogLayout{
      Time:      time.Now(),
      Method:    c.Method(),
      Path:      path,
      Query:     query,
      IP:        c.IP(),
      UserAgent: string(c.Request().Header.UserAgent()),
      Error:     "",
      Cost:      cost,
      Source:    l.Source,
    }
    if l.Filter != nil && !l.Filter(c) {
      layout.Body = string(body)
    }
    if l.AuthProcess != nil {
      // 处理鉴权需要的信息
      l.AuthProcess(c, &layout)
    }
    if l.FilterKeyword != nil {
      // 自行判断key/value 脱敏等
      l.FilterKeyword(&layout)
    }
    // 自行处理日志
    l.Print(layout)
    return err
  }
}

func DefaultLogger() fiber.Handler {
  return Logger{
    Print: func(layout LogLayout) {
      // 标准输出,k8s做收集
      v, _ := json.Marshal(layout)
      fmt.Println(string(v))
    },
    Source: g.CONFIG.System.AppName,
  }.SetLoggerMiddleware()
}
