package routes

import (
	"sinta-backend/controller"
	"sinta-backend/middleware"
	"sinta-backend/service"

	"github.com/gin-gonic/gin"
)

func KaryawanRoutes(router *gin.Engine, karyawanController controller.KaryawanController, jwtService service.JWTService) {
	karyawanRoutes := router.Group("/karyawan", middleware.Authenticate(jwtService))
	{
		karyawanRoutes.GET("", karyawanController.GetAllKaryawan)
		karyawanRoutes.POST("", karyawanController.CreateKaryawan)
	}
}
