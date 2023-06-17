package router

import (
	"github.com/gin-gonic/gin"
	"github.com/laracarvalho/goboards/handler"

	docs "github.com/laracarvalho/goboards/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start() *gin.Engine {
	handler.InitHandler()
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Project reachable.",
		})
	})

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/listings", handler.ListListingsHandler)
		v1.POST("/listing", handler.CreateListingHandler)
		v1.DELETE("/listing", handler.DeleteListingHandler)
		v1.GET("/listing", handler.ShowListingsHandler)
		v1.PUT("/listing", handler.UpdateListingsHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
