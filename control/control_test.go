package control

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akuldr67/Boolean/models"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/akuldr67/Boolean/config"

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

// func TestRoutes(t *testing.T) {
// 	// ts := httptest.NewServer(SetupRoutes())
// 	// defer ts.Close()

// 	// _, err := http.Get(fmt.Sprintf("%s/", ts.URL))
// 	// if err != nil {
// 	// 	t.Fatalf("Expected no error, got %v", err)
// 	// }

// 	router := SetupRoutes()
// 	req, _ := http.NewRequest("GET", "/", nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	// if resp.StatusCode != 200 {
// 	// 	t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
// 	// }

// }

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

func initiateTest(t *testing.T) sqlmock.Sqlmock {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	assert.Nil(t, err, "an error was not expected when opening a stub database connection")

	// config.DB, err = gorm.Open(mysql.Open("db"), &gorm.Config{})
	config.DB, err = gorm.Open("mysql", db)

	assert.Nil(t, err, "error while opening test mysql database")

	return mock
}

func TestGetAllBooleansHelper(t *testing.T) {
	mock := initiateTest(t)

	testBooleans := getTestBooleans()

	rows := sqlmock.NewRows([]string{"id", "key", "value"})

	expected := `[`
	for _, boolean := range testBooleans {
		rows = rows.AddRow(boolean.ID, boolean.Key, boolean.Value)
		expected = expected + `{"id":"` + boolean.ID.String() + `","key":"` + boolean.Key + `","value":`
		if *boolean.Value == true {
			expected += `true`
		} else {
			expected += `false`
		}
		expected += `},`
	}
	expected = expected[:len(expected)-1]
	expected += `]`

	mock.ExpectQuery("SELECT * FROM `booleans`").WillReturnRows(rows)

	var allBooleans []models.Boolean
	err := getAllBooleansHelper(&allBooleans)
	assert.Nil(t, err)

	assert.Equal(t, allBooleans, testBooleans)
}

func TestGetAllBooleans(t *testing.T) {
	mock := initiateTest(t)

	testBooleans := getTestBooleans()

	rows := sqlmock.NewRows([]string{"id", "key", "value"})

	expected := `[`
	for _, boolean := range testBooleans {
		rows = rows.AddRow(boolean.ID, boolean.Key, boolean.Value)
		expected = expected + `{"id":"` + boolean.ID.String() + `","key":"` + boolean.Key + `","value":`
		if *boolean.Value == true {
			expected += `true`
		} else {
			expected += `false`
		}
		expected += `},`
	}
	expected = expected[:len(expected)-1]
	expected += `]`

	mock.ExpectQuery("SELECT * FROM `booleans`").WillReturnRows(rows)

	tr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(tr)
	getAllBooleans(c)

	// assert.Equal(t, *tr.Body, *rows)

	assert.Equal(t, tr.Body.String(), expected)

	assert.Equal(t, tr.Code, http.StatusOK)
}

// func TestGetAllBooleans(t *testing.T) {
// 	testBooleans := getTestBooleans()

// }

func TestCreateBooleansHelper(t *testing.T) {
	mock := initiateTest(t)

	testBooleans := getTestBooleans()

	const sqlInsert = "INSERT INTO `booleans` (`id`,`key`,`value`) VALUES (?,?,?)"

	for _, boolean := range testBooleans {
		mock.ExpectBegin()
		mock.ExpectExec(sqlInsert).
			WithArgs(boolean.ID.String(), boolean.Key, boolean.Value).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		err := createBooleanHelper(&boolean)
		assert.Nil(t, err)
	}
}

func TestCreateBooleans(t *testing.T) {
	_ = initiateTest(t)
	var err error

	testBooleans := getTestBooleans()

	for _, boolean := range testBooleans {
		jsonRequest, err2 := json.Marshal(models.Boolean{
			Key:   boolean.Key,
			Value: boolean.Value,
		})
		assert.Nil(t, err2)

		tr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(tr)

		c.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(jsonRequest))
		if err != nil {
			fmt.Println("******** here *********")
			t.Fatal(err)
		}
		c.Request.Header.Set("Content-Type", "application/json")
		createBoolean(c)
		// fmt.Println(tr.Code)
	}
}

// func TestCreateBooleans(t *testing.T) {
// mock := initiateTest(t)

// 	testBooleans := getTestBooleans()

// 	const sqlInsert = "INSERT INTO `booleans` (`id`,`key`,`value`) VALUES (?,?,?)"

// 	for _, boolean := range testBooleans {
// 		mock.ExpectBegin()
// 		mock.ExpectExec(sqlInsert).
// 			WithArgs(boolean.ID.String(), boolean.Key, boolean.Value).
// 			WillReturnResult(sqlmock.NewResult(0, 0))
// 		mock.ExpectCommit()

