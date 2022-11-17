package controller

import (
	"net/http"
	"sinta-backend/common"
	"sinta-backend/dto"
	"sinta-backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authController struct {
	tokoService service.TokoService
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(ts service.TokoService, as service.AuthService, js service.JWTService) AuthController {
	return &authController{
		tokoService: ts,
		authService: as,
		jwtService:  js,
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var tokoDTO dto.TokoCreateDTO
	errDTO := ctx.ShouldBind(&tokoDTO)
	if errDTO != nil {
		response := common.BuildErrorResponse("Failed to process request", errDTO.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	isDuplicateUsername, _ := c.authService.CheckUsernameDuplicate(ctx.Request.Context(), tokoDTO.Username)
	if isDuplicateUsername {
		response := common.BuildErrorResponse("Failed to process request", "Duplicate Email", common.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	createdToko, err := c.tokoService.CreateToko(ctx.Request.Context(), tokoDTO)
	if err != nil {
		response := common.BuildErrorResponse("Failed to process request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}
	tokoID := strconv.FormatUint(uint64(createdToko.ID), 10)
	token := c.jwtService.GenerateToken(tokoID)

	response := common.BuildResponse(true, "OK", token)
	ctx.JSON(http.StatusCreated, response)
}

func (c *authController) Login(ctx *gin.Context) {
	var tokoLoginDTO dto.TokoLoginDTO
	if errDTO := ctx.ShouldBind(&tokoLoginDTO); errDTO != nil {
		response := common.BuildErrorResponse("Failed to process request", errDTO.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult, _ := c.authService.VerifyCredential(ctx.Request.Context(), tokoLoginDTO.Username, tokoLoginDTO.Password)
	if !authResult {
		response := common.BuildErrorResponse("Error Logging in", "Invalid Credentials", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	toko, err := c.tokoService.GetTokoByUsername(ctx.Request.Context(), tokoLoginDTO.Username)
	if err != nil {
		response := common.BuildErrorResponse("Failed to process request", err.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	tokoID := strconv.FormatUint(uint64(toko.ID), 10)
	generatedToken := c.jwtService.GenerateToken(tokoID)
	response := common.BuildResponse(true, "OK", generatedToken)
	ctx.JSON(http.StatusOK, response)
}
