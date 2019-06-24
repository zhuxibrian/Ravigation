package storage

import (
	Redis "github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
	"Ravigation/ravigation/utils"
	"time"
)

var redisClient *Redis.Pool
func init() {
	config := utils.Config()
	maxIdle := config.GetInt("RedisConfig.MaxIdle")
	maxActive := config.GetInt("RedisConfig.MaxActive")

	// 建立连接池
	redisClient = &Redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(config.GetInt("RedisConfig.MaxIdleTimeout")) * time.Second,
		Wait:        true,
		Dial: func() (Redis.Conn, error) {
			con, err := Redis.Dial("tcp", config.GetString("RedisConfig.Host"),
				Redis.DialPassword(config.GetString("RedisConfig.Password")),
				Redis.DialDatabase(config.GetInt("RedisConfig.Db")),
				Redis.DialConnectTimeout(time.Duration(config.GetInt("RedisConfig.ConnectTimeout"))*time.Second),
				Redis.DialReadTimeout(time.Duration(config.GetInt("RedisConfig.ReadTimeout"))*time.Second),
				Redis.DialWriteTimeout(time.Duration(config.GetInt("RedisConfig.WriteTimeout"))*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}


func GetConn() Redis.Conn {
	// 从池里获取连接
	conn := redisClient.Get()

	// 错误判断
	if conn.Err() != nil {
		log.Error("get connect from pool error: %s", conn.Err())
		return nil
	}

	return conn
}