// 		addString := `{"key":"` + boolean.Key + `","value":`
// 		if *boolean.Value == true {
// 			addString += `true`
// 		} else {
// 			addString += `false`
// 		}
// 		addString += `}`
// 		// var jsonStr = []byte(addString)

// 		jsonRequest, err2 := json.Marshal(models.Boolean{
// 			Key:   boolean.Key,
// 			Value: boolean.Value,
// 		})
// 		assert.Nil(t, err2)

// 		tr := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(tr)

// 		c.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(jsonRequest))
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		c.Request.Header.Set("Content-Type", "application/json")
// 		createBoolean(c)
// 	}
// }

func TestGetBooleanByIDHelper(t *testing.T) {
	mock := initiateTest(t)

	testBooleans := getTestBooleans()

	for _, boolean := range testBooleans {
		rows := sqlmock.
			NewRows([]string{"id", "key", "value"}).
			AddRow(boolean.ID, boolean.Key, boolean.Value)

		mock.ExpectQuery("SELECT * FROM `booleans` WHERE (id = ?) ORDER BY `booleans`.`id` ASC LIMIT 1").
			WithArgs(boolean.ID).
			WillReturnRows(rows)

		var newBoolean models.Boolean
		err := getBooleanByIDHelper(&newBoolean, boolean.ID.String())
		assert.Nil(t, err)
		assert.Equal(t, newBoolean.ID, boolean.ID)
		assert.Equal(t, newBoolean.Key, boolean.Key)
		assert.Equal(t, newBoolean.Value, boolean.Value)
	}
}

func TestGetBooleanByID(t *testing.T) {
	mock := initiateTest(t)

	testBooleans := getTestBooleans()

	for _, boolean := range testBooleans {
		rows := sqlmock.
			NewRows([]string{"id", "key", "value"}).
			AddRow(boolean.ID, boolean.Key, boolean.Value)

		mock.ExpectQuery("SELECT * FROM `booleans` WHERE (id = ?) ORDER BY `booleans`.`id` ASC LIMIT 1").
			WithArgs(boolean.ID).
			WillReturnRows(rows)

		tr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(tr)
		c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
		c.Params = gin.Params{
			{Key: "id", Value: boolean.ID.String()},
		}

		getBooleanByID(c)

		assert.Equal(t, tr.Code, http.StatusOK)

		var b models.Boolean
		err := json.Unmarshal(tr.Body.Bytes(), &b)
		assert.Nil(t, err)
		assert.Equal(t, boolean, b)
	}
}

func TestUpdateBooleanHelper(t *testing.T) {
	mock := initiateTest(t)

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

		err := updateBooleanHelper(&boolean, boolean)
		assert.Nil(t, err)
	}
}

func TestUpdateBoolean(t *testing.T) {
	mock := initiateTest(t)

	testBooleans := getTestBooleans()

	// const query = "SELECT * FROM `booleans`  WHERE (id = ?) ORDER BY `booleans`.`id` ASC LIMIT 1"

	for _, boolean := range testBooleans {
		mock.ExpectBegin()

		// mock.ExpectExec(query).
		// 	WithArgs(boolean.ID).
		// 	WillReturnResult(sqlmock.NewResult(0, 0))

		var sqlUpdate string
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

		jsonRequest, err := json.Marshal(models.Boolean{
			Value: boolean.Value,
			Key:   boolean.Key,
		})
		assert.Nil(t, err)

		tr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(tr)
		c.Request, _ = http.NewRequest(http.MethodGet, "", bytes.NewBuffer(jsonRequest))
		c.Params = gin.Params{
			{Key: "id", Value: boolean.ID.String()},
		}

		updateBoolean(c)
		// fmt.Println(tr.Code)
		// assert.Equal(t, http.StatusOK, tr.Code)

	}
}

func TestDeleteBooleanHelper(t *testing.T) {
	mock := initiateTest(t)
	testBooleans := getTestBooleans()

	const sqlDelete = "DELETE FROM `booleans`  WHERE `booleans`.`id` = ?"

	for _, boolean := range testBooleans {
		mock.ExpectBegin()
		mock.ExpectExec(sqlDelete).
			WithArgs(boolean.ID).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		err := deleteBooleanHelper(&boolean)
		assert.Nil(t, err)
	}

}

func TestDeleteBoolean(t *testing.T) {
	mock := initiateTest(t)

	testBooleans := getTestBooleans()

	const sqlDelete = "DELETE FROM `booleans`  WHERE `booleans`.`id` = ?"

	for _, boolean := range testBooleans {
		mock.ExpectBegin()
		mock.ExpectExec(sqlDelete).
			WithArgs(boolean.ID).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		tr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(tr)
		c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
		c.Params = gin.Params{
			{Key: "id", Value: boolean.ID.String()},
		}

		deleteBoolean(c)
		// fmt.Println(tr.Code)
		// assert.Equal(t, tr.Code, http.StatusNoContent)
	}
}
