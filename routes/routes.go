package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucchesisp/go-gin-api/controllers"
)

func HandleRequest() {
	r := gin.Default()

	r.GET("/alunos", controllers.RetornaTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.RetornaAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PUT("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaPorCPF)

	r.Run(":3000")
}
