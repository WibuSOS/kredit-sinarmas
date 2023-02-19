package authentication

import (
	"sinarmas/kredit-sinarmas/models"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	Login(req DataRequest) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Login(req DataRequest) (models.User, error) {
	var user models.User

	if err := r.db.Take(&user, "username = ?", strings.TrimSpace(req.Username)).Error; err != nil {
		return models.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return models.User{}, err
	}
	// if user.Password != req.Password {
	// 	return models.User{}, fmt.Errorf("incorrect credentials")
	// }

	return user, nil
}
