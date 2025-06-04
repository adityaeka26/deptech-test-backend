package usecase

import (
	"context"
	"time"

	"github.com/adityaeka26/deptech-test-backend/config"
	"github.com/adityaeka26/deptech-test-backend/internal/dto"
	"github.com/adityaeka26/deptech-test-backend/internal/model"
	"github.com/adityaeka26/deptech-test-backend/internal/repository"
	"github.com/adityaeka26/deptech-test-backend/pkg/minio"
	"gorm.io/gorm"
)

type productUsecase struct {
	config            *config.EnvConfig
	db                *gorm.DB
	minio             *minio.Minio
	productRepository repository.ProductRepository
}

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req dto.CreateProductReq) (*dto.CreateProductRes, error)
	GetProductByID(ctx context.Context, req dto.GetProductByIDReq) (*dto.GetProductByIDRes, error)
	GetAllProducts(ctx context.Context) ([]dto.GetProductByIDRes, error)
	UpdateProduct(ctx context.Context, req dto.UpdateProductReq) (*dto.UpdateProductRes, error)
	DeleteProduct(ctx context.Context, req dto.DeleteProductReq) error
}

func NewProductUsecase(config *config.EnvConfig, db *gorm.DB, minio *minio.Minio, productRepository repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		config:            config,
		db:                db,
		minio:             minio,
		productRepository: productRepository,
	}
}

func (u *productUsecase) CreateProduct(ctx context.Context, req dto.CreateProductReq) (*dto.CreateProductRes, error) {
	tx := u.db.WithContext(ctx).Begin()
	productRepositoryTx := u.productRepository.WithTx(tx)

	err := u.minio.Upload(ctx, "products", req.Image)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// fmt.Println(req.Image.Filename)

	product := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		ImagePath:   req.Image.Filename,
		CategoryID:  req.CategoryID,
		Stock:       req.Stock,
	}

	if err := productRepositoryTx.Create(ctx, product); err != nil {
		tx.Rollback()
		return nil, err
	}

	imageUrl, err := u.minio.GeneratePresignedURL(ctx, "products", product.ImagePath, time.Minute*5)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &dto.CreateProductRes{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		ImageUrl:    imageUrl.String(),
		CategoryID:  product.CategoryID,
		Stock:       product.Stock,
	}, nil
}

func (u *productUsecase) GetProductByID(ctx context.Context, req dto.GetProductByIDReq) (*dto.GetProductByIDRes, error) {
	product, err := u.productRepository.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	imageUrl, err := u.minio.GeneratePresignedURL(ctx, "products", product.ImagePath, time.Minute*5)
	if err != nil {
		return nil, err
	}
	return &dto.GetProductByIDRes{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		ImageUrl:    imageUrl.String(),
		Stock:       product.Stock,
		Category: dto.GetCategoryByIDRes{
			ID:          product.Category.ID,
			Name:        product.Category.Name,
			Description: product.Category.Description,
		},
	}, nil
}

func (u *productUsecase) GetAllProducts(ctx context.Context) ([]dto.GetProductByIDRes, error) {
	products, err := u.productRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var res []dto.GetProductByIDRes
	for _, product := range products {
		imageUrl, err := u.minio.GeneratePresignedURL(ctx, "products", product.ImagePath, time.Minute*5)
		if err != nil {
			return nil, err
		}
		res = append(res, dto.GetProductByIDRes{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			ImageUrl:    imageUrl.String(),
			Stock:       product.Stock,
			Category: dto.GetCategoryByIDRes{
				ID:          product.Category.ID,
				Name:        product.Category.Name,
				Description: product.Category.Description,
			},
		})
	}

	return res, nil
}

func (u *productUsecase) UpdateProduct(ctx context.Context, req dto.UpdateProductReq) (*dto.UpdateProductRes, error) {
	_, err := u.productRepository.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	tx := u.db.WithContext(ctx).Begin()
	productRepositoryTx := u.productRepository.WithTx(tx)

	err = u.minio.Upload(ctx, "products", req.Image)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// fmt.Println(req.Image.Filename)

	product := &model.Product{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		ImagePath:   req.Image.Filename,
		CategoryID:  req.CategoryID,
		Stock:       req.Stock,
	}

	if err := productRepositoryTx.Update(ctx, product); err != nil {
		tx.Rollback()
		return nil, err
	}

	imageUrl, err := u.minio.GeneratePresignedURL(ctx, "products", product.ImagePath, time.Minute*5)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &dto.UpdateProductRes{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		ImageUrl:    imageUrl.String(),
		CategoryID:  product.CategoryID,
		Stock:       product.Stock,
	}, nil
}

func (u *productUsecase) DeleteProduct(ctx context.Context, req dto.DeleteProductReq) error {
	tx := u.db.WithContext(ctx).Begin()
	productRepositoryTx := u.productRepository.WithTx(tx)

	if err := productRepositoryTx.Delete(ctx, &model.Product{ID: req.ID}); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
