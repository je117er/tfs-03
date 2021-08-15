package utils

import (
	"fmt"
	"strings"
	"time"
)

type PersonComponent struct {
	Name string
	URL string
}

type AggregateRating struct {
	RatingCount int
	BestRating  float64
	WorstRating float64
	RatingValue float64
}

type Review struct {
	InLanguage string
	Name       string
	ReviewBody string
	DateCreated JsonDate
}
type JsonDate struct {
	time.Time
}

type Movie struct {
	Name            string
	Description     string
	URL string
	Review          Review
	AggregateRating AggregateRating
	Genre           []string
	Keywords string
	Duration string
	Actor []PersonComponent
	Director []PersonComponent
	Creator []PersonComponent
	DatePublished JsonDate

}

const ctLayout = "2006-01-02"

func (ct *JsonDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (ct *JsonDate) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

var nilTime = (time.Time{}).UnixNano()

func (ct *JsonDate) IsSet() bool {
	return ct.UnixNano() != nilTime
}
