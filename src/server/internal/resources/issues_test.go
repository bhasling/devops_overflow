/*
	This file is a unit test for the route /issues
*/
package resources

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"io/ioutil"
	"main/internal/services"
	"fmt"
)
func TestGetAll(t *testing.T) {
	var config = services.NewConfig()
	var providerService = services.NewServiceProvider(config)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("serviceProvider", providerService)
	GetIssues(c)
	b, _ := ioutil.ReadAll(w.Body)
	if w.Code != 200 {
		t.Error(w.Code, string(b))
	}
}

func TestGetOne(t *testing.T) {
	var config = services.NewConfig()
	var providerService = services.NewServiceProvider(config)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("serviceProvider", providerService)
	c.Params = []gin.Param{gin.Param{Key: "id", Value: "1"}}
	GetIssueById(c)

	fmt.Println(w.Code)
	if w.Code != 404 {
		t.Errorf("Expected 404 return code")
	}
}
