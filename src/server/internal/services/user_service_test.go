/*
	Unit test of the user service.
	This uses test helper methods found in services_test_helper_test.go
*/
package services

import (
	"testing"
	"github.com/stretchr/testify/assert"

)
func TestUserHappyPath(t *testing.T) {
	var config = NewConfig()
	var serviceProvider = NewServiceProvider(config)
	mockFileService := CreateMockfileService(serviceProvider)
	var userService = serviceProvider.GetUserService()
	mockFileService.AddGetFolderResult([]string {}, nil)
	mockFileService.AddGetFolderResult([]string {"123"}, nil)
	mockFileService.AddGetFileResult(`{"user_id":"1", "password":"my password"}`, nil)
	mockFileService.AddGetFileResult(`{"user_id":"1", "password":"my password"}`, nil)

	// Get initial list of users
	initialUsers, _ := userService.GetAllUsers()

	// Test Create user
	user,_ := userService.CreateUser("1")
	assert.Equal(t, "1", user.UserId)

	// Test add password to user and save it
	user.Password = "my password"
	err := userService.SaveUser(user)
	assert.Equal(t, nil, err)
	users, _ := userService.GetAllUsers()
	assert.Equal(t, len(initialUsers) + 1, len(users))

	// Test read back user
	readBackUser, _ := userService.GetUserById(user.UserId)
	assert.Equal(t, user.Password, readBackUser.Password)

	// Delete new user
	err = userService.DeleteUserById("1")
	assert.Equal(t, nil, err)
}
