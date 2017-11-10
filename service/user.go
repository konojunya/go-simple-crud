package service

import "github.com/konojunya/go-simple-crud/model"

func CreateUser(name string, password string) error {
	return db.Create(&model.User{
		Name:     name,
		Password: password,
	}).Error
}

func FindUserAll() (*[]model.User, error) {
	var users []model.User
	err := db.Find(&users).Error
	return &users, err
}

func FindUserByName(name string) (*model.User, error) {
	user := &model.User{}
	err := db.Where("name = ?", name).First(&user).Error

	return user, err
}

func FindUserById(id uint) (*model.User, error) {
	user := &model.User{}
	err := db.First(&user, id).Error

	return user, err
}

func UpdatePasswordById(id uint, password string) (*model.User, error) {
	user, err := FindUserById(id)
	if err != nil {
		return nil, err
	}

	user.Password = password
	err = db.Save(&user).Error

	return user, err
}

func DeleteUserById(id uint) error {
	return db.Delete(&model.User{}, id).Error
}

func ExistsUserByName(name string) (bool, error) {
	var users []model.User

	err := db.Where("name = ?", name).Find(&users).Error
	return len(users) >= 1, err
}

func CheckLogin(name string, password string) (bool, *model.User) {
	user, err := FindUserByName(name)
	if err != nil {
		return false, nil
	}

	return user.Password == password, user
}
