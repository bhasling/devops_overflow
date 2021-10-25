/*
	This file is a unit test for the route (controller) for /users
	This uses test helper methods found in resources_test_helper_test.go
*/
package resources

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"main/internal/services"
	"net/http"
	"bytes"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"errors"
)

func TestBlankUserId(t *testing.T) {
	// Mock context for test
	var config = services.NewConfig()
	var serviceProvider = services.NewServiceProvider(config)
	mockFileService := CreateMockfileService(serviceProvider)
	mockFileService.AddGetFolderResult([]string {}, nil)
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	buf := bytes.NewBuffer([]byte (`{"user_id":"", "password":"pw"}`))
	c.Request, _ = http.NewRequest("POST", "/", buf)
	c.Set("serviceProvider", serviceProvider)

	// Call route with request
	PostSignup(c)

	// Verify result
	b, _ := ioutil.ReadAll(w.Body)
	assert.Contains(t, string(b), "User Id is required")
}

func TestUserAlreadyexists(t *testing.T) {
	// Mock context for test
	var config = services.NewConfig()
	var serviceProvider = services.NewServiceProvider(config)
	mockFileService := CreateMockfileService(serviceProvider)
	mockFileService.AddGetFolderResult([]string {}, nil)
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	buf := bytes.NewBuffer([]byte (`{"user_id":"user1", "password":"pw"}`))
	c.Request, _ = http.NewRequest("POST", "/", buf)
	c.Set("serviceProvider", serviceProvider)

	// Call route with request
	PostSignup(c)

	// Verify result
	b, _ := ioutil.ReadAll(w.Body)
	assert.Contains(t, string(b), "already exists")
}

func TestPostSignupUserSuccess(t *testing.T) {
	// Mock context for test
	var config = services.NewConfig()
	var serviceProvider = services.NewServiceProvider(config)
	mockFileService := CreateMockfileService(serviceProvider)
	mockFileService.AddGetFolderResult([]string {}, nil)
	mockFileService.AddGetFileResult("", errors.New("No such key"))
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	buf := bytes.NewBuffer([]byte (`{"user_id":"user2", "password":"pw"}`))
	c.Request, _ = http.NewRequest("POST", "/", buf)
	c.Set("serviceProvider", serviceProvider)

	// Call route with request
	PostSignup(c)

	// Verify result
	b, _ := ioutil.ReadAll(w.Body)
	assert.Contains(t, string(b), "user2")
}
