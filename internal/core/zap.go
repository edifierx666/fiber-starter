package core

import (
  "fmt"
  "os"

  "ats3/internal/core/internal"
  "ats3/internal/g"
  utils "ats3/internal/util"
  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *zap.Logger) {
  if ok, _ := utils.PathExists(g.CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
    fmt.Printf("create %v directory\n", g.CONFIG.Zap.Director)
    _ = os.Mkdir(g.CONFIG.Zap.Director, os.ModePerm)
  }

  cores := internal.Zap.GetZapCores()
  logger = zap.New(zapcore.NewTee(cores...))

  if g.CONFIG.Zap.ShowLine {
    logger = logger.WithOptions(zap.AddCaller())
  }
  return logger
}
