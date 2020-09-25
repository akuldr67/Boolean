package control

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_control "github.com/akuldr67/Boolean/mocks"
	"github.com/gin-gonic/gin"

	"github.com/akuldr67/Boolean/models"
	"github.com/golang/mock/gomock"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func getTestBooleans() []models.Boolean {
	id1, _ := uuid.NewV4()
	id2, _ := uuid.NewV4()
	return []models.Boolean{
		{
			ID:    id1,
			Key:   "key",
			Value: new(bool),
		}, {
			ID:    id2,
			Value: new(bool),
		},
	}
}

func initiateTest(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.Nil(t, err, "an error was not expected when opening a stub database connection")

	// config.DB, err = gorm.Open(mysql.Open("db"), &gorm.Config{})
	DB, err2 := gorm.Open("mysql", db)
	assert.Nil(t, err2, "error while opening test mysql database")
	return mock, DB
}

func TestGetAllBooleansHelper(t *testing.T) {
	mock, db := initiateTest(t)
	testBooleans := getTestBooleans()

	rows := sqlmock.NewRows([]string{"id", "key", "value"})

	for _, boolean := range testBooleans {
		rows = rows.AddRow(boolean.ID, boolean.Key, boolean.Value)
	}

	mock.ExpectQuery("SELECT * FROM `booleans`").WillReturnRows(rows)

	var allBooleans []models.Boolean
	r := NewRepo(db)
	err := r.GetAllBooleansHelper(&allBooleans)
	assert.Nil(t, err)

	assert.Equal(t, allBooleans, testBooleans)

	err = mock.ExpectationsWereMet()
	assert.Nil(t, err)
	// if err != nil {
	// 	t.Errorf("expectations not matched %s", err)
	// }
}

func TestGetAllBooleansStatusOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_control.NewMockRepoInterface(ctrl)
	var b []models.Boolean
	var err error
	err = nil
	mockRepo.EXPECT().GetAllBooleansHelper(&b).Return(err)

	controller := NewController(mockRepo)
	tr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(tr)
	controller.GetAllBooleans(c)

	// assert.Equal(t, "null", tr.Body.String())
	assert.Equal(t, http.StatusOK, tr.Code)
}

func TestGetAllBooleansStatusNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_control.NewMockRepoInterface(ctrl)
	var b []models.Boolean
	var err error
	err = fmt.Errorf("unable to find database")
	mockRepo.EXPECT().GetAllBooleansHelper(&b).Return(err)

	controller := NewController(mockRepo)
	tr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(tr)
	controller.GetAllBooleans(c)

	assert.Equal(t, http.StatusNotFound, tr.Code)
}

func TestCreateBooleansHelper(t *testing.T) {
	mock, db := initiateTest(t)
	testBooleans := getTestBooleans()

	const sqlInsert = "INSERT INTO `booleans` (`id`,`key`,`value`) VALUES (?,?,?)"

	for _, boolean := range testBooleans {
		mock.ExpectBegin()
		mock.ExpectExec(sqlInsert).
			WithArgs(boolean.ID.String(), boolean.Key, boolean.Value).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		r := NewRepo(db)
		err := r.CreateBooleanHelper(&boolean)
		assert.Nil(t, err)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	}
}

func TestCreateBooleansStatusOK(t *testing.T) {
	testBooleans := getTestBooleans()
	ctrl := gomock.NewController(t)
	mockRepo := mock_control.NewMockRepoInterface(ctrl)
	var err error

	for _, boolean := range testBooleans {
		err = nil
		mockRepo.EXPECT().CreateBooleanHelper(gomock.Any()).Return(err)
		jsonRequest, err2 := json.Marshal(models.Boolean{
			Key:   boolean.Key,
			Value: boolean.Value,
		})
		assert.Nil(t, err2)

		controller := NewController(mockRepo)
		tr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(tr)

		c.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(jsonRequest))
		assert.Nil(t, err)

		c.Request.Header.Set("Content-Type", "application/json")
		controller.CreateBoolean(c)
		assert.Equal(t, http.StatusOK, tr.Code)
	}
}

func TestCreateBooleansBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_control.NewMockRepoInterface(ctrl)
	var err error
	err = nil

	mockRepo.EXPECT().CreateBooleanHelper(gomock.Any()).Return(err)
	jsonRequest, err2 := json.Marshal(models.Boolean{
		Key: "key without value",
	})
	assert.Nil(t, err2)

	controller := NewController(mockRepo)
	tr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(tr)

	c.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(jsonRequest))
	assert.Nil(t, err)

	c.Request.Header.Set("Content-Type", "application/json")
	controller.CreateBoolean(c)
	assert.Equal(t, http.StatusBadRequest, tr.Code)
}

func TestCreateBooleansStatusNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_control.NewMockRepoInterface(ctrl)
	testBooleans := getTestBooleans()
	var err error

	for _, boolean := range testBooleans {
		err = fmt.Errorf("unable to find database")
		mockRepo.EXPECT().CreateBooleanHelper(gomock.Any()).Return(err)
		jsonRequest, err2 := json.Marshal(models.Boolean{
			Key:   boolean.Key,
			Value: boolean.Value,
		})
		assert.Nil(t, err2)

		controller := NewController(mockRepo)
		tr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(tr)

		c.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(jsonRequest))
		assert.Nil(t, err)

		c.Request.Header.Set("Content-Type", "application/json")
		controller.CreateBoolean(c)
		assert.Equal(t, http.StatusNotFound, tr.Code)
	}
}

func TestGetBooleanByIDHelper(t *testing.T) {
	mock, db := initiateTest(t)

	testBooleans := getTestBooleans()

	for _, boolean := range testBooleans {
		rows := sqlmock.
			NewRows([]string{"id", "key", "value"}).
			AddRow(boolean.ID, boolean.Key, boolean.Value)

		mock.ExpectQuery("SELECT * FROM `booleans` WHERE (id = ?) ORDER BY `booleans`.`id` ASC LIMIT 1").
			WithArgs(boolean.ID).
			WillReturnRows(rows)

		var newBoolean models.Boolean
		r := NewRepo(db)
		err := r.GetBooleanByIDHelper(&newBoolean, boolean.ID.String())
		assert.Nil(t, err)
		assert.Equal(t, newBoolean.ID, boolean.ID)
		assert.Equal(t, newBoolean.Key, boolean.Key)
		assert.Equal(t, newBoolean.Value, boolean.Value)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	}
}

func TestGetBooleanByIDStatusOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_control.NewMockRepoInterface(ctrl)
	var err error
	err = nil
	mockRepo.EXPECT().GetBooleanByIDHelper(gomock.Any(), gomock.Any()).Return(err)

	controller := NewController(mockRepo)
	tr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(tr)
	controller.GetBooleanByID(c)

	assert.Equal(t, http.StatusOK, tr.Code)
}

func TestGetBooleanByIDStatusNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_control.NewMockRepoInterface(ctrl)
	var err error
	err = fmt.Errorf("unable to find database")
	mockRepo.EXPECT().GetBooleanByIDHelper(gomock.Any(), gomock.Any()).Return(err)

	controller := NewController(mockRepo)
	tr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(tr)
	controller.GetBooleanByID(c)

	assert.Equal(t, http.StatusNotFound, tr.Code)
}

func TestUpdateBooleanHelper(t *testing.T) {
	mock, db := initiateTest(t)

	testBooleans := getTestBooleans()

	for _, boolean := range testBooleans {
		var sqlUpdate string
		mock.ExpectBegin()
		if boolean.Key != "" {
			sqlUpdate = "UPDATE `booleans` SET `id` = ?, `key` = ?, `value` = ? WHERE `booleans`.`id` = ?"
			mock.ExpectExec(sqlUpdate).
				WithArgs(boolean.ID, boolean.Key, boolean.Value, boolean.ID).
				WillReturnResult(sqlmock.NewResult(1, 1))
		} else {
			sqlUpdate = "UPDATE `booleans` SET `id` = ?, `value` = ? WHERE `booleans`.`id` = ?"
			mock.ExpectExec(sqlUpdate).
				WithArgs(boolean.ID, boolean.Value, boolean.ID).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}
		mock.ExpectCommit()

		r := NewRepo(db)
		err := r.UpdateBooleanHelper(&boolean, boolean)
		assert.Nil(t, err)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	}
}

func TestUpdateBooleanStatusOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_control.NewMockRepoInterface(ctrl)
	var err error
	testBooleans := getTestBooleans()

	for _, boolean := range testBooleans {
		err = nil
		mockRepo.EXPECT().GetBooleanByIDHelper(gomock.Any(), gomock.Any()).Return(err)
		mockRepo.EXPECT().UpdateBooleanHelper(gomock.Any(), gomock.Any()).Return(err)

		jsonRequest, err2 := json.Marshal(models.Boolean{
			Key:   boolean.Key,
			Value: boolean.Value,
		})
		assert.Nil(t, err2)

		controller := NewController(mockRepo)
		tr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(tr)

		c.Request, err = http.NewRequest("PATCH", "/", bytes.NewBuffer(jsonRequest))
		assert.Nil(t, err)

		c.Request.Header.Set("Content-Type", "application/json")
		c.Params = gin.Params{
			{Key: "id", Value: boolean.ID.String()},
		}
		controller.UpdateBoolean(c)

		assert.Equal(t, http.StatusOK, tr.Code)
	}
}

func TestUpdateBooleanStatusNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_control.NewMockRepoInterface(ctrl)
	var err, errID, errDB error
	err = nil
	testBooleans := getTestBooleans()

	for i, boolean := range testBooleans {
		if i == 0 { //ID not found case
			errID = fmt.Errorf("unable to find ID")
			errDB = nil
			mockRepo.EXPECT().GetBooleanByIDHelper(gomock.Any(), gomock.Any()).Return(errID)
		} else { //error while updating DB case
			errID = nil
			errDB = fmt.Errorf("unable to find database")
			mockRepo.EXPECT().GetBooleanByIDHelper(gomock.Any(), gomock.Any()).Return(errID)
			mockRepo.EXPECT().UpdateBooleanHelper(gomock.Any(), gomock.Any()).Return(errDB)
		}

		jsonRequest, err2 := json.Marshal(models.Boolean{
			Key:   boolean.Key,
			Value: boolean.Value,
		})
		assert.Nil(t, err2)

		controller := NewController(mockRepo)
		tr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(tr)

		c.Request, err = http.NewRequest("PATCH", "/", bytes.NewBuffer(jsonRequest))
		assert.Nil(t, err)

		c.Request.Header.Set("Content-Type", "application/json")
		c.Params = gin.Params{
			{Key: "id", Value: boolean.ID.String()},
		}
		controller.UpdateBoolean(c)

		assert.Equal(t, http.StatusNotFound, tr.Code)
	}
}

func TestUpdateBooleanBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_control.NewMockRepoInterface(ctrl)
	var err error
	err = nil
	mockRepo.EXPECT().GetBooleanByIDHelper(gomock.Any(), gomock.Any()).Return(err)
	mockRepo.EXPECT().UpdateBooleanHelper(gomock.Any(), gomock.Any()).Return(err)

	jsonRequest, err2 := json.Marshal(models.Boolean{
		Key: "key without value",
	})
	assert.Nil(t, err2)

	controller := NewController(mockRepo)
	tr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(tr)
	c.Request, err = http.NewRequest("PATCH", "/", bytes.NewBuffer(jsonRequest))
	assert.Nil(t, err)

	c.Request.Header.Set("Content-Type", "application/json")
	tempID, _ := uuid.NewV4()
	c.Params = gin.Params{
		{Key: "id", Value: tempID.String()},
	}
	controller.UpdateBoolean(c)

	assert.Equal(t, http.StatusBadRequest, tr.Code)
}

func TestDeleteBooleanHelper(t *testing.T) {
	mock, db := initiateTest(t)
	testBooleans := getTestBooleans()

	const sqlDelete = "DELETE FROM `booleans`  WHERE `booleans`.`id` = ?"

	for _, boolean := range testBooleans {
		mock.ExpectBegin()
		mock.ExpectExec(sqlDelete).
			WithArgs(boolean.ID).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		r := NewRepo(db)
		err := r.DeleteBooleanHelper(&boolean)
		assert.Nil(t, err)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	}
}

func TestDeleteBooleanStatusOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_control.NewMockRepoInterface(ctrl)
	var err error
	err = nil
	mockRepo.EXPECT().GetBooleanByIDHelper(gomock.Any(), gomock.Any()).Return(err)
	mockRepo.EXPECT().DeleteBooleanHelper(gomock.Any()).Return(err)

	controller := NewController(mockRepo)
	tr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(tr)
	controller.DeleteBoolean(c)

	// fmt.Println(c.Request.Response)
	assert.Equal(t, http.StatusOK, tr.Code)
}

func TestDeleteBooleanStatusNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_control.NewMockRepoInterface(ctrl)
	var errID, errDB error
	for i := 0; i < 2; i++ {
		if i == 0 {
			errID = fmt.Errorf("unable to find ID")
			errDB = nil
			mockRepo.EXPECT().GetBooleanByIDHelper(gomock.Any(), gomock.Any()).Return(errID)
		} else {
			errID = nil
			errDB = fmt.Errorf("unable to find database")
			mockRepo.EXPECT().GetBooleanByIDHelper(gomock.Any(), gomock.Any()).Return(errID)
			mockRepo.EXPECT().DeleteBooleanHelper(gomock.Any()).Return(errDB)
		}

		controller := NewController(mockRepo)
		tr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(tr)
		controller.DeleteBoolean(c)

		assert.Equal(t, http.StatusNotFound, tr.Code)
	}
}

// func TestRoutes(t *testing.T) {
// 	ts := httptest.NewServer(SetupRoutes())
// 	defer ts.Close()

// 	_, err := http.Get(fmt.Sprintf("%s/", ts.URL))
// 	if err != nil {
// 		t.Fatalf("Expected no error, got %v", err)
// 	}

// 	router := SetupRoutes()
// 	req, _ := http.NewRequest("GET", "/", nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	// if resp.StatusCode != 200 {
// 	// 	t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
// 	// }

// }
