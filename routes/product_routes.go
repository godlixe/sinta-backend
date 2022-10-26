package routes

// import (
// 	"jwt-hacktiv8/controller"
// 	"jwt-hacktiv8/middleware"
// 	"jwt-hacktiv8/service"

// 	"github.com/gin-gonic/gin"
// )

// func ProductRoutes(router *gin.Engine, productController controller.ProductController, jwtService service.JWTService, productService service.ProductService) {
// 	productRoutes := router.Group("/product", middleware.Authenticate(jwtService))
// 	{
// 		productRoutes.POST("/", productController.InsertProduct)
// 		productRoutes.PUT("/update/:productId", middleware.ProductAuthorization(jwtService, productService), productController.UpdateProduct)
// 	}
// }
