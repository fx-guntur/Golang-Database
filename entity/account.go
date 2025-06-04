package entity

import (
	"database/sql"
	"time"
)

type Account struct {
	Id         uint
	Name       string
	Email      sql.NullString
	Balance    int
	Rating     float64
	Created_at time.Time
	Birth_date sql.NullTime
	Married    bool
}
