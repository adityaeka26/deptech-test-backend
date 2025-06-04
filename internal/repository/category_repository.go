package repository

import (
	"context"

	"github.com/adityaeka26/deptech-test-backend/internal/model"
	"github.com/adityaeka26/deptech-test-backend/pkg/mysql"
	"gorm.io/gorm"
)

type categoryRepository struct {
	mysql *mysql.MySql
}

type CategoryRepository interface {
	GetByID(ctx context.Context, id uint) (*model.Category, error)
	GetByEmail(ctx context.Context, email string) (*model.Category, error)
	GetAll(ctx context.Context) ([]model.Category, error)

	BeginTx(ctx context.Context) (CategoryTxRepository, error)
}

func NewCategoryRepository(mysql *mysql.MySql) CategoryRepository {
	return &categoryRepository{
		mysql: mysql,
	}
}

func (r *categoryRepository) GetByID(ctx context.Context, id uint) (*model.Category, error) {
	var category model.Category
	if err := r.mysql.Db.WithContext(ctx).Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) GetByEmail(ctx context.Context, email string) (*model.Category, error) {
	var category model.Category
	if err := r.mysql.Db.WithContext(ctx).Where("email = ?", email).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	if err := r.mysql.Db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) BeginTx(ctx context.Context) (CategoryTxRepository, error) {
	tx := r.mysql.Db.WithContext(ctx).Begin()
	return &categoryTxRepository{
		tx: tx,
	}, nil
}

// Transactional
type categoryTxRepository struct {
	tx *gorm.DB
}

type CategoryTxRepository interface {
	Create(ctx context.Context, category *model.Category) error
	Update(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, category *model.Category) error

	Commit() error
	Rollback() error
}

func (r *categoryTxRepository) Create(ctx context.Context, category *model.Category) error {
	return r.tx.WithContext(ctx).Create(category).Error
}

func (r *categoryTxRepository) Update(ctx context.Context, category *model.Category) error {
	return r.tx.WithContext(ctx).Save(category).Error
}

func (r *categoryTxRepository) Delete(ctx context.Context, category *model.Category) error {
	return r.tx.WithContext(ctx).Delete(category).Error
}

func (r *categoryTxRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *categoryTxRepository) Rollback() error {
	return r.tx.Rollback().Error
}
