package routes

import (
	"sinta-backend/controller"
	"sinta-backend/middleware"
	"sinta-backend/service"

	"github.com/gin-gonic/gin"
)

func TransaksiRoutes(router *gin.Engine, transaksiController controller.TransaksiController, jwtService service.JWTService) {
	transaksiRoutes := router.Group("/transaksi")
	{
		transaksiRoutes.GET("", middleware.Authenticate(jwtService, "toko"), transaksiController.GetAllTransaksiByTokoID)
		transaksiRoutes.POST("", middleware.Authenticate(jwtService, "toko"), transaksiController.CreateTransaksi)
		transaksiRoutes.GET("/harian/:tokoID", transaksiController.GetHarianTransaksiByTokoID)
		transaksiRoutes.GET("/mingguan/:tokoID", transaksiController.GetMingguanTransaksiByTokoID)
		transaksiRoutes.GET("/bulanan/:tokoID", transaksiController.GetBulananTransaksiByTokoID)
		transaksiRoutes.GET("/harian/total", transaksiController.GetHarianTotal)
		transaksiRoutes.GET("/mingguan/total", transaksiController.GetMingguanTotal)
		transaksiRoutes.GET("/bulanan/total", transaksiController.GetBulananTotal)
	}
}
