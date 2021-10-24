/*
	This file is a unit test for the route /issues
	This uses test helper methods found in resources_test_helper_test.go
*/
package resources

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"io/ioutil"
	"main/internal/services"
	"encoding/json"
)

func TestShouldReadAllIsses(t *testing.T) {
	// Mock up context
	var config = services.NewConfig()
	var serviceProvider = services.NewServiceProvider(config)
	mockFileService := CreateMockfileService(serviceProvider)
	mockFileService.AddGetFolderResult([]string {"1", "2"}, nil)
	mockFileService.AddGetFileResult(`{"issue_id":"1", "title":"my title1"}`, nil)
	mockFileService.AddGetFileResult(`{"issue_id":"2", "title":"my title2"}`, nil)
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("serviceProvider", serviceProvider)

	// Make request to route
	GetIssues(c)

	// Verify response returned by the http request
	b, _ := ioutil.ReadAll(w.Body)
	ExpectEquals(t, w.Code, 200)
	ExpectNotEquals(t, b, nil)
	var issues []services.Issue
	json.Unmarshal(b, &issues)
	ExpectEquals(t, len(issues), 2)
}

func TestShouldReadOneIssue(t *testing.T) {
	// Mock up context
	var config = services.NewConfig()
	var serviceProvider = services.NewServiceProvider(config)
	mockFileService := CreateMockfileService(serviceProvider)
	mockFileService.AddGetFolderResult([]string {}, nil)
	mockFileService.AddGetFolderResult([]string {"1", "2"}, nil)
	mockFileService.AddGetFileResult(`{"issue_id":"1", "title":"my title1"}`, nil)
	mockFileService.AddGetFileResult(`{"issue_id":"2", "title":"my title2"}`, nil)
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("serviceProvider", serviceProvider)
	c.Params = []gin.Param{gin.Param{Key: "id", Value: "1"}}

	// Make request to route
	GetIssueById(c)

	// Verify response returned by the http request
	ExpectEquals(t, w.Code, 200)
	b, _ := ioutil.ReadAll(w.Body)
	var issue services.Issue
	json.Unmarshal(b, &issue)
	ExpectEquals(t, issue.Title, "my title1")
}
