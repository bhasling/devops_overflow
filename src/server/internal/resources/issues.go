package resources

import (
    "net/http"
	"github.com/gin-gonic/gin"
    "api/internal/services"
)

// getAlbums responds with the list of all albums as JSON.
func GetIssues(c *gin.Context) {
    p, _ := c.Get("serviceProvider")
    var serviceProvider *services.ServiceProvider
    serviceProvider = p.(*services.ServiceProvider)
    issueService := serviceProvider.GetIssueService()
    issues, _ := issueService.GetAllIssues()
    c.IndentedJSON(http.StatusOK, issues)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
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

func PostIssue(c *gin.Context) {
    var newIssue services.Issue

    if err := c.BindJSON(&newIssue); err != nil {
        return
    }

    // Now newIssue contains the posted JSON information deserialized
}

func PutIssue(c *gin.Context) {
    var newIssue services.Issue

    if err := c.BindJSON(&newIssue); err != nil {
        return
    }

    // Now newIssue contains the posted JSON information deserialized
}

func DeleteIssueById(c *gin.Context) {

}

