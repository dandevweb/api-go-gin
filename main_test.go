package main

import (
	studentsController "api-go-gin/controllers"
	"api-go-gin/database"
	"api-go-gin/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestStudentByIdHandler(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RoutesSetup()
	r.GET("/alunos/:id", studentsController.GetById)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var studentMock models.Student
	json.Unmarshal(res.Body.Bytes(), &studentMock)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, studentMock.Name, "Nome do aluno de teste")
	assert.Equal(t, studentMock.CPF, "12345678910")
	assert.Equal(t, studentMock.RG, "123456789")
}

func TestDeleteStudentHandler(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	r := RoutesSetup()
	r.DELETE("/alunos/:id", studentsController.Delete)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNoContent, res.Code)
}

func TestUpdateStudentHandler(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RoutesSetup()
	r.PUT("/alunos/:id", studentsController.Update)
	toUpdateStudent := models.Student{Name: "Nome do aluno de teste updated", CPF: "12345678911", RG: "123456780"}
	studentJson, _ := json.Marshal(toUpdateStudent)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PUT", path, bytes.NewBuffer(studentJson))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var studentUpdated models.Student
	json.Unmarshal(res.Body.Bytes(), &studentUpdated)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, studentUpdated.Name, toUpdateStudent.Name)
	assert.Equal(t, studentUpdated.CPF, toUpdateStudent.CPF)
	assert.Equal(t, studentUpdated.RG, toUpdateStudent.RG)

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
