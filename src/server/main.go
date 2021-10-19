/*
    This file is the main routine of the DevOps Overflow web server.
    This file creates a web server listening on a configured port when run locally.
    When installed on Lambda this creates the Lambda main routine to listen for
    incoming requests in AWS cloud from API gateway.

    In both cases the gin-gonic routes process the incoming Http requests for the
    microservice.
*/
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
    router.GET("/_next/static/chunks/pages/:page", resources.GetPageFile)
    router.GET("/_next/static/:level2/:level3", resources.GetLevel3File)
    router.GET("/:level1", resources.GetLevel1File)
    router.GET("/", resources.GetIndexFile)

    if inLambda() {
        log.Fatal(gateway.ListenAndServe(":8080", router))
    } else {
        log.Fatal(http.ListenAndServe(":8080", router))
    }
}

