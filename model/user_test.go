package model

import "testing"

func TestMigration(t *testing.T) {
	db.AutoMigrate(&User{})
}

func TestUserStore(t *testing.T) {
	err := db.Create(&User{
		Name:     "konojunya",
		Password: "konojunya",
	}).Error

	if err != nil {
		t.Error(err)
	}
}

func TestFindUserById(t *testing.T) {
	user := &User{}
	user.ID = 1

	db.First(&user)
}

func TestFindUserByName(t *testing.T) {
	user := &User{}
	db.Where("name = ?", "konojunya").First(&user)
}

func TestUpdateUserById(t *testing.T) {
	user := &User{}
	db.First(&user, 1)

	user.Name = "makki"
	db.Save(&user)
}

func TestDeleteUserById(t *testing.T) {
	user := &User{}

	db.Delete(&user, 1)
}

func TestDeleteUserByName(t *testing.T) {
	user := &User{}

	db.Delete(&user, "name = ?", "konojunya")
}
