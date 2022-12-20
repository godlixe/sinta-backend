package routes

import (
	"sinta-backend/controller"
	"sinta-backend/middleware"
	"sinta-backend/service"

	"github.com/gin-gonic/gin"
)

func ProdukRoutes(router *gin.Engine, produkController controller.ProdukController, jwtService service.JWTService) {
	produkRoutes := router.Group("/produk")
	{
		produkRoutes.GET("", produkController.GetAllProduk)
		produkRoutes.POST("", middleware.Authenticate(jwtService, "admin"), produkController.CreateProduk)
	}
}
