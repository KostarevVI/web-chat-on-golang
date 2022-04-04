package dbstorage_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"web-chat-on-golang/dbstorage"
	"web-chat-on-golang/internal/app/model"
)

func TestUserDataRepo_Create(t *testing.T) {
	d, teardown := dbstorage.TestBD(t, databaseURL)
	defer teardown("user_data")

	u, err := d.UserData().Create(&model.UserData{
		Email:        "user@gmail.com",
		PasswordHash: "123",
		FirstName:    "John",
		LastName:     "Doe",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserDataRepo_FindUser(t *testing.T) {
	d, teardown := dbstorage.TestBD(t, databaseURL)
	defer teardown("user_data")

	email := "user@gmail.com"
	_, err := d.UserData().FindUser(email)
	assert.Error(t, err)

	d.UserData().Create(&model.UserData{
		Email:        "user@gmail.com",
		PasswordHash: "123",
		FirstName:    "John",
		LastName:     "Doe",
	})
	u, err := d.UserData().FindUser(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
