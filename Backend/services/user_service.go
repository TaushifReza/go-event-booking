package services

import (
	"errors"
	"fmt"

	"github.com/TaushifReza/go-event-booking-api/dto"
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
	hashedPassword, err := utils.HashPassword(reqDto.Password)
	if err != nil {
		fmt.Println("Password hashing ERROR: ", err)
		return nil, errors.New("something went wrong. please try again")
	}

	// check if email already exists
	exists, err := s.Repo.ExistsByEmail(reqDto.Email)
	if err != nil {
		return nil, errors.New("something went wrong. please try again")
	}
	if exists {
		return nil, fmt.Errorf("email %v already exists", reqDto.Email)
	}

	user := &models.User{
		Email: reqDto.Email,
		Password: hashedPassword,
	}
	
	if err := s.Repo.Create(user); err != nil{
		return nil, errors.New("something went wrong. please try again")
	}

	return user, nil
}

func (s *UserService) LoginUser(reqDto *dto.LoginRequest) (*dto.UserLoginResponse, error) {
	user, err := s.Repo.GetByEmail(reqDto.Email)
	if err != nil {
		return nil, errors.New("something went wrong. please try again")
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// validate password
	if err := utils.CheckPassword(reqDto.Password, user.Password); err != nil{
		return nil, errors.New("invalid email or password")
	}

	// generate token
	accessToken, _ := utils.GenerateAccessToken(user.ID, user.Email)
	refreshToken, _ := utils.GenerateRefreshToken(user.ID, user.Email)

	res := &dto.UserLoginResponse{
		ID: user.ID,
		Email: user.Email,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}

	return res, nil
}