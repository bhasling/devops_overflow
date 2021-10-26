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
)

func TestLoginWithIllegalPassword(t *testing.T) {
	// Mock context for test
	var config = services.NewConfig()
	var serviceProvider = services.NewServiceProvider(config)
	mockFileService := CreateMockfileService(serviceProvider)
	mockFileService.AddGetFileResult(`{"user_id":"", "password":"password123"}`, nil)
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	buf := bytes.NewBuffer([]byte (`{"user_id":"", "password":"badpassword"}`))
	c.Request, _ = http.NewRequest("POST", "/", buf)
	c.Set("serviceProvider", serviceProvider)

	// Call route with request
	PostLogin(c)

	// Verify result
	b, _ := ioutil.ReadAll(w.Body)
	assert.Contains(t, string(b), "Invalid login")
}


func TestLoginWithLegalPassword(t *testing.T) {
	// Mock context for test
	var config = services.NewConfig()
	var serviceProvider = services.NewServiceProvider(config)
	mockFileService := CreateMockfileService(serviceProvider)
	mockFileService.AddGetFileResult(`{"user_id":"legal_user", "password":"password123"}`, nil)
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	buf := bytes.NewBuffer([]byte (`{"user_id":"legal_user", "password":"password123"}`))
	c.Request, _ = http.NewRequest("POST", "/", buf)
	c.Set("serviceProvider", serviceProvider)

	// Call route with request
	PostLogin(c)

	// Verify result
	b, _ := ioutil.ReadAll(w.Body)
	assert.Contains(t, string(b), "legal_user")
	assert.Contains(t, w.Header()["Set-Cookie"][0], "Bearer ")
}
