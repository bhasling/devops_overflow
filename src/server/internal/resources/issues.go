/*
    This file is the route (controller) for DevOps issues.
    An issue is a DevOps issue submitted by a DevOps logged in user.
    An issue may have a list oassociated answers contained with an issue.
*/
package resources

import (
    "net/http"
	"github.com/gin-gonic/gin"
    "api/internal/services"
)

// Returns a list of issues in response to GET /issues
func GetIssues(c *gin.Context) {
    p, _ := c.Get("serviceProvider")
    var serviceProvider *services.ServiceProvider
    serviceProvider = p.(*services.ServiceProvider)
    issueService := serviceProvider.GetIssueService()
    issues, _ := issueService.GetAllIssues()
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
    var newIssue services.Issue

    if err := c.BindJSON(&newIssue); err != nil {
        return
    }

    // Now newIssue contains the posted JSON information deserialized
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

