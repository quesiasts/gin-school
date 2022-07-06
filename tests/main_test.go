package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/quesiasts/gin-school/controllers"
	"github.com/quesiasts/gin-school/database"
	"github.com/quesiasts/gin-school/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{Name: "Quésia", Document: "12345678901", Age: 25}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func RemoveStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestStatusCodeIntroduction(t *testing.T) {
	r := SetupRoutes()
	r.GET("/:name", controllers.Introduction)
	req, _ := http.NewRequest("GET", "/quesia", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "It must be equals")

	mockResponse := `{"API says:":"Hi quesia, are you okay?"}`
	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, mockResponse, string(resBody))
}

func TestListAllStudents(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer RemoveStudentMock()
	r := SetupRoutes()
	r.GET("/students", controllers.ListAllStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestSearchStudentsDocument(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer RemoveStudentMock()
	r := SetupRoutes()
	r.GET("/students/document/:document", controllers.SearchForDocument)
	req, _ := http.NewRequest("GET", "/students/document/12345678901", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestSearchStudentID(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer RemoveStudentMock()
	r := SetupRoutes()
	r.GET("/students/:id", controllers.SearchForID)
	pathSearch := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathSearch, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var studentMock models.Student
	json.Unmarshal(res.Body.Bytes(), &studentMock)

	assert.Equal(t, "Quésia", studentMock.Name)
	assert.Equal(t, "12345678901", studentMock.Document)
	assert.Equal(t, 25, studentMock.Age)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestRemoveStudent(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	r := SetupRoutes()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	pathSearch := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathSearch, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestEditStudent(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer RemoveStudentMock()
	r := SetupRoutes()
	r.PATCH("/students/:id", controllers.EditStudent)
	student := models.Student{Name: "Quésia", Document: "47123456789", Age: 25}
	valueJson, _ := json.Marshal(student)
	pathToEdit := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathToEdit, bytes.NewBuffer(valueJson))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var studentMockReleased models.Student
	json.Unmarshal(res.Body.Bytes(), &studentMockReleased)

	assert.Equal(t, "Quésia", studentMockReleased.Name)
	assert.Equal(t, "47123456789", studentMockReleased.Document)
	assert.Equal(t, 25, studentMockReleased.Age)

}
