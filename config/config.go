package config

import (
	"time"

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

	JwtPrivateKeyFile string
	JwtPublickKeyFile string

	TokenExpiresIn time.Duration
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
	Cfg.JwtPrivateKeyFile = viper.GetString("jwt-private-key-file")
	Cfg.JwtPublickKeyFile = viper.GetString("jwt-public-key-file")
	Cfg.TokenExpiresIn = viper.GetDuration("token-expires-in")
	log.Println(Cfg)
}
