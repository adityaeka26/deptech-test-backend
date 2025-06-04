package repository

import (
	"context"

	"github.com/adityaeka26/deptech-test-backend/internal/model"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

type TransactionRepository interface {
	WithTx(tx *gorm.DB) TransactionRepository

	GetByID(ctx context.Context, id uint) (*model.Transaction, error)
	GetAll(ctx context.Context) ([]model.Transaction, error)

	Create(ctx context.Context, data *model.Transaction) error
	Update(ctx context.Context, data *model.Transaction) error
	Delete(ctx context.Context, data *model.Transaction) error
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) WithTx(tx *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: tx,
	}
}

func (r *transactionRepository) GetByID(ctx context.Context, id uint) (*model.Transaction, error) {
	var data model.Transaction
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *transactionRepository) GetAll(ctx context.Context) ([]model.Transaction, error) {
	var data []model.Transaction
	if err := r.db.WithContext(ctx).
		Preload("Items").
		Preload("Items.Product").
		Preload("User").
		Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *transactionRepository) Create(ctx context.Context, data *model.Transaction) error {
	return r.db.WithContext(ctx).Create(data).Error
}

func (r *transactionRepository) Update(ctx context.Context, data *model.Transaction) error {
	return r.db.WithContext(ctx).Save(data).Error
}

func (r *transactionRepository) Delete(ctx context.Context, data *model.Transaction) error {
	return r.db.WithContext(ctx).Delete(data).Error
}
