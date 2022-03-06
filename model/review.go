package model

import (
	"database/sql"
	"time"

	"gopkg.in/gorp.v1"
)

// Review レビュー情報
type Review struct {
	ReviewID    int            `db:"review_id" json:"review_id"`
	MemberId    int            `db:"member_id" json:"member_id"`
	DrinkingDay time.Time      `db:"drinking_day" json:"drinking_day"`
	IsPublic    bool           `db:"is_public" json:"is_public"`
	Brewery     sql.NullString `db:"brewery" json:"brewery"`
	Beer_name   string         `db:"beer_name" json:"beer_name"`
	Store       sql.NullString `db:"store" json:"store"`
	Bar         sql.NullString `db:"bar" json:"bar"`
	Aroma       sql.NullInt64  `db:"aroma" json:"aroma"`
	BitterTaste sql.NullInt64  `db:"bitterTaste" json:"bitterTaste"`
	SweetTaste  sql.NullInt64  `db:"sweetTaste" json:"sweetTaste"`
	Body        sql.NullInt64  `db:"body" json:"body"`
	Sharpness   sql.NullInt64  `db:"sharpness" json:"sharpness"`
	CountryId   sql.NullInt64  `db:"country_id" json:"country_id"`
	Memo        sql.NullString `db:"memo" json:"memo"`
	Create_date time.Time      `db:"create_date" json:"create_date"`
	Update_date time.Time      `db:"update_date" json:"update_date"`
	Evaluation  sql.NullInt64  `db:"evaluation" json:"evaluation"`
	StyleId     sql.NullInt64  `db:"style_id" json:"style_id"`
}

// GetReviews レビュー情報を取得します
func GetReviews(tx *gorp.Transaction) ([]Review, error) {

	reviews, err := selectToReviews(tx)
	if err != nil {
		return reviews, err
	}

	return reviews, nil
}

// selectToReviews レビュー情報を検索します
func selectToReviews(tx *gorp.Transaction) ([]Review, error) {
	var reviews []Review
	_, err := tx.Select(&reviews, `
		select
		  *
		from
		  review
		order by
		  review_id
	`)
	if err != nil {
		return reviews, err
	}

	return reviews, nil
}
