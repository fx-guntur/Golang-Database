package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	accountRepository := NewAccountRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	account := entity.Account{
		Name:       "Gengar",
		Email:      sql.NullString{String: "gengar@email.com", Valid: true},
		Balance:    100000,
		Rating:     5.0,
		Birth_date: sql.NullTime{Time: time.Date(1999, 9, 9, 0, 0, 0, 0, time.UTC), Valid: true},
		Married:    false,
	}

	result, err := accountRepository.Insert(ctx, account)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	accountRepository := NewAccountRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	account, err := accountRepository.FindById(ctx, 5)
	if err != nil {
		panic(err)
	}
	fmt.Println(account)
}

func TestFindAll(t *testing.T) {
	accountRepository := NewAccountRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	accounts, err := accountRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, account := range accounts {
		fmt.Println(account)
	}
}
