package util

import (
  "ats3fx/internal/config"
  "github.com/flipped-aurora/gin-vue-admin/server/global"
  "github.com/glebarez/sqlite"
  "gorm.io/gorm"
)

// GormSqlite 初始化Sqlite数据库
func GormSqlite() *gorm.DB {
  s := global.GVA_CONFIG.Sqlite
  if s.Dbname == "" {
    return nil
  }

  if db, err := gorm.Open(sqlite.Open(s.Dsn())); err != nil {
    panic(err)
  } else {
    sqlDB, _ := db.DB()
    sqlDB.SetMaxIdleConns(s.MaxIdleConns)
    sqlDB.SetMaxOpenConns(s.MaxOpenConns)
    return db
  }
}

// GormSqliteByConfig 初始化Sqlite数据库用过传入配置
func GormSqliteByConfig(s config.Sqlite, option ...gorm.Option) *gorm.DB {
  if s.Dbname == "" {
    return nil
  }

  if db, err := gorm.Open(sqlite.Open(s.Dsn()), option...); err != nil {
    panic(err)
  } else {
    sqlDB, _ := db.DB()
    sqlDB.SetMaxIdleConns(s.MaxIdleConns)
    sqlDB.SetMaxOpenConns(s.MaxOpenConns)
    return db
  }
}
