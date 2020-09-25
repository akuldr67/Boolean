package control

import (
	"net/http"

	"github.com/akuldr67/Boolean/models"
	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	repo RepoInterface
}

type ControllerInterface interface {
	GetAllBooleans(ctx *gin.Context)
	CreateBoolean(ctx *gin.Context)
	GetBooleanByID(ctx *gin.Context)
	UpdateBoolean(ctx *gin.Context)
	DeleteBoolean(ctx *gin.Context)
}

func NewController(repo RepoInterface) Controller {
	return Controller{
		repo: repo,
	}
}

func (c Controller) GetAllBooleans(ctx *gin.Context) {
	var bools []models.Boolean

	// err := config.DB.Find(&bools).Error
	err := c.repo.GetAllBooleansHelper(&bools)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, bools)
	}
}

func (c Controller) CreateBoolean(ctx *gin.Context) {
	var newBool models.Boolean
	err := ctx.ShouldBindJSON(&newBool)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	newBool.ID, _ = uuid.NewV4()
	// newBool = models.Boolean{ID: uuid, Key: newBool.Key, Value: newBool.Value}

	// err = c.repo.DB.Create(&newBool).Error
	err = c.repo.CreateBooleanHelper(&newBool)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, newBool)
	}
}

func (c Controller) GetBooleanByID(ctx *gin.Context) {
	var boolean models.Boolean

	id := ctx.Params.ByName("id")

	err := c.repo.GetBooleanByIDHelper(&boolean, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, boolean)
	}
}

func (c Controller) UpdateBoolean(ctx *gin.Context) {
	var boolean models.Boolean
	id := ctx.Params.ByName("id")

	err := c.repo.GetBooleanByIDHelper(&boolean, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	// ctx.BindJSON(&boolean)
	// c.repo.DB.Save(&boolean)

	var input models.Boolean
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	// crepo.DB.Model(&boolean).Updates(input)
	err = c.repo.UpdateBooleanHelper(&boolean, input)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, boolean)
	}
}

func (c Controller) DeleteBoolean(ctx *gin.Context) {
	var boolean models.Boolean
	id := ctx.Params.ByName("id")

	err := c.repo.GetBooleanByIDHelper(&boolean, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	// c.repo.DB.Delete(&boolean)
	err = c.repo.DeleteBooleanHelper(&boolean)

	// ctx.JSON(http.StatusOK, gin.H{"id" + id: "deleted"})
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.Writer.WriteHeader(204)
	}
}
