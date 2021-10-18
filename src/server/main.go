package main

import (
    "github.com/gin-gonic/gin"
    "api/internal/resources"
    "api/internal/services"
    "net/http"
    "log"
    "github.com/apex/gateway"
    "os"
)

func inLambda() bool {
    if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
       return true
    }
    return false
 }

 var config = services.NewConfig()
 var providerService = services.NewServiceProvider(config)

func ProviderServiceMiddleware(c *gin.Context) {
    c.Set("serviceProvider", providerService)
    c.Next()
}
func main() {
    router := gin.Default()
    router.Use(ProviderServiceMiddleware)
    router.GET("/issues", resources.GetIssues)
	router.GET("/issues/:id", resources.GetIssueById)
    router.POST("/issues", resources.PostIssue)
    router.PUT("/issues", resources.PutIssue)
	router.DELETE("/issues/:id", resources.DeleteIssueById)
    router.GET("/users", resources.GetUsers)
	router.GET("/users/:id", resources.GetUsersById)

    if inLambda() {
        log.Fatal(gateway.ListenAndServe(":8080", router))
    } else {
        log.Fatal(http.ListenAndServe(":8080", router))
    }
}

