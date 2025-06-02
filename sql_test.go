package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(name, email, balance, rating, birth_date, married) VALUES('Doobi', 'doobi@email', 100000, 5.0, '1999-9-9', true)"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestExecSqlNull(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(name, balance, rating, married) VALUES('Doobi', 100000, 5.0, true)"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestDeleteSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "DELETE FROM customer"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success delete all customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id uint
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id :", id)
		fmt.Println("Name :", name)
	}

	defer rows.Close()
}

/*
		kalau terjadi error

		Scan error on column index 5, name "birth_date": unsupported Scan, storing driver.Value type []uint8 into type
		*time.Time [recovered]
	    panic: sql: Scan error on column index 5, name "birth_date": unsupported Scan, storing driver.Value type
		[]uint8 into type *time.Time

		tinggal tambahkan ?parseTime=true di connection nya
*/
func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id uint
		var name string
		var email sql.NullString
		var balance int
		var rating float64
		var created_at time.Time
		var birth_date sql.NullTime
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("===================")
		fmt.Println("id :", id)
		fmt.Println("Name :", name)

		if email.Valid {
			fmt.Println("Email :", email.String)
		}

		fmt.Println("Balance :", balance)

		if birth_date.Valid {
			fmt.Println("Date of Birth :", birth_date.Time)
		}

		fmt.Println("Married status :", married)
		fmt.Println("Created at :", created_at)
	}

	defer rows.Close()
}
