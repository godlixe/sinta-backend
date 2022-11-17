package main

import (
	"log"
	"os"
	"sinta-backend/config"
	"sinta-backend/controller"
	"sinta-backend/repository"
	"sinta-backend/routes"
	"sinta-backend/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	var (
		db             *gorm.DB                  = config.SetupDatabaseConnection()
		tokoRepository repository.TokoRepository = repository.NewTokoRepository(db)
		// productRepository repository.ProductRepository = repository.NewProductRepository(db)

		jwtService  service.JWTService  = service.NewJWTService()
		tokoService service.TokoService = service.NewTokoService(tokoRepository)
		authService service.AuthService = service.NewAuthService(tokoRepository)
		// productService service.ProductService = service.NewProductService(productRepository)

		// productController controller.ProductController = controller.NewProductController(productService, jwtService)
		authController controller.AuthController = controller.NewAuthController(tokoService, authService, jwtService)
	)

	defer config.CloseDatabaseConnection(db)

	server := gin.Default()

	routes.AuthRoutes(server, authController)
	// routes.ProductRoutes(server, productController, jwtService, productService)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
