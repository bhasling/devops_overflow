/*
	This file is a unit test for the route /issues
*/
package resources

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"io/ioutil"
	"encoding/json"
	"encoding/json"
	"api/internal/services"
)
func TestGetAll(t *testing.T) {
	var config = services.NewConfig()
	var providerService = services.NewServiceProvider(config)

	gin.SetMode(gin.TestMode)
	c.Set("serviceProvider", providerService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	GetIssues(c)
	b, _ := ioutil.ReadAll(w.Body)
	if w.Code != 200 {
		t.Error(w.Code, string(b))
	}
}

func TestGetOne(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{gin.Param{Key: "id", Value: "1"}}
	GetIssueById(c)

	if w.Code == 200 {
		b, _ := ioutil.ReadAll(w.Body)
		var a issue
		json.Unmarshal(b, &a)
		var expected = "Problem 1"
		if a.Name != expected {
			t.Errorf("Expected name %s, got %s in body %s", expected, a.Name, string(b))
		}
	} else {
		b, _ := ioutil.ReadAll(w.Body)
		t.Errorf("Expected 200 status code, got %d and body %s", w.Code, string(b))
	}
}

func TestGetOneError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{gin.Param{Key: "id", Value: "xxx"}}
	GetIssueById(c)

	if w.Code != 404 {
		b, _ := ioutil.ReadAll(w.Body)
		t.Errorf("Expected 404 status code, got %d and body %s", w.Code, string(b))
	}
}
