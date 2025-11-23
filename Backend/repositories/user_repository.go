package repositories

import (
	"errors"

	"github.com/TaushifReza/go-event-booking-api/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository{
	return &UserRepository{DB: db}
}

// CREATE
func(r *UserRepository) Create(user *models.User)  error{
	return r.DB.Create(user).Error
}

// Find by email
func (r *UserRepository) GetByEmail(email string) (*models.User, error){
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Not found
		}
		return nil, err
	}
	return &user, nil
}

// Check if email exists
func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	if err := r.DB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}