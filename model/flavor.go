package model

import (
	"gopkg.in/gorp.v1"
)

/**
 * Flavor フレーバー情報
 */
type Flavor struct {
	FlavorID   int    `db:"flavor_id" json:"flavor_id"`
	FlavorName string `db:"flavor_name" json:"flavor_name"`
}

// GetFlavors フレーバー情報を取得します
func GetFlavors(tx *gorp.Transaction) ([]Flavor, error) {

	f, err := selectToFlavors(tx)
	if err != nil {
		return f, err
	}

	return f, nil
}

// selectToFlavors フレーバー情報を検索します
func selectToFlavors(tx *gorp.Transaction) ([]Flavor, error) {
	var f []Flavor
	_, err := tx.Select(&f, `
		select
		  *
		from
		  flavor
		order by
		  flavor_id
	`)
	if err != nil {
		return f, err
	}

	return f, nil
}
