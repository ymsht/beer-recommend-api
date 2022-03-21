package model

import (
	"gopkg.in/gorp.v1"
)

/**
 * Style スタイル情報
 */
type Style struct {
	StyleID   int    `db:"style_id" json:"style_id"`
	StyleName string `db:"style_name" json:"style_name"`
}

// GetStyles スタイル情報を取得します
func GetStyles(tx *gorp.Transaction) ([]Style, error) {

	s, err := selectToStyles(tx)
	if err != nil {
		return s, err
	}

	return s, nil
}

// selectToStyles スタイル情報を検索します
func selectToStyles(tx *gorp.Transaction) ([]Style, error) {
	var s []Style
	_, err := tx.Select(&s, `
		select
		  *
		from
		  style
		order by
		  style_id
	`)
	if err != nil {
		return s, err
	}

	return s, nil
}
