package model

import (
	"gopkg.in/gorp.v1"
)

/**
 * Beer スタイル情報
 */
type Beer struct {
	BeerID   int    `db:"beer_id" json:"beer_id"`
	BeerName string `db:"beer_name" json:"beer_name"`
}

// GetBeers
func GetBeers(tx *gorp.Transaction, id int) ([]Beer, error) {

	s, err := selectToBeers(tx, id)
	if err != nil {
		return s, err
	}

	return s, nil
}

// selectToBeers
func selectToBeers(tx *gorp.Transaction, id int) ([]Beer, error) {
	var s []Beer
	_, err := tx.Select(&s, `
		select
		  beer_id,
			beer_name
		from
		  beer
		where
			brewery_id = ?
		order by
		  beer_id
	`, id)
	if err != nil {
		return s, err
	}

	return s, nil
}
