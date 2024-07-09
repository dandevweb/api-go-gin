package main

import (
	studentsController "api-go-gin/controllers"
	"api-go-gin/database"
	"api-go-gin/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func RoutesSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()

	return routes
}

var ID int

func TestAllStudentsHandler(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RoutesSetup()
	r.GET("/alunos", studentsController.All)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestSearchByCpfHandle(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RoutesSetup()
	r.GET("/alunos/cpf/:cpf", studentsController.GetByCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678910", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

}

func CreateStudentMock() {
	student := models.Student{Name: "Nome do aluno de teste", CPF: "12345678910", RG: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}
