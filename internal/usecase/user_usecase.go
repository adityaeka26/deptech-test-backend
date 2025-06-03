package usecase

import (
	"context"
	"time"

	"github.com/adityaeka26/deptech-test-backend/internal/dto"
	"github.com/adityaeka26/deptech-test-backend/internal/model"
	"github.com/adityaeka26/deptech-test-backend/internal/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

type UserUsecase interface {
	CreateUser(ctx context.Context, req dto.CreateUserReq) (*dto.CreateUserRes, error)
	GetUserByID(ctx context.Context, req dto.GetUserByIDReq) (*dto.GetUserByIDRes, error)
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) CreateUser(ctx context.Context, req dto.CreateUserReq) (*dto.CreateUserRes, error) {
	txRepo, err := u.userRepository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txRepo.Rollback()
			panic(err)
		}
	}()

	dob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		txRepo.Rollback()
		return nil, err
	}

	user := model.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		DateOfBirth: dob,
		Gender:      req.Gender,
	}

	if err := txRepo.Create(ctx, &user); err != nil {
		txRepo.Rollback()
		return nil, err
	}

	if err := txRepo.Commit(); err != nil {
		txRepo.Rollback()
		return nil, err
	}

	return &dto.CreateUserRes{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth.Format("2006-01-02"),
		Gender:      user.Gender,
	}, nil
}

func (u *userUsecase) GetUserByID(ctx context.Context, req dto.GetUserByIDReq) (*dto.GetUserByIDRes, error) {
	user, err := u.userRepository.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &dto.GetUserByIDRes{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth.Format("2006-01-02"),
		Gender:      user.Gender,
	}, nil
}
