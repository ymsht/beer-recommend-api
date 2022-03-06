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
	Evaluation  int            `db:"evaluation" json:"evaluation"`
	Aroma       int            `db:"aroma" json:"aroma"`
	BitterTaste int            `db:"bitterTaste" json:"bitterTaste"`
	SweetTaste  int            `db:"sweetTaste" json:"sweetTaste"`
	Body        int            `db:"body" json:"body"`
	Sharpness   int            `db:"sharpness" json:"sharpness"`
	CountryId   int            `db:"country_id" json:"country_id"`
	Memo        sql.NullString `db:"memo" json:"memo"`
	Create_date time.Time      `db:"create_date" json:"create_date"`
	Update_date time.Time      `db:"update_date" json:"update_date"`
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
