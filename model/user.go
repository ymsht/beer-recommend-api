package model

import (
	"time"

	"gopkg.in/gorp.v1"
)

type User struct {
	UserId      int       `db:"user_id" json:"user_id"`
	UserName    string    `db:"user_name" json:"user_name"`
	Password    string    `db:"password" json:"password"`
	Create_date time.Time `db:"create_date" json:"create_date"`
	Update_date time.Time `db:"update_date" json:"update_date"`
}

// GethUser ユーザ情報を返します
func GethUser(tx *gorp.Transaction, n string) (User, error) {
	u, err := searchUser(tx, n)
	if err != nil {
		return u, err
	}

	return u, nil
}

// searchUser ユーザを検索します
func searchUser(tx *gorp.Transaction, n string) (User, error) {
	var u User
	err := tx.SelectOne(&u, `
		select
		  *
		from
		  user
		where
		  user_name = ?
	`, n)
	if err != nil {
		return u, err
	}

	return u, nil
}

// CreateUser ユーザを登録します
func CreateUser(tx *gorp.Transaction, u User) error {
	err := tx.Insert(&u)
	if err != nil {
		return err
	}

	return nil
}
