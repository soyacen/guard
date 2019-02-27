package model

import (
	"errors"

	"github.com/yacen/guard/db"
)

/**
CREATE TABLE IF NOT EXISTS account (
  id INT AUTO_INCREMENT PRIMARY KEY ,
  username VARCHAR(16) DEFAULT '',
  phone CHAR(15) DEFAULT '',
  email CHAR(24) DEFAULT '',
  password CHAR(64) NOT NULL ,
  salt CHAR(36) NOT NULL
);

*/

type Account struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Salt     string `json:"salt"`
}

func CreateAccount(username, password, salt string) (err error) {
	stmt, err := db.DB.Prepare("INSERT INTO account SET username=?, password=?, salt=?")
	if err != nil {
		return
	}
	result, err := stmt.Exec(username, password, salt)
	if err != nil {
		return
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return
	}
	if rows <= 0 {
		return errors.New("create failed")
	}
	return nil
}

func FindAccountsByUsername(username string) (accounts []*Account, err error) {
	stmt, err := db.DB.Prepare("SELECT id, username, phone, email, password, salt FROM account WHERE username=?")
	if err != nil {
		return
	}
	rows, err := stmt.Query(username)
	if rows.Next() {
		account := new(Account)
		err = rows.Scan(&account.ID, &account.Username, &account.Phone, &account.Email, &account.Password, &account.Salt)
		accounts = append(accounts, account)
	}
	return
}

func (account *Account) FindFromMysql() {

}
