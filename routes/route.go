package routes

import (
	studentsController "api-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", studentsController.All)
	r.GET("/alunos/cpf/:cpf", studentsController.GetByCpf)
	r.GET("/alunos/:id", studentsController.GetById)
	r.PUT("/alunos/:id", studentsController.Update)
	r.DELETE("/alunos/:id", studentsController.Delete)
	r.POST("alunos", studentsController.Create)
	r.GET("/index", studentsController.ShowIndexPage)
	r.NoRoute(studentsController.NotFoundRoute)

	r.Run()
}
