package repository

import (
	"context"

	"github.com/adityaeka26/deptech-test-backend/internal/model"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

type CategoryRepository interface {
	WithTx(tx *gorm.DB) CategoryRepository

	GetByID(ctx context.Context, id uint) (*model.Category, error)
	GetAll(ctx context.Context) ([]model.Category, error)

	Create(ctx context.Context, category *model.Category) error
	Update(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, category *model.Category) error
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) WithTx(tx *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: tx,
	}
}

func (r *categoryRepository) GetByID(ctx context.Context, id uint) (*model.Category, error) {
	var category model.Category
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	if err := r.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) Create(ctx context.Context, category *model.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *categoryRepository) Update(ctx context.Context, category *model.Category) error {
	return r.db.WithContext(ctx).Save(category).Error
}

func (r *categoryRepository) Delete(ctx context.Context, category *model.Category) error {
	return r.db.WithContext(ctx).Delete(category).Error
}
