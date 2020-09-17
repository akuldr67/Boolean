package control

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	group1 := r.Group("/")
	{
		group1.GET("", getAllBooleans)
		group1.POST("", createBoolean)
	}

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	// })

	return r
}
