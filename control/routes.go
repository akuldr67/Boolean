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
		group1.GET(":id", getBooleanByID)
		group1.PATCH(":id", updateBoolean)
		group1.DELETE(":id", deleteBoolean)
	}

	return r
}
