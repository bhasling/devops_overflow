/*
	This file is a unit test for the route (controller) for /users
*/
package resources

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"io/ioutil"
	"encoding/json"
	"api/internal/services"
)

func TestGetOneUser(t *testing.T) {
	var config = services.NewConfig()
	var providerService = services.NewServiceProvider(config)
   
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("serviceProvider", providerService)
	c.Params = []gin.Param{gin.Param{Key: "id", Value: "user1"}}
	GetUsersById(c)

	if w.Code == 200 {
		b, _ := ioutil.ReadAll(w.Body)
		var u services.User
		json.Unmarshal(b, &u)
		var expected = "user1pw"
		if u.Password != expected {
			t.Errorf("Expected FirstName %s, got %s in body %s", expected, u.Password, string(b))
		}
	} else {
		b, _ := ioutil.ReadAll(w.Body)
		t.Errorf("Expected 200 status code, got %d and body %s", w.Code, string(b))
	}
}
