package middleware

import (
	"fmt"
	"net/http"
	"sinta-backend/common"
	"sinta-backend/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService service.JWTService, role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := common.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := common.BuildErrorResponse("Invalid Token", "Token is invalid", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := common.BuildErrorResponse("Invalid Token", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !token.Valid {
			response := common.BuildErrorResponse("Access denied", "You dont have access", nil)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		tokoRole, err := jwtService.GetRoleByToken(string(authHeader))
		fmt.Println("ROLE", tokoRole)
		if err != nil || (tokoRole != "admin" && tokoRole != role) {
			response := common.BuildErrorResponse("Access denied", "You dont have access", nil)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		// get userID from token
		tokoID, err := jwtService.GetTokoIDByToken(authHeader)
		if err != nil {
			response := common.BuildErrorResponse("Invalid Token", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		fmt.Println("ROLE", tokoRole)
		c.Set("token", authHeader)
		c.Set("tokoID", tokoID)
		c.Next()
	}
}
