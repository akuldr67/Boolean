package control

import (
	"github.com/akuldr67/Boolean/config"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	var boolRepo RepoInterface
	boolRepo = NewRepo(config.DB)

	ctrl := NewController(boolRepo)

	group1 := r.Group("/")
	{
		group1.GET("", ctrl.GetAllBooleans)
		group1.POST("", ctrl.CreateBoolean)
		group1.GET(":id", ctrl.GetBooleanByID)
		group1.PATCH(":id", ctrl.UpdateBoolean)
		group1.DELETE(":id", ctrl.DeleteBoolean)
	}

	return r
}
