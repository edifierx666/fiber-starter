package config

type Server struct {
  AppName string `mapstructure:"name" json:"name" yaml:"name"`
  Zap     Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
  Redis   Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`

  // gorm
  Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

  Sqlite Sqlite          `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
  DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
}
