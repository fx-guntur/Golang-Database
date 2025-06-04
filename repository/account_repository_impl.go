package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"
)

type accountRepositoryImpl struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepositoryImpl{DB: db}
}

// bisa juga tambah parameter tx / transaction jika dibutuhkan
func (repository *accountRepositoryImpl) Insert(ctx context.Context, account entity.Account) (entity.Account, error) {
	script := "INSERT INTO customer(name, email, balance, rating, birth_date, married) VALUES(?, ?, ?, ?, ?, ?)"

	var email sql.NullString
	if account.Email.Valid {
		email = account.Email
	} else {
		email = sql.NullString{String: "", Valid: false}
	}

	var birthDate sql.NullTime
	if account.Birth_date.Valid {
		birthDate = account.Birth_date
	} else {
		birthDate = sql.NullTime{Time: time.Time{}, Valid: false}
	}

	result, err := repository.DB.ExecContext(ctx, script, account.Name, email.String, account.Balance, account.Rating, birthDate.Time, account.Married)
	if err != nil {
		return account, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return account, err
	}
	account.Id = uint(id)
	return account, nil
}

func (repository *accountRepositoryImpl) FindById(ctx context.Context, id uint) (entity.Account, error) {
	script := "SELECT id, name, email FROM customer WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	account := entity.Account{}
	if err != nil {
		return account, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&account.Id, &account.Name, &account.Email)
		return account, nil
	} else {
		return account, errors.New("Id " + strconv.Itoa(int(id)) + " Not found")
	}
}

func (repository *accountRepositoryImpl) FindAll(ctx context.Context) ([]entity.Account, error) {
	script := "SELECT id, name, email FROM customer"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var accounts []entity.Account
	for rows.Next() {
		account := entity.Account{}
		rows.Scan(&account.Id, &account.Name, &account.Email)
		accounts = append(accounts, account)
	}
	return accounts, nil
}
