package services

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/TaushifReza/go-event-booking-api/dto"
	"github.com/TaushifReza/go-event-booking-api/internal/common"
	"github.com/TaushifReza/go-event-booking-api/models"
	"github.com/TaushifReza/go-event-booking-api/repositories"
	"github.com/TaushifReza/go-event-booking-api/utils"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService{
	return &UserService{Repo: repo}
}

func (s *UserService) Register(reqDto *dto.CreateUserRequest) (*models.User, error){
	reqDto.Email = strings.TrimSpace(strings.ToLower(reqDto.Email))

	hashedPassword, err := utils.HashPassword(reqDto.Password)
	if err != nil {
		return nil, &common.AppError{
			Code: http.StatusInternalServerError,
			Message: fmt.Errorf("something went wrong. please try again."),

		}
	}

	// check if email already exists
	exists, err := s.Repo.ExistsByEmail(reqDto.Email)
	if err != nil {
		return nil, errors.New("something went wrong. please try again")
	}
	if exists {
		return nil, &common.AppError{
			Code: http.StatusBadRequest,
			Message: fmt.Errorf("email %v already exists", reqDto.Email),
		}
	}

	user := &models.User{
		Email: reqDto.Email,
		Password: hashedPassword,
	}
	
	if err := s.Repo.Create(user); err != nil{
		return nil, &common.AppError{
			Code: http.StatusInternalServerError,
			Message: fmt.Errorf("something went wrong. please try again"),
		}
	}

	return user, nil
}

func (s *UserService) LoginUser(reqDto *dto.LoginRequest) (*dto.UserLoginResponse, error) {
	user, err := s.Repo.GetByEmail(reqDto.Email)
	if err != nil {
		return nil, &common.AppError{
			Code: http.StatusInternalServerError,
			Message: fmt.Errorf("something went wrong. please try again"),
		}
	}
	if user == nil {
		return nil, &common.AppError{
			Code: http.StatusBadRequest,
			Message: fmt.Errorf("invalid email or password"),
		}
	}

	// validate password
	if err := utils.CheckPassword(reqDto.Password, user.Password); err != nil{
		return nil, &common.AppError{
			Code: http.StatusBadRequest,
			Message: fmt.Errorf("invalid email or password"),
		}
	}

	// generate token
	accessToken, _ := utils.GenerateAccessToken(user.ID, user.Email)
	refreshToken, _ := utils.GenerateRefreshToken(user.ID, user.Email)

	res := &dto.UserLoginResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}

	return res, nil
}

func (s *UserService) VerifyRefreshToken(token string) (*dto.RefreshTokenResponse, error) {
	claims, err := utils.VerifyToken(token, "refresh")
	if err != nil {
		return nil, &common.AppError{
			Code: http.StatusInternalServerError,
			Message: err,
		}
	}

	// generate token
	accessToken, _ := utils.GenerateAccessToken(claims.ID, claims.Email)

	return &dto.RefreshTokenResponse{
		AccessToken: accessToken,
	}, nil
}


func (s *UserService) GetUserInfo(email string) (*dto.UserResponse, error){
	user, err := s.Repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("something went wrong. please try again")
	}
	if user == nil {
		return nil, errors.New("something went wrong. please try again")
	}

	res := &dto.UserResponse{
		ID: user.ID,
		Email: user.Email,
	}

	return res, nil
}