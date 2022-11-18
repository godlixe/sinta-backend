package routes

import (
	"sinta-backend/controller"
	"sinta-backend/middleware"
	"sinta-backend/service"

	"github.com/gin-gonic/gin"
)

func TransaksiRoutes(router *gin.Engine, transaksiController controller.TransaksiController, jwtService service.JWTService) {
	transaksiRoutes := router.Group("/transaksi", middleware.Authenticate(jwtService))
	{
		transaksiRoutes.GET("", transaksiController.GetAllTransaksiByTokoID)
		transaksiRoutes.POST("", transaksiController.CreateTransaksi)
	}
}
