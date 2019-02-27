package db

import (
	"database/sql"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yacen/guard/config"
	"github.com/yacen/guard/util/log"
)

var DB *sql.DB

var Redis *redis.Client

func InitMysql() {
	var err error
	log.Printf("database driver is %s, dns is %s", config.Cfg.DBDriver, config.Cfg.DNS)
	DB, err = sql.Open(config.Cfg.DBDriver, config.Cfg.DNS)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	DB.SetMaxIdleConns(1000)
}

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Cfg.RedisAddr,
		Password: config.Cfg.RedisPwd,
		DB:       config.Cfg.RedisDB,
		PoolSize: 100,
	})
	err := Redis.Ping().Err()
	if err != nil {
		log.Fatal(err)
	}
}
