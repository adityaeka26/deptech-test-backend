package usecase

import (
	"context"

	"github.com/adityaeka26/deptech-test-backend/config"
	"github.com/adityaeka26/deptech-test-backend/internal/dto"
	"github.com/adityaeka26/deptech-test-backend/internal/model"
	"github.com/adityaeka26/deptech-test-backend/internal/repository"
)

type categoryUsecase struct {
	config             *config.EnvConfig
	categoryRepository repository.CategoryRepository
}

type CategoryUsecase interface {
	CreateCategory(ctx context.Context, req dto.CreateCategoryReq) (*dto.CreateCategoryRes, error)
	GetCategoryByID(ctx context.Context, id uint) (*dto.GetCategoryByIDRes, error)
	GetAllCategories(ctx context.Context) ([]dto.GetCategoryByIDRes, error)
	UpdateCategory(ctx context.Context, req dto.UpdateCategoryReq) (*dto.UpdateCategoryRes, error)
	DeleteCategory(ctx context.Context, req dto.DeleteCategoryReq) error
}

func NewCategoryUsecase(config *config.EnvConfig, categoryRepository repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{
		config:             config,
		categoryRepository: categoryRepository,
	}
}

func (u *categoryUsecase) CreateCategory(ctx context.Context, req dto.CreateCategoryReq) (*dto.CreateCategoryRes, error) {
	txRepo, err := u.categoryRepository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txRepo.Rollback()
			panic(err)
		}
	}()

	category := &model.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := txRepo.Create(ctx, category); err != nil {
		txRepo.Rollback()
		return nil, err
	}

	if err := txRepo.Commit(); err != nil {
		return nil, err
	}

	return &dto.CreateCategoryRes{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (u *categoryUsecase) GetCategoryByID(ctx context.Context, id uint) (*dto.GetCategoryByIDRes, error) {
	category, err := u.categoryRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.GetCategoryByIDRes{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (u *categoryUsecase) GetAllCategories(ctx context.Context) ([]dto.GetCategoryByIDRes, error) {
	categories, err := u.categoryRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var res []dto.GetCategoryByIDRes
	for _, category := range categories {
		res = append(res, dto.GetCategoryByIDRes{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return res, nil
}

func (u *categoryUsecase) UpdateCategory(ctx context.Context, req dto.UpdateCategoryReq) (*dto.UpdateCategoryRes, error) {
	txRepo, err := u.categoryRepository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txRepo.Rollback()
			panic(err)
		}
	}()

	category := &model.Category{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := txRepo.Update(ctx, category); err != nil {
		txRepo.Rollback()
		return nil, err
	}

	if err := txRepo.Commit(); err != nil {
		txRepo.Rollback()
		return nil, err
	}

	return &dto.UpdateCategoryRes{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (u *categoryUsecase) DeleteCategory(ctx context.Context, req dto.DeleteCategoryReq) error {
	txRepo, err := u.categoryRepository.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			txRepo.Rollback()
			panic(err)
		}
	}()

	if err := txRepo.Delete(ctx, &model.Category{ID: req.ID}); err != nil {
		txRepo.Rollback()
		return err
	}

	if err := txRepo.Commit(); err != nil {
		txRepo.Rollback()
		return err
	}

	return nil
}
