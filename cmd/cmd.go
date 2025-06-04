package cmd

import (
	"context"

	"github.com/adityaeka26/deptech-test-backend/cmd/rest"
	"github.com/adityaeka26/deptech-test-backend/config"
	"github.com/adityaeka26/deptech-test-backend/internal/middleware"
	"github.com/adityaeka26/deptech-test-backend/internal/model"
	"github.com/adityaeka26/deptech-test-backend/internal/repository"
	"github.com/adityaeka26/deptech-test-backend/internal/usecase"
	"github.com/adityaeka26/deptech-test-backend/pkg/minio"
	"github.com/adityaeka26/deptech-test-backend/pkg/mysql"
)

func Execute() {
	config, err := config.Load(".env")
	if err != nil {
		panic(err)
	}

	mysql, err := mysql.NewMySql(
		config.MySqlUsername,
		config.MySqlPassword,
		config.MySqlHost,
		config.MySqlPort,
		config.MySqlDbName,
		config.MySqlSslMode,
	)
	if err != nil {
		panic(err)
	}

	err = mysql.Db.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Product{},
		&model.Transaction{},
		&model.TransactionItem{},
	)
	if err != nil {
		panic(err)
	}

	minio, err := minio.NewMinio(config)
	if err != nil {
		panic(err)
	}

	err = minio.CreateBucket(context.Background(), "products")
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(mysql.Db)
	categoryRepository := repository.NewCategoryRepository(mysql.Db)
	productRepository := repository.NewProductRepository(mysql.Db)
	transactionRepository := repository.NewTransactionRepository(mysql.Db)
	transactionItemRepository := repository.NewTransactionItemRepository(mysql.Db)

	userUsecase := usecase.NewUserUsecase(config, mysql.Db, userRepository)
	categoryUsecase := usecase.NewCategoryUsecase(config, mysql.Db, categoryRepository)
	productUsecase := usecase.NewProductUsecase(config, mysql.Db, minio, productRepository)
	transactionUsecase := usecase.NewTransactionUsecase(config, mysql.Db, productRepository, transactionRepository, transactionItemRepository)

	middleware := middleware.NewMiddleware()

	err = rest.ServeRest(config, userUsecase, categoryUsecase, productUsecase, transactionUsecase, middleware)
	if err != nil {
		panic(err)
	}
}
