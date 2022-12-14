package routes

import (
	"sinta-backend/controller"
	"sinta-backend/middleware"
	"sinta-backend/service"

	"github.com/gin-gonic/gin"
)

func StokRoutes(router *gin.Engine, stokController controller.StokController, jwtService service.JWTService) {
	stokRoutes := router.Group("/stok", middleware.Authenticate(jwtService, "toko"))
	{
		stokRoutes.GET("", stokController.GetStokByTokoID)
		stokRoutes.GET("/produk", stokController.GetProdukStokByTokoID)
		stokRoutes.POST("", stokController.InsertStok)
		stokRoutes.PUT("", stokController.UpdateStok)

	}
}
