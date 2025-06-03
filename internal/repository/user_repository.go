package repository

import (
	"context"

	"github.com/adityaeka26/deptech-test-backend/internal/model"
	"github.com/adityaeka26/deptech-test-backend/pkg/mysql"
	"gorm.io/gorm"
)

type userRepository struct {
	mysql *mysql.MySql
}

type UserRepository interface {
	GetByID(ctx context.Context, id uint) (*model.User, error)

	BeginTx(ctx context.Context) (UserTxRepository, error)
}

func NewUserRepository(mysql *mysql.MySql) UserRepository {
	return &userRepository{
		mysql: mysql,
	}
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	if err := r.mysql.Db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) BeginTx(ctx context.Context) (UserTxRepository, error) {
	tx := r.mysql.Db.WithContext(ctx).Begin()
	return &userTxRepository{
		tx: tx,
	}, nil
}

// Transactional
type userTxRepository struct {
	tx *gorm.DB
}

type UserTxRepository interface {
	Create(ctx context.Context, user *model.User) error

	Commit() error
	Rollback() error
}

func (r *userTxRepository) Create(ctx context.Context, user *model.User) error {
	return r.tx.WithContext(ctx).Create(user).Error
}

func (r *userTxRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *userTxRepository) Rollback() error {
	return r.tx.Rollback().Error
}
