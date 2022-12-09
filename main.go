package main

import (
	"log"
	"os"
	"sinta-backend/config"
	"sinta-backend/controller"
	"sinta-backend/repository"
	"sinta-backend/routes"
	"sinta-backend/service"

	"github.com/gin-contrib/cors"
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
		db                  *gorm.DB                       = config.SetupDatabaseConnection()
		tokoRepository      repository.TokoRepository      = repository.NewTokoRepository(db)
		produkRepository    repository.ProdukRepository    = repository.NewProdukRepository(db)
		transaksiRepository repository.TransaksiRepository = repository.NewTransaksiRepository(db)
		stokRepository      repository.StokRepository      = repository.NewStokRepository(db)
		ajuanRepository     repository.AjuanRepository     = repository.NewAjuanRepository(db)

		jwtService       service.JWTService       = service.NewJWTService()
		tokoService      service.TokoService      = service.NewTokoService(tokoRepository)
		authService      service.AuthService      = service.NewAuthService(tokoRepository)
		produkService    service.ProdukService    = service.NewProdukService(produkRepository)
		transaksiService service.TransaksiService = service.NewTransaksiService(transaksiRepository, stokRepository)
		stokService      service.StokService      = service.NewStokService(stokRepository)
		ajuanService     service.AjuanService     = service.NewAjuanService(ajuanRepository, stokRepository)
		// productService service.ProductService = service.NewProductService(productRepository)

		// productController controller.ProductController = controller.NewProductController(productService, jwtService)
		authController      controller.AuthController      = controller.NewAuthController(tokoService, authService, jwtService)
		tokoController      controller.TokoController      = controller.NewTokoController(tokoService, jwtService)
		produkController    controller.ProdukController    = controller.NewProdukController(produkService)
		transaksiController controller.TransaksiController = controller.NewTransaksiController(transaksiService, jwtService)
		stokController      controller.StokController      = controller.NewStokController(stokService, jwtService)
		ajuanController     controller.AjuanController     = controller.NewAjuanController(ajuanService, jwtService)
	)

	defer config.CloseDatabaseConnection(db)

	server := gin.Default()
	server.Use(cors.Default())
	routes.AuthRoutes(server, authController)
	routes.TokoRoutes(server, tokoController, jwtService)
	routes.ProdukRoutes(server, produkController, jwtService)
	routes.TransaksiRoutes(server, transaksiController, jwtService)
	routes.StokRoutes(server, stokController, jwtService)
	routes.AjuanRoutes(server, ajuanController, jwtService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8087"
	}
	server.Run(":" + port)
}
