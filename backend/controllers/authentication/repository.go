package authentication

import (
	"log"
	"sinarmas/kredit-sinarmas/models"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	Login(req DataRequest) (models.User, error)
	ChangePassword(req *RequestChangePassword, userID uint) error
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
	log.Printf("%+v", user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return models.User{}, err
	}
	// if user.Password != req.Password {
	// 	return models.User{}, fmt.Errorf("incorrect credentials")
	// }

	return user, nil
}

func (r *repository) ChangePassword(req *RequestChangePassword, userID uint) error {
	var user models.User
	user.ID = userID

	if err := r.db.First(&user).Error; err != nil {
		return err
	}
	log.Printf("%+v", user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return err
	}
	// // if user.Password != req.Password {
	// // 	return models.User{}, fmt.Errorf("incorrect credentials")
	// // }

	pb, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 8)
	user.Password = string(pb)
	if err := r.db.Select("password").Updates(&user).Error; err != nil {
		return err
	}

	return nil
}
