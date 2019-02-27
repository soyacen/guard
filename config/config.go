package config

import (
	"github.com/spf13/viper"
	"github.com/yacen/guard/util/log"
)

type Config struct {
	CfgFile  string
	Port     int
	Https    bool
	CertFile string
	KeyFile  string

	UsernamePattern string
	PasswordPattern string

	DBDriver string
	DNS      string

	RedisAddr string
	RedisPwd  string
	RedisDB   int

	JwtKeyFile string
}

var Cfg Config

func GetFlagsFromConfigFile() {
	Cfg.Port = viper.GetInt("port")
	Cfg.Https = viper.GetBool("https")
	Cfg.CertFile = viper.GetString("tls-cert-file")
	Cfg.KeyFile = viper.GetString("tls-private-key-file")
	Cfg.UsernamePattern = viper.GetString("username-pattern")
	Cfg.PasswordPattern = viper.GetString("password-pattern")
	Cfg.DNS = viper.GetString("dns")
	Cfg.DBDriver = viper.GetString("driver")
	Cfg.RedisAddr = viper.GetString("redis")
	Cfg.RedisPwd = viper.GetString("redis-pwd")
	Cfg.RedisDB = viper.GetInt("redis-db")
	Cfg.JwtKeyFile = viper.GetString("jwt-key-file")
	log.Println(Cfg)
}
