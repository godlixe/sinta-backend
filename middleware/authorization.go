package middleware

// import (
// 	"jwt-hacktiv8/common"
// 	"jwt-hacktiv8/service"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// func ProductAuthorization(jwtService service.JWTService, productService service.ProductService) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		productId, err := strconv.Atoi(ctx.Param("productId"))
// 		if err != nil {
// 			response := common.BuildErrorResponse("Failed to process request", "No parameter ID", nil)
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}
// 		product, err := productService.GetProductByID(ctx.Request.Context(), uint(productId))
// 		if err != nil {
// 			response := common.BuildErrorResponse("Failed to process request", "No product with that ID", nil)
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}
// 		token := ctx.MustGet("token").(string)
// 		userId, _ := jwtService.GetUserIDByToken(token)

// 		if product.UserID != userId {
// 			response := common.BuildErrorResponse("Failed to process request", "You are not authorized to access this data", nil)
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}
// 		ctx.Set("product_id", uint(productId))
// 		ctx.Next()
// 	}
// }
