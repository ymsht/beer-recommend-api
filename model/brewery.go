package model

import (
	"gopkg.in/gorp.v1"
)

/**
 * Brewery
 */
type Brewery struct {
	BreweryId   int    `db:"brewery_id" json:"brewery_id"`
	BreweryName string `db:"brewery_name" json:"brewery_name"`
}

// GetBreweries
func GetBreweries(tx *gorp.Transaction) ([]Brewery, error) {

	s, err := selectToBreweries(tx)
	if err != nil {
		return s, err
	}

	return s, nil
}

// selectToBreweries
func selectToBreweries(tx *gorp.Transaction) ([]Brewery, error) {
	var s []Brewery
	_, err := tx.Select(&s, `
		select
		brewery_id,
			brewery_name
		from
		  brewery
		order by
		  brewery_id
	`)
	if err != nil {
		return s, err
	}

	return s, nil
}
