package business

import (
	"rsv-system-go/database"
	"rsv-system-go/middleweres"
	"rsv-system-go/models"

	"gorm.io/gorm"
)

var db *gorm.DB = database.InitDb()

func FetchUsers(users *[]models.User) (err error) {
	return db.Find(users).Error
}

func CreateUser(user *models.User) (err error) {
	return db.Create(user).Error
}

func FetchUserById(id int, user *models.User) (err error) {
	return db.Where("id = ?", id).First(user).Error
}

func UpdateUser(id int, user *models.User) (err error) {
	user.ID = id
	return db.Save(user).Error
}

func DeleteUser(id int, user *models.User) (err error) {
	user.ID = id
	return db.Delete(user).Error
}

func Login(user *models.User) (interface{}, error) {
	if err := db.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	token, err := middleweres.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	return &token, nil
}
