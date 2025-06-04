package repository

import (
	"context"

	"github.com/adityaeka26/deptech-test-backend/internal/model"
	"gorm.io/gorm"
)

type transactionItemRepository struct {
	db *gorm.DB
}

type TransactionItemRepository interface {
	WithTx(tx *gorm.DB) TransactionItemRepository

	GetByID(ctx context.Context, id uint) (*model.TransactionItem, error)
	GetAll(ctx context.Context) ([]model.TransactionItem, error)

	Create(ctx context.Context, data *model.TransactionItem) error
	Update(ctx context.Context, data *model.TransactionItem) error
	Delete(ctx context.Context, data *model.TransactionItem) error
}

func NewTransactionItemRepository(db *gorm.DB) TransactionItemRepository {
	return &transactionItemRepository{
		db: db,
	}
}

func (r *transactionItemRepository) WithTx(tx *gorm.DB) TransactionItemRepository {
	return &transactionItemRepository{
		db: tx,
	}
}

func (r *transactionItemRepository) GetByID(ctx context.Context, id uint) (*model.TransactionItem, error) {
	var data model.TransactionItem
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *transactionItemRepository) GetAll(ctx context.Context) ([]model.TransactionItem, error) {
	var data []model.TransactionItem
	if err := r.db.WithContext(ctx).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *transactionItemRepository) Create(ctx context.Context, data *model.TransactionItem) error {
	return r.db.WithContext(ctx).Create(data).Error
}

func (r *transactionItemRepository) Update(ctx context.Context, data *model.TransactionItem) error {
	return r.db.WithContext(ctx).Save(data).Error
}

func (r *transactionItemRepository) Delete(ctx context.Context, data *model.TransactionItem) error {
	return r.db.WithContext(ctx).Delete(data).Error
}
