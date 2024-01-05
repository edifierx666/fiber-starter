package util

import (
  "ats3/internal/config"
  _ "github.com/go-sql-driver/mysql"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func GormMysql(m *config.Mysql) *gorm.DB {
  if m.Dbname == "" {
    return nil
  }
  mysqlConfig := mysql.Config{
    DSN:                       m.Dsn(), // DSN data source name
    DefaultStringSize:         191,     // string 类型字段的默认长度
    SkipInitializeWithVersion: false,   // 根据版本自动配置
  }
  if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
    return nil
  } else {
    db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
    sqlDB, _ := db.DB()
    sqlDB.SetMaxIdleConns(m.MaxIdleConns)
    sqlDB.SetMaxOpenConns(m.MaxOpenConns)
    return db
  }
}

// GormMysqlByConfig 初始化Mysql数据库用过传入配置
func GormMysqlByConfig(m config.Mysql, option ...gorm.Option) *gorm.DB {
  if m.Dbname == "" {
    return nil
  }
  mysqlConfig := mysql.Config{
    DSN:                       m.Dsn(), // DSN data source name
    DefaultStringSize:         191,     // string 类型字段的默认长度
    SkipInitializeWithVersion: false,   // 根据版本自动配置
  }
  if db, err := gorm.Open(mysql.New(mysqlConfig), option...); err != nil {
    panic(err)
  } else {
    db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
    sqlDB, _ := db.DB()
    sqlDB.SetMaxIdleConns(m.MaxIdleConns)
    sqlDB.SetMaxOpenConns(m.MaxOpenConns)
    return db
  }
}
