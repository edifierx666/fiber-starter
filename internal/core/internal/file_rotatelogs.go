package internal

import (
  "os"

  "ats3fx/internal/g"
  "go.uber.org/zap/zapcore"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
// Author [SliverHorn](https://github.com/SliverHorn)
func (r *fileRotatelogs) GetWriteSyncer(level string) zapcore.WriteSyncer {
  fileWriter := NewCutter(g.CONFIG.Zap.Director, level, WithCutterFormat("2006-01-02"))
  if g.CONFIG.Zap.LogInConsole {
    return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
  }

  return zapcore.AddSync(fileWriter)
}
