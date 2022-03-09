package model

import (
	"database/sql"
	"encoding/json"
	"time"

	"gopkg.in/gorp.v1"
)

type NullString struct {
	sql.NullString
}

// Review レビュー情報
type Review struct {
	ReviewID     int           `db:"review_id" json:"review_id"`
	MemberId     int           `db:"member_id" json:"member_id"`
	DrinkingDay  sql.NullTime  `db:"drinking_day" json:"drinking_day"`
	IsPublic     bool          `db:"is_public" json:"is_public"`
	Brewery      NullString    `db:"brewery" json:"brewery"`
	Beer_name    string        `db:"beer_name" json:"beer_name"`
	Store        NullString    `db:"store" json:"store"`
	Bar          NullString    `db:"bar" json:"bar"`
	Aroma        sql.NullInt64 `db:"aroma" json:"aroma"`
	BitterTaste  sql.NullInt64 `db:"bitterTaste" json:"bitterTaste"`
	SweetTaste   sql.NullInt64 `db:"sweetTaste" json:"sweetTaste"`
	Body         sql.NullInt64 `db:"body" json:"body"`
	Sharpness    sql.NullInt64 `db:"sharpness" json:"sharpness"`
	CountryId    sql.NullInt64 `db:"country_id" json:"country_id"`
	Memo         NullString    `db:"memo" json:"memo"`
	Create_date  time.Time     `db:"create_date" json:"create_date"`
	Update_date  time.Time     `db:"update_date" json:"update_date"`
	Evaluation   sql.NullInt64 `db:"evaluation" json:"evaluation"`
	StyleId      sql.NullInt64 `db:"style_id" json:"style_id"`
	PurchaseDate sql.NullTime  `db:"purchase_date" json:"purchase_date"`
}

func (s NullString) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	s.String = str
	s.Valid = str != ""
	return nil
}

// GetReviews レビュー情報を取得します
func GetReviews(tx *gorp.Transaction) ([]Review, error) {

	reviews, err := selectToReviews(tx)
	if err != nil {
		return reviews, err
	}

	return reviews, nil
}

// GetReview レビュー情報を取得します
func GetReview(tx *gorp.Transaction) (Review, error) {

	review, err := selectToReview(tx)
	if err != nil {
		return review, err
	}

	return review, nil
}

func CreateReview(tx *gorp.Transaction, r Review) error {

	err := insertToReview(tx, r)
	if err != nil {
		return err
	}

	return nil
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

// selectToReview レビュー情報を検索します
func selectToReview(tx *gorp.Transaction) (Review, error) {
	var review Review
	err := tx.SelectOne(&review, `
		select
		  *
		from
		  review
		where
			review_id = ?
	`, 1)
	if err != nil {
		return review, err
	}

	return review, nil
}

func insertToReview(tx *gorp.Transaction, r Review) error {
	err := tx.Insert(r)
	if err != nil {
		return err
	}

	return nil
}
