package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/tiagoncardoso/golang-api/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log/slog"
	"testing"
)

func dbConnectUser() *gorm.DB {
	dsn := "file::memory:"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}

	return db
}

func createUser(db *gorm.DB, user *entity.User) (*User, error) {

	userDB := NewUser(db)
	err := userDB.Create(user)
	if err != nil {
		slog.Error("err creating user", "msg", err)
		return userDB, err
	}

	return userDB, nil
}

func TestUser_Create(t *testing.T) {
	db := dbConnectUser()
	user, _ := entity.NewUser("John Doe", "john@doe.com", "123456")

	_, err := createUser(db, user)
	if err != nil {
		slog.Error("err connect", "msg", err)
	}
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)
}

func TestFindUserByEmail(t *testing.T) {
	db := dbConnectUser()

	user, _ := entity.NewUser("John Doe", "john@doe", "123456")
	userDb, err := createUser(db, user)

	assert.Nil(t, err)

	userFound, err := userDb.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)
}
