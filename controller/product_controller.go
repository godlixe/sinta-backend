package controller

// import (
// 	"jwt-hacktiv8/common"
// 	"jwt-hacktiv8/dto"
// 	"jwt-hacktiv8/service"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type ProductController interface {
// 	InsertProduct(ctx *gin.Context)
// 	UpdateProduct(ctx *gin.Context)
// }

// type productController struct {
// 	productService service.ProductService
// 	jwtService     service.JWTService
// }

// func NewProductController(ps service.ProductService, js service.JWTService) ProductController {
// 	return &productController{
// 		productService: ps,
// 		jwtService:     js,
// 	}
// }

// func (c *productController) InsertProduct(ctx *gin.Context) {
// 	var productDTO dto.ProductCreateDTO

// 	if err := ctx.ShouldBind(&productDTO); err != nil {
// 		res := common.BuildErrorResponse("Failed to bind product request", err.Error(), common.EmptyObj{})
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	token := ctx.MustGet("token").(string)
// 	userId, _ := c.jwtService.GetUserIDByToken(token)

// 	productDTO.UserID = userId

// 	result, err := c.productService.InsertProduct(ctx.Request.Context(), productDTO)
// 	if err != nil {
// 		res := common.BuildErrorResponse("Failed to create product", err.Error(), common.EmptyObj{})
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, result)
// }

// func (c *productController) UpdateProduct(ctx *gin.Context) {
// 	var productDTO dto.ProductUpdateDTO

// 	if err := ctx.ShouldBind(&productDTO); err != nil {
// 		res := common.BuildErrorResponse("Failed to bind product request", err.Error(), common.EmptyObj{})
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}
// 	productDTO.ID = ctx.MustGet("product_id").(uint)
// 	result, err := c.productService.UpdateProduct(ctx.Request.Context(), productDTO)
// 	if err != nil {
// 		res := common.BuildErrorResponse("Failed to create product", err.Error(), common.EmptyObj{})
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, result)
// }
