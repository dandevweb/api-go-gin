package routes

import (
	studentsController "api-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", studentsController.All)
	r.GET("/alunos/cpf/:cpf", studentsController.GetByCpf)
	r.GET("/alunos/:id", studentsController.GetById)
	r.PUT("/alunos/:id", studentsController.Update)
	r.DELETE("/alunos/:id", studentsController.Delete)
	r.POST("alunos", studentsController.Create)

	r.Run()
}
