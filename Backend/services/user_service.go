package services

import (
	"errors"

	"github.com/TaushifReza/go-event-booking-api/dto"
	"github.com/TaushifReza/go-event-booking-api/models"
	"github.com/TaushifReza/go-event-booking-api/utils"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) Create(resDto *dto.CreateUserRequest) (*models.User, error){
	hashedPassword, err := utils.HashPassword(resDto.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Email: resDto.Email,
		Password: hashedPassword,
	}
	
	if err := s.DB.Create(&user).Error; err != nil{
		return nil, err
	}

	return &user, nil
}

func (s *UserService) LoginUser(reqDto *dto.LoginRequest) (*dto.UserLoginResponse, error) {
	var user models.User

	if err := s.DB.Where("email = ?", reqDto.Email).First(&user).Error; err != nil{
		return nil, errors.New("invalid email or password")
	}

	// validate password
	if err := utils.CheckPassword(reqDto.Password, user.Password); err != nil{
		return nil, errors.New("invalid email or password")
	}

	// generate token
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	res := dto.UserLoginResponse{
		ID: user.ID,
		Email: user.Email,
		Token: token,
	}

	return &res, nil
}

func NewUserService(db *gorm.DB) *UserService{
	return &UserService{DB: db}
}