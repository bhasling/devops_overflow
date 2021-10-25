/*
    This file is the route (controller) for /users.
    This manages user objects for the login accounts for DevOps Overflow
*/
package resources

import (
    "net/http"
	"github.com/gin-gonic/gin"
    "main/internal/services"
    "fmt"
)

// Post a message to sign up a new user 
func PostSignup(c *gin.Context) {
    var postedUser services.User
    err := c.BindJSON(&postedUser);
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to sign up."})       
        return
    }
    if postedUser.UserId == "" {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User Id is required."})       
        return
    }
    if postedUser.Password == "" {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Password is required."})       
        return
    }

    p, _ := c.Get("serviceProvider")
    var serviceProvider *services.ServiceProvider
    serviceProvider = p.(*services.ServiceProvider)
    userService := serviceProvider.GetUserService()
    existingUser, _ := userService.GetUserById(postedUser.UserId)
    if existingUser != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("User '%s' already exists.", postedUser.UserId)})
        return
    }
    user, _ := userService.CreateUser(postedUser.UserId)
    user.Password = postedUser.Password
    err = userService.SaveUser(user)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Unable to create user '%s'.", postedUser.UserId)})
    }
    c.IndentedJSON(http.StatusOK, user)
}
