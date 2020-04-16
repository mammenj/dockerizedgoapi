package repository

import (
	"github.com/jinzhu/gorm"

	"myapp/app/myauth"
	"myapp/model"
)

func ListUsers(db *gorm.DB) (model.Users, error) {
	users := make([]*model.User, 0)
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}
func GetUser(db *gorm.DB, id uint) (*model.User, error) {
	user := &model.User{}
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByUserid(db *gorm.DB, userid string) (*model.User, error) {
	user := &model.User{}

	if err := db.Preload("Roles").First(&user, "userid = ?", userid).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(db *gorm.DB, user *model.User) (*model.User, error) {

	user.Password = myauth.HashAndSalt([]byte(user.Password))
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(db *gorm.DB, user *model.User) error {
	if err := db.First(&model.Book{}, user.ID).Update(user).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(db *gorm.DB, id uint) error {
	user := &model.User{}
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
