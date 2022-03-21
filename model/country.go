package model

import (
	"gopkg.in/gorp.v1"
)

/**
 * Country 原産国情報
 */
type Country struct {
	CountryID   int    `db:"country_id" json:"country_id"`
	CountryName string `db:"country_name" json:"country_name"`
}

// GetCountries 原産国情報を取得します
func GetCountries(tx *gorp.Transaction) ([]Country, error) {

	c, err := selectToCountries(tx)
	if err != nil {
		return c, err
	}

	return c, nil
}

// selectToCountries 原産国情報を検索します
func selectToCountries(tx *gorp.Transaction) ([]Country, error) {
	var c []Country
	_, err := tx.Select(&c, `
		select
		  *
		from
		  country
		order by
		country_id
	`)
	if err != nil {
		return c, err
	}

	return c, nil
}
