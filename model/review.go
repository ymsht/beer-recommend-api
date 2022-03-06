package model

import (
	"database/sql"
	"time"

	"gopkg.in/gorp.v1"
)

// Review レビュー情報
type Review struct {
	ReviewID    int            `db:"review_id" json:"reviewId"`
	MemberId    int            `db:"member_id" json:"memberId"`
	DrinkingDay time.Time      `db:"drinking_day" json:"drinkingDay"`
	isPublic    bool           `db:"isPublic" json:"isPublic"`
	brewery     sql.NullString `db:"brewery" json:"brewery"`
	beer_name   string         `db:"beer_name" json:"beer_name"`
	store       sql.NullString `db:"store" json:"store"`
	bar         sql.NullString `db:"bar" json:"bar"`
	aroma       int            `db:"aroma" json:"aroma"`
	bitterTaste int            `db:"bitterTaste" json:"bitterTaste"`
	sweetTaste  int            `db:"sweetTaste" json:"sweetTaste"`
	body        int            `db:"body" json:"body"`
	sharpness   int            `db:"sharpness" json:"sharpness"`
	country_id  int            `db:"country_id" json:"country_id"`
	memo        sql.NullString `db:"memo" json:"memo"`
	create_date time.Time      `db:"create_date" json:"create_date"`
	update_date time.Time      `db:"update_date" json:"update_date"`
}

// GetReviews レビュー情報を取得します
func GetReviews(tx *gorp.Transaction) ([]Along, error) {

	alongs, err := selectToReviews(tx)
	if err != nil {
		return alongs, err
	}

	return alongs, nil
}

// selectToReviews レビュー情報を検索します
func selectToReviews(tx *gorp.Transaction) ([]Along, error) {
	var alongs []Along
	_, err := tx.Select(&alongs, `
		select
		  *
		from
		  review
		order by
		  review_id
	`)
	if err != nil {
		return alongs, err
	}

	return alongs, nil
}
