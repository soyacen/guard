package service

import (
	"testing"

	"github.com/yacen/guard/db"
)

func TestMain(m *testing.M) {
	db.InitMysql()
	db.InitRedis()
	m.Run()
}

func TestSignIn(t *testing.T) {
	token, err := SignIn("jax", "123456")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
