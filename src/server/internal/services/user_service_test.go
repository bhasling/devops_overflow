/*
	Unit test or the user service.
*/
package services

import (
	"testing"
)
func TestUserHappyPath(t *testing.T) {
	var config = NewConfig()
	var serviceProvider = NewServiceProvider(config)
	var userService = serviceProvider.GetUserService()
	users, _ := userService.GetAllUsers()
	userCount := len(users)
	user,_ := userService.CreateUser("1")
	if (user.UserId != "1") {
		t.Errorf("Expected userId 1")
	}
	user.Password = "my password"
	err := userService.SaveUser(user)
	if (err != nil) {
		t.Errorf("Save user failed %s", err.Error())
	}
	user, _ = userService.GetUserById("1")
	if (user.Password != "my password") {
		t.Errorf("Expected the password when getting user")
	}
	users, _ = userService.GetAllUsers()
	if (len(users) != userCount + 1) {
		t.Errorf("Expected %d users got %d.", userCount + 1, len(users))
	}
	err = userService.DeleteUserById("1")
	if (err != nil) {
		t.Errorf("Unable to delete user")
	}
}
