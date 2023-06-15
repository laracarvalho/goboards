package router

import "github.com/gin-gonic/gin"

func Start() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Project reachable.",
		})
	})

	router.Run(":8080")
}