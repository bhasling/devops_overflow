/*
    This file is the route (controller) for /users.
    This manages user objects for the login accounts for DevOps Overflow
*/
package resources

import (
    "net/http"
	"github.com/gin-gonic/gin"
    "api/internal/services"
)

// getUsers responds with the list of all users as JSON.
func GetUsers(c *gin.Context) {
    p, _ := c.Get("serviceProvider")
    var serviceProvider *services.ServiceProvider
    serviceProvider = p.(*services.ServiceProvider)
    userService := serviceProvider.GetUserService()
    users, _ := userService.GetAllUsers()
    c.IndentedJSON(http.StatusOK, users)
}

// GetUsersById locates the user whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func GetUsersById(c *gin.Context) {
    id := c.Param("id")
    p, _ := c.Get("serviceProvider")
    var serviceProvider *services.ServiceProvider
    serviceProvider = p.(*services.ServiceProvider)
    userService := serviceProvider.GetUserService()
    user, _ := userService.GetUserById(id)
    if (user == nil) {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
    } else {
        c.IndentedJSON(http.StatusOK, user)
    }
}


