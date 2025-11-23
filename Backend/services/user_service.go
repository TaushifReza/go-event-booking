package services

import (
	"github.com/TaushifReza/go-event-booking-api/dto"
	"github.com/TaushifReza/go-event-booking-api/models"
	"github.com/TaushifReza/go-event-booking-api/utils"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) Create(dto *dto.CreateUserRequest) (*models.User, error){
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Email: dto.Email,
		Password: hashedPassword,
	}
	
	if err := s.DB.Create(&user).Error; err != nil{
		return nil, err
	}

	return &user, nil
}


func NewUserService(db *gorm.DB) *UserService{
	return &UserService{DB: db}
}