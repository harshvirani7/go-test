package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Ensure the raw URL path is captured
	r.UseRawPath = true

	// Add the middleware
	r.Use(removePathParamMiddleware())

	r.GET("/api/v3/videosearch/events/groups/:groupId/aggregates/value/:valueId", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("Original path: %s\nNew path: %s", c.Request.URL.Path, c.GetString("newPath")))
	})

	r.GET("/api/v3/videosearch/events/groups/:groupId/aggregates/var1/:var1Id/var2/:var2Id", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("Original path: %s\nNew path: %s", c.Request.URL.Path, c.GetString("newPath")))
	})

	// Start the server
	r.Run(":8080")
}

// Original function from VideoSearch
func removePathParam(c *gin.Context) string {
	var newPath string
	newPath = c.Request.URL.Path
	for _, paramValue := range c.Params {
		newPath = strings.Replace(newPath, paramValue.Value, "*", -1)
	}
	return newPath
}

// Middleware to use removePathParamOriginal
func removePathParamMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		newPath := removePathParam(c)
		c.Set("newPath", newPath)
		c.Next()
	}
}

// Original path: /api/v3/videosearch/events/groups/3/C-L1-28/aggregates/value/1234
// New path: /api/v3/videosearch/events/groups/*/aggregates/value/*

// Original path: /api/v3/videosearch/events/groups/3 C-L1-28/aggregates/value/12/34
// New path: /api/v3/videosearch/events/groups/*/aggregates/value/*
