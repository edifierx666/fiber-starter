package config

type System struct {
  AppName string `mapstructure:"name" json:"name" yaml:"name"`
  Port    string `mapstructure:"port" json:"port" yaml:"port"`
}
