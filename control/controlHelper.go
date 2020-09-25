package control

import (
	"github.com/akuldr67/Boolean/models"

	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
)

type Repo struct {
	DB *gorm.DB
}

type RepoInterface interface {
	GetAllBooleansHelper(bools *[]models.Boolean) (err error)
	CreateBooleanHelper(newBool *models.Boolean) (err error)
	GetBooleanByIDHelper(boolean *models.Boolean, id string) error
	UpdateBooleanHelper(oldBool *models.Boolean, newBool models.Boolean) error
	DeleteBooleanHelper(boolean *models.Boolean) error
}

func NewRepo(db *gorm.DB) Repo {
	return Repo{
		DB: db,
	}
}

func (r Repo) GetAllBooleansHelper(bools *[]models.Boolean) (err error) {
	if err = r.DB.Find(bools).Error; err != nil {
		return err
	}
	return nil
}

func (r Repo) CreateBooleanHelper(newBool *models.Boolean) (err error) {
	if err = r.DB.Create(newBool).Error; err != nil {
		return err
	}
	return nil
}

func (r Repo) GetBooleanByIDHelper(boolean *models.Boolean, id string) error {
	if err := r.DB.Where("id = ?", id).First(boolean).Error; err != nil {
		return err
	}
	return nil
}

func (r Repo) UpdateBooleanHelper(oldBool *models.Boolean, newBool models.Boolean) error {
	r.DB.Model(oldBool).Update(newBool)
	return nil
}

func (r Repo) DeleteBooleanHelper(boolean *models.Boolean) error {
	r.DB.Delete(boolean)
	return nil
}
