package core

import (
  "fmt"
  "os"

  "ats3fx/internal/core/internal"
  "dario.cat/mergo"
  "github.com/flipped-aurora/gin-vue-admin/server/utils"
  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
)

type Params struct {
  Level         string `mapstructure:"level" json:"level" yaml:"level"`                                                 // 级别
  Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                                              // 日志前缀
  Format        string `mapstructure:"format" json:"format" yaml:"format"`                                              // 输出
  Director      string `mapstructure:"director" json:"director"  yaml:"director" name:"core-module-filerotatelogs-dir"` // 日志文件夹
  EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`                            // 编码级
  StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"`                      // 栈名

  MaxAge       int  `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                      // 日志留存时间
  ShowLine     bool `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
  LogInConsole bool `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
}

func Zap(p *Params) (logger *zap.Logger) {
  if p.Director != "" {
    if ok, _ := utils.PathExists(p.Director); !ok { // 判断是否有Director文件夹
      fmt.Printf("create %v directory\n", p.Director)
      _ = os.Mkdir(p.Director, os.ModePerm)
    }
  }

  internalZap := &internal.InternalZap{}

  _ = mergo.Merge(internalZap, p)
  cores := internalZap.GetZapCores()
  logger = zap.New(zapcore.NewTee(cores...))

  if internalZap.ShowLine {
    logger = logger.WithOptions(zap.AddCaller())
  }
  return logger
}
