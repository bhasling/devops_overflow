/*
    This file is the route (controller) for /signup.
    This supports a sign up request from the UI.
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

    // Check if this user already exists
    existingUser, _ := userService.GetUserById(postedUser.UserId)
    if existingUser != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("User '%s' already exists.", postedUser.UserId)})
        return
    }

    // If user does not exist then create the user with the user ID and password
    user, _ := userService.CreateUser(postedUser.UserId)
    user.Password = postedUser.Password
    err = userService.SaveUser(user)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Unable to create user '%s'.", postedUser.UserId)})
    }

    // Return the user object as the result
    c.IndentedJSON(http.StatusOK, user)
}
