package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// sample VideoSearch APIs test
	// Route to capture the groupId with a manual split
	r.GET("/api/v3/videosearch/events/groups/:groupId/aggregates", func(c *gin.Context) {
		newPath := removePathParam(c)
		// newPath := removePathParamOriginal(c)
		c.String(200, fmt.Sprintf("Original path: %s\nNew path: %s", c.Request.URL.Path, newPath))
	})

	// Start the server
	r.Run(":8080")
}

// Original function from VideoSearch
func removePathParam(c *gin.Context) string {
	var newPath string
	for _, str := range strings.Split(c.Request.URL.Path, "/") {
		found := false
		for _, paramValue := range c.Params {
			if str == paramValue.Value {
				newPath = newPath + "*/"
				found = true
				break
			}
		}
		if !found {
			newPath = newPath + str + "/"
		}
	}
	return newPath
}
