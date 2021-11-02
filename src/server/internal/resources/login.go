/*
    This file is the route (controller) for /login.
    This supports a login request from the UI.
*/
package resources

import (
    "net/http"
	"github.com/gin-gonic/gin"
    "main/internal/services"
    "fmt"
    // "github.com/golang-jwt/jwt"  This library requires GO 1.22
    "github.com/dgrijalva/jwt-go"
    "time"
    "log"
)

// Post a message to authenticate and login a user
// If successful return json {"user_id": "username"} and a session cookie with a JWT token.
func PostLogin(c *gin.Context) {
    var postedUser services.User
    err := c.BindJSON(&postedUser);
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid login.")})
        return
    }

    p, _ := c.Get("serviceProvider")
    var serviceProvider *services.ServiceProvider
    serviceProvider = p.(*services.ServiceProvider)
    userService := serviceProvider.GetUserService()

    // Check if this user already exists
    existingUser, _ := userService.GetUserById(postedUser.UserId)
    if existingUser == nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid login.")})
        return
    }

    // If user does not exist then create the user with the user ID and password
    if postedUser.Password != existingUser.Password {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid login.")})
        return
    }

    jwtToken, err, jwtExpirationTime := createJwtToken(existingUser.UserId)
    if err != nil {
        log.Println(err)
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Unable to create authentication token.")})
        return
    }

    // Return the user object as the result
    c.IndentedJSON(http.StatusOK, existingUser)

    // Return an Authentication cookie with the JWT token
    http.SetCookie(c.Writer, &http.Cookie {
        Name: "Authentication",
        Value: fmt.Sprintf("Bearer %s", jwtToken),
        Expires: jwtExpirationTime,
    })
}

type JwtClaims struct {
    jwt.StandardClaims
}
func createJwtToken(userId string) (string, error, time.Time) {
    jwtSecretString := []byte("devops secret password 43453")
    jwtExpirationTime := time.Now().Add(60 * time.Minute)
    claims := &JwtClaims {
        StandardClaims: jwt.StandardClaims {
            Subject: userId,
            ExpiresAt: jwtExpirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtSecretString)
    if err != nil {
        return "", err, jwtExpirationTime
    } else {
        return tokenString, nil, jwtExpirationTime
    }
}