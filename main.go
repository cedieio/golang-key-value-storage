package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var kv = make(map[string]string)

type keyVal struct {
	Key   string `json:"key" xml:"key" binding:"required"`
	Value string `json:"value" xml:"value" binding:"required"`
}

func main() {
	server := gin.Default()
	server.GET("/storage/:key", func(c *gin.Context) {
		key := c.Param("key")
		val, ok := kv[key]

		if !ok {
			c.JSON(http.StatusNotFound, gin.H{
				"response": "data not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": val,
		})
	})

	server.POST("/storage", func(c *gin.Context) {
		var jsonVal keyVal
		if err := c.ShouldBindJSON(&jsonVal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		kv[jsonVal.Key] = jsonVal.Value
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("saved %v", jsonVal.Value),
		})
	})

	server.Run()
}
