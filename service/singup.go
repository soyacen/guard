package service

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/google/uuid"
	"github.com/yacen/guard/db"
	"github.com/yacen/guard/model"
	"github.com/yacen/guard/resp"
	"github.com/yacen/guard/util/log"
	"golang.org/x/crypto/pbkdf2"
)

func SignUp(username, password string) (err error) {

	// check redis
	exists, err := db.Redis.SIsMember(db.KeyUsernameSet, username).Result()
	if err != nil {
		log.Error(err)
		err = resp.NewError(resp.SignUp, "create account failed")
		return
	}
	if exists {
		log.Error("username exists in redis")
		err = resp.NewError(resp.SignUp, "username already exists")
		return
	}

	// 加密密码
	salt, err := uuid.New().MarshalText()
	if err != nil {
		log.Error(err)
		err = resp.NewError(resp.SignUp, "create account failed")
		return
	}
	sePwd := hex.EncodeToString(pbkdf2.Key([]byte(password), salt, 4096, 32, sha1.New))

	// 添加到数据库
	err = model.CreateAccount(username, sePwd, string(salt))
	if err != nil {
		log.Error(err)
		err = resp.NewError(resp.SignUp, "create account failed")
		return
	}

	// 添加到redis
	_, err = db.Redis.SAdd(db.KeyUsernameSet, username).Result()
	if err != nil {
		log.Error("insert redis failed", err)
	}
	return
}
