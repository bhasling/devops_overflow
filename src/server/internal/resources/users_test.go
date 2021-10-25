/*
	This file is a unit test for the route (controller) for /users
	This uses test helper methods found in resources_test_helper_test.go
*/
package resources

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"io/ioutil"
	"encoding/json"
	"main/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestUserHappyPath(t *testing.T) {
	// Mock context for test
	var config = services.NewConfig()
	var serviceProvider = services.NewServiceProvider(config)
	mockFileService := CreateMockfileService(serviceProvider)
	mockFileService.AddGetFolderResult([]string {}, nil)
	mockFileService.AddGetFolderResult([]string {"1", "2"}, nil)
	mockFileService.AddGetFileResult(`{"user_id":"1", "password":"user1pw"}`, nil)
	mockFileService.AddGetFileResult(`{"user_id":"2", "password":"user2pw"}`, nil)
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("serviceProvider", serviceProvider)
	c.Params = []gin.Param{gin.Param{Key: "id", Value: "user1"}}

	// Call route with request
	GetUsersById(c)

	// Verify response
	assert.Equal(t, 200, w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var u services.User
	json.Unmarshal(b, &u)
	assert.Equal(t, "user1pw", u.Password)
}
