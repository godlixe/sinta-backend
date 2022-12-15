package routes

import (
	"sinta-backend/controller"
	"sinta-backend/service"

	"github.com/gin-gonic/gin"
)

func KaryawanRoutes(router *gin.Engine, karyawanController controller.KaryawanController, jwtService service.JWTService) {
	karyawanRoutes := router.Group("/karyawan")
	{
		karyawanRoutes.GET("", karyawanController.GetAllKaryawan)
		karyawanRoutes.POST("", karyawanController.CreateKaryawan)
		karyawanRoutes.PUT("/update/:id", karyawanController.UpdateKaryawan)
		karyawanRoutes.DELETE("/delete/:id", karyawanController.DeleteKaryawan)
	}
}
