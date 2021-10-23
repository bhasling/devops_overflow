/*
    This file is the route (controller) for DevOps issues.
    An issue is a DevOps issue submitted by a DevOps logged in user.
    An issue may have a list oassociated answers contained with an issue.
*/
package resources

import (
    "net/http"
	"github.com/gin-gonic/gin"
    "main/internal/services"
    "log"
)

// Returns a list of issues in response to GET /issues
func GetIssues(c *gin.Context) {
    log.Println("Called resource GetIssues")
    p, _ := c.Get("serviceProvider")
    var serviceProvider *services.ServiceProvider
    serviceProvider = p.(*services.ServiceProvider)
    issueService := serviceProvider.GetIssueService()
    issues, _ := issueService.GetAllIssues()
    log.Printf("Returned %d records", len(issues))
    c.IndentedJSON(http.StatusOK, issues)
}

// Returns a single issue in response to GET /issues/:issueId
func GetIssueById(c *gin.Context) {
    id := c.Param("id")
    p, _ := c.Get("serviceProvider")
    var serviceProvider *services.ServiceProvider
    serviceProvider = p.(*services.ServiceProvider)
    issueService := serviceProvider.GetIssueService()

    issue, _ := issueService.GetIssueById(id)
    if (issue == nil) {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "issue not found"})
    } else {
        c.IndentedJSON(http.StatusOK, issue)
    }
}

// Creates a new issue in response to POST /issues
func PostIssue(c *gin.Context) {
    var postedIssue services.Issue
    err := c.BindJSON(&postedIssue);
    if err != nil {
        return
    }
    p, _ := c.Get("serviceProvider")
    var serviceProvider *services.ServiceProvider
    serviceProvider = p.(*services.ServiceProvider)
    issueService := serviceProvider.GetIssueService()
    newIssue, _ := issueService.CreateIssue()
    newIssue.Title = postedIssue.Title
    newIssue.Description = postedIssue.Description
    newIssue.Product = postedIssue.Product
    err = issueService.SaveIssue(newIssue)
    if (err != nil) {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to create new issue."})       
    }
    c.IndentedJSON(http.StatusOK, gin.H{"message": "Submitted issue."})
}

// Updates an existing issue in response to PUT /issues
func PutIssue(c *gin.Context) {
    var newIssue services.Issue

    if err := c.BindJSON(&newIssue); err != nil {
        return
    }

    // Now newIssue contains the posted JSON information deserialized
}

// Deletes an existing issue by ID in reponse to DELETE /ssues/:issueId
func DeleteIssueById(c *gin.Context) {

}

