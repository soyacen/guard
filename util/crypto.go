package util

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"github.com/yacen/guard/config"
	"github.com/yacen/guard/util/log"
)

var RsaPriv *rsa.PrivateKey

func InitJwtKeyFile() {
	der, err := ioutil.ReadFile(config.Cfg.JwtKeyFile)
	if err != nil {
		log.Fatal(err)
	}
	block, _ := pem.Decode(der)
	RsaPriv, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
