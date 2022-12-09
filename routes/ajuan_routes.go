package routes

import (
	"sinta-backend/controller"
	"sinta-backend/middleware"
	"sinta-backend/service"

	"github.com/gin-gonic/gin"
)

func AjuanRoutes(router *gin.Engine, ajuanController controller.AjuanController, jwtService service.JWTService) {
	ajuanRoutes := router.Group("/ajuan", middleware.Authenticate(jwtService))
	{
		ajuanRoutes.GET("", ajuanController.GetAllAjuan)
		ajuanRoutes.GET("/:id", ajuanController.GetAjuanByID)
		ajuanRoutes.POST("", ajuanController.CreateAjuan)
		ajuanRoutes.POST("/accept/:id", ajuanController.AcceptAjuan)
		ajuanRoutes.POST("/decline/:id", ajuanController.DeclineAjuan)
	}
}
