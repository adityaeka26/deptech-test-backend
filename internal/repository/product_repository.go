package repository

import (
	"context"

	"github.com/adityaeka26/deptech-test-backend/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	WithTx(tx *gorm.DB) ProductRepository

	GetByID(ctx context.Context, id uint) (*model.Product, error)
	GetByIDLock(ctx context.Context, id uint) (*model.Product, error)
	GetAll(ctx context.Context) ([]model.Product, error)

	Create(ctx context.Context, product *model.Product) error
	Update(ctx context.Context, product *model.Product) error
	Delete(ctx context.Context, product *model.Product) error
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) WithTx(tx *gorm.DB) ProductRepository {
	return &productRepository{
		db: tx,
	}
}

func (r *productRepository) GetByID(ctx context.Context, id uint) (*model.Product, error) {
	var product model.Product
	if err := r.db.WithContext(ctx).Where("id = ?", id).Preload("Category").First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetByIDLock(ctx context.Context, id uint) (*model.Product, error) {
	var product model.Product
	if err := r.db.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", id).
		First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetAll(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.WithContext(ctx).Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) Create(ctx context.Context, product *model.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *productRepository) Update(ctx context.Context, product *model.Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

func (r *productRepository) Delete(ctx context.Context, product *model.Product) error {
	return r.db.WithContext(ctx).Delete(product).Error
}
