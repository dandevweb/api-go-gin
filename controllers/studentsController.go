package studentsController

import (
	"api-go-gin/database"
	"api-go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func All(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)

	c.JSON(200, students)
}

func Create(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.Validate(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func GetById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(200, student)
}

func Update(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.Validate(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&student)

	c.JSON(200, student)
}

func Delete(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)

	c.JSON(http.StatusNoContent, gin.H{"Not content": "Aluno excluído com sucesso"})
}

func GetByCpf(c *gin.Context) {
	var student models.Student
	cpf := c.Params.ByName("cpf")
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(200, student)
}

func ShowIndexPage(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.HTML(http.StatusOK, "index.html", gin.H{"students": students})
}

func NotFoundRoute(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
