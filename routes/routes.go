package routes
import (
	"github.com/gin-gonic/gin"
	"mongodb-com-go/controller"
)


func LidaComRequisicoes() {
	ginRouter := gin.Default()
	ginRouter.GET("/", controller.BemVindo)
	ginRouter.POST("/usuarios", controller.CriaUsuario)
	ginRouter.GET("/usuarios", controller.ListaUsuarios)
	ginRouter.GET("/usuarios/:id", controller.PegaUmUsuario)
	ginRouter.PATCH("/usuarios/mudanome/:id", controller.AtualizaUsuario)
	ginRouter.DELETE("/usuarios/:id", controller.ApagaUsuario)
	ginRouter.Run(":8000")
}