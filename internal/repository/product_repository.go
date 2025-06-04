package repository

import (
	"context"

	"github.com/adityaeka26/deptech-test-backend/internal/model"
	"github.com/adityaeka26/deptech-test-backend/pkg/mysql"
	"gorm.io/gorm"
)

type productRepository struct {
	mysql *mysql.MySql
}

type ProductRepository interface {
	GetByID(ctx context.Context, id uint) (*model.Product, error)
	GetByEmail(ctx context.Context, email string) (*model.Product, error)
	GetAll(ctx context.Context) ([]model.Product, error)

	BeginTx(ctx context.Context) (ProductTxRepository, error)
}

func NewProductRepository(mysql *mysql.MySql) ProductRepository {
	return &productRepository{
		mysql: mysql,
	}
}

func (r *productRepository) GetByID(ctx context.Context, id uint) (*model.Product, error) {
	var product model.Product
	if err := r.mysql.Db.WithContext(ctx).Where("id = ?", id).Preload("Category").First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetByEmail(ctx context.Context, email string) (*model.Product, error) {
	var product model.Product
	if err := r.mysql.Db.WithContext(ctx).Where("email = ?", email).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetAll(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	if err := r.mysql.Db.WithContext(ctx).Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) BeginTx(ctx context.Context) (ProductTxRepository, error) {
	tx := r.mysql.Db.WithContext(ctx).Begin()
	return &productTxRepository{
		tx: tx,
	}, nil
}

// Transactional
type productTxRepository struct {
	tx *gorm.DB
}

type ProductTxRepository interface {
	Create(ctx context.Context, product *model.Product) error
	Update(ctx context.Context, product *model.Product) error
	Delete(ctx context.Context, product *model.Product) error

	Commit() error
	Rollback() error
}

func (r *productTxRepository) Create(ctx context.Context, product *model.Product) error {
	return r.tx.WithContext(ctx).Create(product).Error
}

func (r *productTxRepository) Update(ctx context.Context, product *model.Product) error {
	return r.tx.WithContext(ctx).Save(product).Error
}

func (r *productTxRepository) Delete(ctx context.Context, product *model.Product) error {
	return r.tx.WithContext(ctx).Delete(product).Error
}

func (r *productTxRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *productTxRepository) Rollback() error {
	return r.tx.Rollback().Error
}
