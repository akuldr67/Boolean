package control

import (
	"net/http"

	"github.com/akuldr67/Boolean/config"
	"github.com/akuldr67/Boolean/models"

	"github.com/gin-gonic/gin"
)

func getAllBooleans(c *gin.Context) {
	var bools []models.Boolean
	// err := models.GetAllBooleans(&bools)
	err := config.DB.Find(&bools).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, bools)
	}
}
