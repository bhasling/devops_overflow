/*
    This file is the main routine of the DevOps Overflow web server.
    This file creates a web server listening on port 8080 when run locally.
    When installed on Lambda this creates the Lambda main routine to listen for
    incoming requests in AWS cloud from API gateway.

    In both cases the gin-gonic routes process the incoming Http requests.
*/
package main

import (
    "github.com/gin-gonic/gin"
    "main/internal/resources"
    "main/internal/services"
    "net/http"
    "log"
    "github.com/apex/gateway"
    "os"
)

func inLambda() bool {
    lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT")
    if lambdaTaskRoot != "" {
        return true
    } else {
        return false
    }
 }

func main() {
    config := services.NewConfig()
    err := config.LoadConfig("./config.yaml")
    if err != nil {
        log.Println(err)
        log.Fatal("Unable to read configuration file.")
    }
    if inLambda() {
        config.StaticFolder = "."
    } else {
        _, err = os.Stat("./html")
        if err == nil {
            // Use the static files put here by docker build
            config.StaticFolder = "./html"
        } else {
            // Use the static file from local development
            config.StaticFolder = "../ui/out"
        }
    }
    serviceProvider := services.NewServiceProvider(config)
    router := gin.Default()
    router.Use(func(c *gin.Context) {
        c.Set("serviceProvider", serviceProvider)
        })

    router.GET("/api/issues", resources.GetIssues)
	router.GET("/api/issues/:id", resources.GetIssueById)
    router.POST("/api/issues", resources.PostIssue)
    router.PUT("/api/issues", resources.PutIssue)
	router.DELETE("/api/issues/:id", resources.DeleteIssueById)
    router.GET("/api/users", resources.GetUsers)
	router.GET("/api/users/:id", resources.GetUsersById)
    router.POST("/api/signup", resources.PostSignup)
    router.POST("/api/login", resources.PostLogin)
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

