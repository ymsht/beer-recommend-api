package model

import (
	"time"

	"gopkg.in/gorp.v1"
)

type Login struct {
	UserId      int       `db:"user_id" json:"user_id"`
	UserName    string    `db:"user_name" json:"user_name"`
	Password    string    `db:"password" json:"password"`
	Create_date time.Time `db:"create_date" json:"create_date"`
	Update_date time.Time `db:"update_date" json:"update_date"`
}

func GethUser(tx *gorp.Transaction, n string) (Login, error) {
	u, err := searchUser(tx, n)
	if err != nil {
		return u, err
	}

	return u, nil
}

func searchUser(tx *gorp.Transaction, n string) (Login, error) {
	var u Login
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
