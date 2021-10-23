/*
	This file implements UserService. This provides business logic operations to support
	the Users that can login to the application.
*/
package services

import (
	"fmt"
	"encoding/json"
	"log"
)

type UserServiceImpl struct {
	persistedFileService		PersistedFileService
}

type User struct {
    UserId            			string          `json:"user_id"`
	Password					string			`json:"password"`
}
func NewUserService(persistedFileService PersistedFileService) * UserServiceImpl {
	return &UserServiceImpl{
		persistedFileService: persistedFileService,
	}
}

func (s *UserServiceImpl ) GetAllUsers() ([]User, error) {
	result := make([]User, 0)
	keys, err := s.persistedFileService.GetFolders("users/")
	if err != nil {
		log.Println(err)
		return result, err
	}
	for _, key := range keys {
		content, _ := s.persistedFileService.GetFile(key)
		user := &User{}
		json.Unmarshal([]byte(content), &user)
		result = append(result, *user)
	}
	return result, nil
}

func (s *UserServiceImpl ) CreateUser(userId string) (*User, error) {
	result := User {UserId: userId}
	return &result, nil
}

func (s *UserServiceImpl ) SaveUser(user *User) error {
	content, _ := json.Marshal(user)
	var key = fmt.Sprintf("users/%s", user.UserId)
	err := s.persistedFileService.WriteFile(key, string(content))
	if err != nil {
		log.Printf("Unable to save user '%s': %v", key, err)
		return err
	}
	return nil
}
func (s *UserServiceImpl ) GetUserById(userId string) (*User, error) {
	var key = fmt.Sprintf("users/%s", userId)
	content, err := s.persistedFileService.GetFile(key)
	if (err != nil) {
		log.Println(err)
		return nil, err
	}
	user := &User{}
	json.Unmarshal([]byte(content), &user)
	return user, nil
}

func (s *UserServiceImpl ) DeleteUserById(userId string) error {
	var key = fmt.Sprintf("users/%s", userId)
	err := s.persistedFileService.DeleteFile(key)
	if (err != nil) {
		log.Println(err)
		return err
	}
	return nil
}

