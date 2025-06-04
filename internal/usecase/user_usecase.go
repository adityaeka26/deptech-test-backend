package usecase

import (
	"context"
	"time"

	"github.com/adityaeka26/deptech-test-backend/config"
	"github.com/adityaeka26/deptech-test-backend/internal/dto"
	"github.com/adityaeka26/deptech-test-backend/internal/model"
	"github.com/adityaeka26/deptech-test-backend/internal/repository"
	"github.com/adityaeka26/deptech-test-backend/pkg/helper"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	config         *config.EnvConfig
	userRepository repository.UserRepository
}

type UserUsecase interface {
	CreateUser(ctx context.Context, req dto.CreateUserReq) (*dto.CreateUserRes, error)
	GetUserByID(ctx context.Context, req dto.GetUserByIDReq) (*dto.GetUserByIDRes, error)
	UpdateUser(ctx context.Context, req dto.UpdateUserReq) (*dto.UpdateUserRes, error)
	DeleteUser(ctx context.Context, req dto.DeleteUserReq) error
	GetAllUsers(ctx context.Context) ([]dto.GetUserByIDRes, error)

	LoginUser(ctx context.Context, req dto.LoginUserReq) (*dto.LoginUserRes, error)
}

func NewUserUsecase(config *config.EnvConfig, userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		config:         config,
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := model.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    string(hashedPassword),
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

func (u *userUsecase) UpdateUser(ctx context.Context, req dto.UpdateUserReq) (*dto.UpdateUserRes, error) {
	user, err := u.userRepository.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user = &model.User{
		ID:          req.ID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    string(hashedPassword),
		DateOfBirth: dob,
		Gender:      req.Gender,
	}

	if err := txRepo.Update(ctx, user); err != nil {
		txRepo.Rollback()
		return nil, err
	}

	if err := txRepo.Commit(); err != nil {
		txRepo.Rollback()
		return nil, err
	}

	return &dto.UpdateUserRes{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth.Format("2006-01-02"),
		Gender:      user.Gender,
	}, nil
}

func (u *userUsecase) DeleteUser(ctx context.Context, req dto.DeleteUserReq) error {
	txRepo, err := u.userRepository.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			txRepo.Rollback()
			panic(err)
		}
	}()

	if err := txRepo.Delete(ctx, &model.User{ID: req.ID}); err != nil {
		txRepo.Rollback()
		return err
	}

	if err := txRepo.Commit(); err != nil {
		txRepo.Rollback()
		return err
	}

	return nil
}

func (u *userUsecase) GetAllUsers(ctx context.Context) ([]dto.GetUserByIDRes, error) {
	users, err := u.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var res []dto.GetUserByIDRes
	for _, user := range users {
		res = append(res, dto.GetUserByIDRes{
			ID:          user.ID,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			DateOfBirth: user.DateOfBirth.Format("2006-01-02"),
			Gender:      user.Gender,
		})
	}

	return res, nil
}

func (u *userUsecase) LoginUser(ctx context.Context, req dto.LoginUserReq) (*dto.LoginUserRes, error) {
	user, err := u.userRepository.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	claims := make(jwt.MapClaims)
	claims["data"] = map[string]any{
		"id":    user.ID,
		"email": user.Email,
	}
	token, err := helper.GenerateToken(u.config.JwtPrivateKey, claims)
	if err != nil {
		return nil, err
	}

	return &dto.LoginUserRes{
		Token: *token,
	}, nil
}
