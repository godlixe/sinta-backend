package routes

import (
	"sinta-backend/controller"
	"sinta-backend/middleware"
	"sinta-backend/service"

	"github.com/gin-gonic/gin"
)

func TokoRoutes(router *gin.Engine, tokoController controller.TokoController, jwtService service.JWTService) {
	tokoRoutes := router.Group("/toko", middleware.Authenticate(jwtService))
	{
		tokoRoutes.GET("", tokoController.GetAllToko)
		tokoRoutes.PUT("/:id", tokoController.UpdateToko)
		tokoRoutes.DELETE("/:id", tokoController.DeleteToko)
	}
}
