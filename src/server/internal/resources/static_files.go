/*
	This file is a resource (controller) to support the root paths to return
	the nextjs (react) static UI files to the browser.
*/
package resources

import (
	"github.com/gin-gonic/gin"
    "api/internal/services"
	"fmt"
)

// Returns the index.html file in response to GET /
func GetIndexFile(c *gin.Context) {
    p, _ := c.Get("serviceProvider")
    var serviceProvider *services.ServiceProvider
    serviceProvider = p.(*services.ServiceProvider)
	htmlFolder := serviceProvider.GetConfig().StaticFolder
	path := fmt.Sprintf("%s/%s", htmlFolder, "index.html")
	c.File(path)
}

// Returns any files in response to GET /:level1
func GetLevel1File(c *gin.Context) {
    p, _ := c.Get("serviceProvider")
    var serviceProvider *services.ServiceProvider
    serviceProvider = p.(*services.ServiceProvider)
	fileName := c.Param("level1")
	htmlFolder := serviceProvider.GetConfig().StaticFolder
	path := fmt.Sprintf("%s/%s", htmlFolder, fileName)
	c.File(path)
}

// Returns any files in response to GET /:level1/:level2/pages/:page
func GetPageFile(c *gin.Context) {
    p, _ := c.Get("serviceProvider")
    var serviceProvider *services.ServiceProvider
    serviceProvider = p.(*services.ServiceProvider)
	fileName := c.Param("page")
	htmlFolder := serviceProvider.GetConfig().StaticFolder
	path := fmt.Sprintf("%s/_next/static/chunks/pages/%s", htmlFolder, fileName)
	c.File(path)
}

// Returns any files in response to GET /:level1/:level2/:level3
func GetLevel3File(c *gin.Context) {
    p, _ := c.Get("serviceProvider")
    var serviceProvider *services.ServiceProvider
    serviceProvider = p.(*services.ServiceProvider)
	level2 := c.Param("level2")
	level3 := c.Param("level3")
	htmlFolder := serviceProvider.GetConfig().StaticFolder
	path := fmt.Sprintf("%s/_next/static/%s/%s", htmlFolder, level2, level3)
	c.File(path)
}

