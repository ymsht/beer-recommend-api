package model

import (
	"gopkg.in/gorp.v1"
)

/**
 * Area 地域情報
 */
type Area struct {
	AreaID    int    `db:"area_id" json:"area_id"`
	CountryID int    `db:"country_id" json:"country_id"`
	AreaName  string `db:"area_name" json:"area_name"`
}

// GetAreas 地域情報を取得します
func GetAreas(tx *gorp.Transaction) ([]Area, error) {

	a, err := selectToAreas(tx)
	if err != nil {
		return a, err
	}

	return a, nil
}

// selectToAreas 地域情報を検索します
func selectToAreas(tx *gorp.Transaction) ([]Area, error) {
	var a []Area
	_, err := tx.Select(&a, `
		select
		  *
		from
		  area
		order by
		area_id
	`)
	if err != nil {
		return a, err
	}

	return a, nil
}
