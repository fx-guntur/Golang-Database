package repository

import (
	"belajar-golang-database/entity"
	"context"
)

type AccountRepository interface {
	// bisa juga tambah parameter tx / transaction jika dibutuhkan
	Insert(ctx context.Context, account entity.Account) (entity.Account, error)
	FindById(ctx context.Context, id uint) (entity.Account, error)
	FindAll(ctx context.Context) ([]entity.Account, error)
}
