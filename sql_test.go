package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

// gunakan parameter pada query sql agar aman dari sql injection

func TestQuerySqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	name := "Bobi'; #"
	email := "bobi@email"

	script := "SELECT name FROM customer WHERE name = ? AND email = ?"
	rows, err := db.QueryContext(ctx, script, name, email)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Data benar", name)
	} else {
		fmt.Println("Data salah")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	name := "Surti"
	email := "surti@email"
	balance := 100000
	rating := 5.0
	dob := "1999-9-9"
	married := true

	script := "INSERT INTO customer(name, email, balance, rating, birth_date, married) VALUES(?, ?, ?, ?, ?, ?)"
	_, err := db.ExecContext(ctx, script, name, email, balance, rating, dob, married)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	name := "Burti"
	email := "burti@email"
	balance := 100000
	rating := 5.0
	dob := "1999-9-9"
	married := true

	script := "INSERT INTO customer(name, email, balance, rating, birth_date, married) VALUES(?, ?, ?, ?, ?, ?)"
	result, err := db.ExecContext(ctx, script, name, email, balance, rating, dob, married)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Last insert id : ", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(name, email, balance, rating, birth_date, married) VALUES(?, ?, ?, ?, ?, ?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		name := "Burti" + strconv.Itoa(i)
		email := "burti" + strconv.Itoa(i) + "@email"
		balance := 100000
		rating := 5.0
		dob := "1999-9-9"
		married := true

		result, err := statement.ExecContext(ctx, name, email, balance, rating, dob, married)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Account id : ", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO customer(name, email, balance, rating, birth_date, married) VALUES(?, ?, ?, ?, ?, ?)"
	// disini tidak mengguankan prepare statement tapi lebih baik jika menggunakannya
	// lakukan transaksi
	for i := 0; i < 10; i++ {
		name := "Borti" + strconv.Itoa(i)
		email := "borti" + strconv.Itoa(i) + "@email"
		balance := 100000
		rating := 5.0
		dob := "1999-9-9"
		married := true

		result, err := tx.ExecContext(ctx, script, name, email, balance, rating, dob, married)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Account id : ", id)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

func TestTransactionRollback(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO customer(name, email, balance, rating, birth_date, married) VALUES(?, ?, ?, ?, ?, ?)"
	// disini tidak mengguankan prepare statement tapi lebih baik jika menggunakannya
	// lakukan transaksi
	for i := 0; i < 10; i++ {
		name := "Borti" + strconv.Itoa(i)
		email := "borti" + strconv.Itoa(i) + "@email"
		balance := 100000
		rating := 5.0
		dob := "1999-9-9"
		married := true

		result, err := tx.ExecContext(ctx, script, name, email, balance, rating, dob, married)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Account id : ", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
