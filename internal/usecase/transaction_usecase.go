package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/adityaeka26/deptech-test-backend/config"
	"github.com/adityaeka26/deptech-test-backend/internal/dto"
	"github.com/adityaeka26/deptech-test-backend/internal/model"
	"github.com/adityaeka26/deptech-test-backend/internal/repository"
	pkgError "github.com/adityaeka26/deptech-test-backend/pkg/error"
	"gorm.io/gorm"
)

type transactionUsecase struct {
	config                    *config.EnvConfig
	db                        *gorm.DB
	productRepository         repository.ProductRepository
	transactionRepository     repository.TransactionRepository
	transactionItemRepository repository.TransactionItemRepository
}

type TransactionUsecase interface {
	CreateTransaction(ctx context.Context, req dto.CreateTransactionReq) (*dto.CreateTransactionRes, error)
	GetAllTransactions(ctx context.Context) ([]dto.GetAllTransactionsRes, error)
}

func NewTransactionUsecase(config *config.EnvConfig, db *gorm.DB, productRepository repository.ProductRepository, transactionRepository repository.TransactionRepository, transactionItemRepository repository.TransactionItemRepository) TransactionUsecase {
	return &transactionUsecase{
		config:                    config,
		transactionRepository:     transactionRepository,
		productRepository:         productRepository,
		db:                        db,
		transactionItemRepository: transactionItemRepository,
	}
}

func (u *transactionUsecase) CreateTransaction(ctx context.Context, req dto.CreateTransactionReq) (*dto.CreateTransactionRes, error) {
	tx := u.db.WithContext(ctx).Begin()
	transactionRepositoryTx := u.transactionRepository.WithTx(tx)
	transactionItemRepositoryTx := u.transactionItemRepository.WithTx(tx)
	productRepositoryTx := u.productRepository.WithTx(tx)

	products := make([]*model.Product, 0, len(req.Items))
	for _, item := range req.Items {
		product, err := productRepositoryTx.GetByIDLock(ctx, item.ProductID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		if req.Type == "out" {
			if product.Stock < item.Quantity {
				tx.Rollback()
				return nil, pkgError.Conflict(fmt.Sprintf("insufficient stock for product %s", product.Name))
			}
			product.Stock -= item.Quantity
		} else {
			product.Stock += item.Quantity
		}

		err = productRepositoryTx.Update(ctx, product)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		products = append(products, product)
	}

	transaction := &model.Transaction{
		UserID:    req.UserID,
		Type:      req.Type,
		DeletedAt: time.Time{},
	}
	err := transactionRepositoryTx.Create(ctx, transaction)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, product := range products {
		transactionItem := &model.TransactionItem{
			TransactionID: transaction.ID,
			ProductID:     product.ID,
			Quantity:      product.Stock,
		}

		if err := transactionItemRepositoryTx.Create(ctx, transactionItem); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return nil, nil
}

func (u *transactionUsecase) GetAllTransactions(ctx context.Context) ([]dto.GetAllTransactionsRes, error) {
	transactions, err := u.transactionRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var res []dto.GetAllTransactionsRes
	for _, transaction := range transactions {
		items := []dto.GetAllTransactionsResItems{}
		for _, item := range transaction.Items {
			items = append(items, dto.GetAllTransactionsResItems{
				ID:       item.ID,
				Quantity: item.Quantity,
				Product: dto.GetAllTransactionsResItemsProduct{
					ID:          item.Product.ID,
					Name:        item.Product.Name,
					Description: item.Product.Description,
					ImageUrl:    item.Product.ImagePath,
					CategoryID:  item.Product.CategoryID,
					Stock:       item.Product.Stock,
				},
			})
		}

		res = append(res, dto.GetAllTransactionsRes{
			ID: transaction.ID,
			User: dto.GetAllTransactionsResUser{
				ID:          transaction.User.ID,
				FirstName:   transaction.User.FirstName,
				LastName:    transaction.User.LastName,
				Email:       transaction.User.Email,
				DateOfBirth: transaction.User.DateOfBirth.Format("2006-01-02"),
				Gender:      transaction.User.Gender,
			},
			Type:      transaction.Type,
			CreatedAt: transaction.CreatedAt,
			Items:     items,
		})
	}

	return res, nil
}
