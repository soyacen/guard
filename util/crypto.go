package util

import (
	"crypto/rsa"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/yacen/guard/util/log"

	"github.com/yacen/guard/config"
)

var RsaPriv *rsa.PrivateKey
var RsaPub *rsa.PublicKey

func InitJwtKeyFile() {
	Der, err := ioutil.ReadFile(config.Cfg.JwtPrivateKeyFile)
	if err != nil {
		log.Fatal(err)
	}
	RsaPriv, err = jwt.ParseRSAPrivateKeyFromPEM(Der)
	if err != nil {
		log.Fatal(err)
	}

	Der, err = ioutil.ReadFile(config.Cfg.JwtPublickKeyFile)
	if err != nil {
		log.Fatal(err)
	}
	RsaPub, err = jwt.ParseRSAPublicKeyFromPEM(Der)
	if err != nil {
		log.Fatal(err)
	}

}
