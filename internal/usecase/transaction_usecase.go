package usecase

import (
	"github.com/adityaeka26/deptech-test-backend/config"
	"github.com/adityaeka26/deptech-test-backend/internal/repository"
)

type transactionUsecase struct {
	config                *config.EnvConfig
	productRepository     repository.ProductRepository
	transactionRepository repository.TransactionRepository
}

type TransactionUsecase interface {
}

func NewTransactionRepository(config *config.EnvConfig, productRepository repository.ProductRepository, transactionRepository repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{
		config:                config,
		transactionRepository: transactionRepository,
		productRepository:     productRepository,
	}
}

// func (u *transactionUsecase) CreateTransaction(ctx context.Context, req dto.CreateTransactionReq) (*dto.CreateTransactionRes, error) {
// 	txRepo, err := u.transactionRepository.BeginTx(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer func() {
// 		if r := recover(); r != nil {
// 			txRepo.Rollback()
// 			panic(err)
// 		}
// 	}()

// 	IDs := []uint{}
// 	for _, item := range req.Items {
// 		IDs = append(IDs, item.ProductID)
// 	}
// 	products, err := txRepo.GetByMultipleIDs(ctx, IDs)

// 	if err := txRepo.Commit(); err != nil {
// 		txRepo.Rollback()
// 		return nil, err
// 	}

// 	return nil, nil
// }
