package routes

import (
	"sinta-backend/controller"
	"sinta-backend/middleware"
	"sinta-backend/service"

	"github.com/gin-gonic/gin"
)

func StokRoutes(router *gin.Engine, stokController controller.StokController, jwtService service.JWTService) {
	stokRoutes := router.Group("/stok", middleware.Authenticate(jwtService))
	{
		stokRoutes.GET("", stokController.GetStokByTokoID)
		stokRoutes.POST("", stokController.InsertStok)
		stokRoutes.PUT("", stokController.UpdateStok)

	}
}
