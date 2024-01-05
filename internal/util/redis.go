package util

import (
  "context"

  "ats3/internal/config"
  "ats3/internal/g"
  "github.com/flipped-aurora/gin-vue-admin/server/global"
  "github.com/redis/go-redis/v9"
  "go.uber.org/zap"
)

func Redis(redisCfg *config.Redis) {
  client := redis.NewClient(&redis.Options{
    Addr:     redisCfg.Addr,
    Password: redisCfg.Password, // no password set
    DB:       redisCfg.DB,       // use default DB
  })
  pong, err := client.Ping(context.Background()).Result()
  if err != nil {
    g.Log.Error("redis connect ping failed, err:", zap.Error(err))
  } else {
    g.Log.Info("redis connect ping response:", zap.String("pong", pong))
    global.GVA_REDIS = client
  }
}
