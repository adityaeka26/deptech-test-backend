package repository

import (
	"context"

	"github.com/adityaeka26/deptech-test-backend/internal/model"
	"github.com/adityaeka26/deptech-test-backend/pkg/mysql"
	"gorm.io/gorm"
)

type transactionRepository struct {
	mysql *mysql.MySql
}

type TransactionRepository interface {
	GetByID(ctx context.Context, id uint) (*model.Transaction, error)
	GetAll(ctx context.Context) ([]model.Transaction, error)

	BeginTx(ctx context.Context) (TransactionTxRepository, error)
}

func NewTransactionRepository(mysql *mysql.MySql) TransactionRepository {
	return &transactionRepository{
		mysql: mysql,
	}
}

func (r *transactionRepository) GetByID(ctx context.Context, id uint) (*model.Transaction, error) {
	var data model.Transaction
	if err := r.mysql.Db.WithContext(ctx).Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *transactionRepository) GetAll(ctx context.Context) ([]model.Transaction, error) {
	var data []model.Transaction
	if err := r.mysql.Db.WithContext(ctx).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *transactionRepository) BeginTx(ctx context.Context) (TransactionTxRepository, error) {
	tx := r.mysql.Db.WithContext(ctx).Begin()
	return &transactionTxRepository{
		tx: tx,
	}, nil
}

// Transactional
type transactionTxRepository struct {
	tx *gorm.DB
}

type TransactionTxRepository interface {
	Create(ctx context.Context, data *model.Transaction) error
	Update(ctx context.Context, data *model.Transaction) error
	Delete(ctx context.Context, data *model.Transaction) error

	Commit() error
	Rollback() error
}

func (r *transactionTxRepository) Create(ctx context.Context, data *model.Transaction) error {
	return r.tx.WithContext(ctx).Create(data).Error
}

func (r *transactionTxRepository) Update(ctx context.Context, data *model.Transaction) error {
	return r.tx.WithContext(ctx).Save(data).Error
}

func (r *transactionTxRepository) Delete(ctx context.Context, data *model.Transaction) error {
	return r.tx.WithContext(ctx).Delete(data).Error
}

func (r *transactionTxRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *transactionTxRepository) Rollback() error {
	return r.tx.Rollback().Error
}
