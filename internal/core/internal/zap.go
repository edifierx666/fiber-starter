package internal

import (
  "strings"
  "time"

  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
)

type InternalZap struct {
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

// GetEncoder 获取 zapcore.Encoder
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *InternalZap) GetEncoder() zapcore.Encoder {
  if z.Format == "json" {
    return zapcore.NewJSONEncoder(z.GetEncoderConfig())
  }
  return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

// GetEncoderConfig 获取zapcore.EncoderConfig
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *InternalZap) GetEncoderConfig() zapcore.EncoderConfig {
  return zapcore.EncoderConfig{
    MessageKey:     "message",
    LevelKey:       "level",
    TimeKey:        "time",
    NameKey:        "logger",
    CallerKey:      "caller",
    StacktraceKey:  z.StacktraceKey,
    LineEnding:     zapcore.DefaultLineEnding,
    EncodeLevel:    z.ZapEncodeLevel(),
    EncodeTime:     z.CustomTimeEncoder,
    EncodeDuration: zapcore.SecondsDurationEncoder,
    EncodeCaller:   zapcore.FullCallerEncoder,
  }
}

// GetEncoderCore 获取Encoder的 zapcore.Core
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *InternalZap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
  writer := FileRotatelogs.GetWriteSyncer(l.String()) // 日志分割
  return zapcore.NewCore(z.GetEncoder(), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *InternalZap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
  encoder.AppendString(z.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

// GetZapCores 根据配置文件的Level获取 []zapcore.Core
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *InternalZap) GetZapCores() []zapcore.Core {
  cores := make([]zapcore.Core, 0, 7)
  for level := z.TransportLevel(); level <= zapcore.FatalLevel; level++ {
    cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
  }
  return cores
}

// ZapEncodeLevel 根据 EncodeLevel 返回 zapcore.LevelEncoder
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *InternalZap) ZapEncodeLevel() zapcore.LevelEncoder {
  switch {
  case z.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
    return zapcore.LowercaseLevelEncoder
  case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
    return zapcore.LowercaseColorLevelEncoder
  case z.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
    return zapcore.CapitalLevelEncoder
  case z.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
    return zapcore.CapitalColorLevelEncoder
  default:
    return zapcore.LowercaseLevelEncoder
  }
}

// TransportLevel 根据字符串转化为 zapcore.Level
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *InternalZap) TransportLevel() zapcore.Level {
  z.Level = strings.ToLower(z.Level)
  switch z.Level {
  case "debug":
    return zapcore.DebugLevel
  case "info":
    return zapcore.InfoLevel
  case "warn":
    return zapcore.WarnLevel
  case "error":
    return zapcore.WarnLevel
  case "dpanic":
    return zapcore.DPanicLevel
  case "panic":
    return zapcore.PanicLevel
  case "fatal":
    return zapcore.FatalLevel
  default:
    return zapcore.DebugLevel
  }
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *InternalZap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
  switch level {
  case zapcore.DebugLevel:
    return func(level zapcore.Level) bool { // 调试级别
      return level == zap.DebugLevel
    }
  case zapcore.InfoLevel:
    return func(level zapcore.Level) bool { // 日志级别
      return level == zap.InfoLevel
    }
  case zapcore.WarnLevel:
    return func(level zapcore.Level) bool { // 警告级别
      return level == zap.WarnLevel
    }
  case zapcore.ErrorLevel:
    return func(level zapcore.Level) bool { // 错误级别
      return level == zap.ErrorLevel
    }
  case zapcore.DPanicLevel:
    return func(level zapcore.Level) bool { // dpanic级别
      return level == zap.DPanicLevel
    }
  case zapcore.PanicLevel:
    return func(level zapcore.Level) bool { // panic级别
      return level == zap.PanicLevel
    }
  case zapcore.FatalLevel:
    return func(level zapcore.Level) bool { // 终止级别
      return level == zap.FatalLevel
    }
  default:
    return func(level zapcore.Level) bool { // 调试级别
      return level == zap.DebugLevel
    }
  }
}
