package service

import (
	"crypto/sha1"
	"encoding/hex"
	"time"

	"github.com/yacen/guard/config"

	"github.com/yacen/guard/util"

	"github.com/dgrijalva/jwt-go"

	"github.com/yacen/guard/model"

	"github.com/yacen/guard/db"
	"github.com/yacen/guard/resp"
	"github.com/yacen/guard/util/log"
	"golang.org/x/crypto/pbkdf2"
)

func SignIn(username, password string) (token string, err error) {
	// redis 里是否存在
	exists, err := db.Redis.SIsMember(db.KeyUsernameSet, username).Result()
	if err != nil {
		log.Error(err)
		err = resp.NewError(resp.SignInError, "sign in failed")
		return
	}
	// 不存在就返回错误
	if !exists {
		log.Error("username not exists in redis")
		err = resp.NewError(resp.SignInError, "please sign up")
		return
	}

	// 查询数据库
	accounts, err := model.FindAccountsByUsername(username)
	if err != nil {
		log.Error(err)
		err = resp.NewError(resp.SignInError, "sign in failed")
		return
	}
	if len(accounts) <= 0 {
		log.Error("username not exists in database")
		err = resp.NewError(resp.SignInError, "please sign up")
		return
	} else if len(accounts) > 1 {
		log.Error("There are multiple usernames")
		err = resp.NewError(resp.SignInError, "sign in failed")
		return
	}
	account := accounts[0]

	// 加密密码
	salt := []byte(account.Salt)
	sePwd := hex.EncodeToString(pbkdf2.Key([]byte(password), salt, 4096, 32, sha1.New))

	if sePwd != account.Password {
		log.Error("password wrong")
		err = resp.NewError(resp.SignInError, "sign in failed")
		return
	}

	type UserClaims struct {
		User  string `json:"user"`
		Phone string `json:"phone"`
		Email string `json:"email"`
		jwt.StandardClaims
	}

	now := time.Now()
	expiresIn := time.Millisecond * config.Cfg.TokenExpiresIn
	claims := UserClaims{
		User:  account.Username,
		Phone: account.Phone,
		Email: account.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(expiresIn).Unix(),
			IssuedAt:  now.Unix(),
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodRS512, claims).SignedString(util.RsaPriv)
	if err != nil {
		log.Error("generate token error,", err)
		err = resp.NewError(resp.SignInError, "sign in failed")
		return
	}
	/*	// 存入redis
		_, err = db.Redis.Set(db.KeyUsernameToken(username), token, expiresIn).Result()
		if err != nil {
			log.Error("save token on redis failed,", err)
			err = resp.NewError(resp.SignInError, "sign in failed")
			return
		}*/
	return
}
