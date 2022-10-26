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
	userService service.UserService
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(us service.UserService, as service.AuthService, js service.JWTService) AuthController {
	return &authController{
		userService: us,
		authService: as,
		jwtService:  js,
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.UserRegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := common.BuildErrorResponse("Failed to process request", errDTO.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	isDuplicateEmail, _ := c.authService.CheckEmailDuplicate(ctx.Request.Context(), registerDTO.Email)
	if isDuplicateEmail {
		response := common.BuildErrorResponse("Failed to process request", "Duplicate Email", common.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	createdUser, err := c.userService.CreateUser(ctx.Request.Context(), registerDTO)
	if err != nil {
		response := common.BuildErrorResponse("Failed to process request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}
	userId := strconv.FormatUint(uint64(createdUser.ID), 10)
	token := c.jwtService.GenerateToken(userId)
	response := common.BuildResponse(true, "OK", token)
	ctx.JSON(http.StatusCreated, response)
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.UserLoginDTO
	if errDTO := ctx.ShouldBind(&loginDTO); errDTO != nil {
		response := common.BuildErrorResponse("Failed to process request", errDTO.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult, _ := c.authService.VerifyCredential(ctx.Request.Context(), loginDTO.Email, loginDTO.Password)
	if !authResult {
		response := common.BuildErrorResponse("Error Logging in", "Invalid Credentials", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	user, err := c.userService.GetUserByEmail(ctx.Request.Context(), loginDTO.Email)
	if err != nil {
		response := common.BuildErrorResponse("Failed to process request", err.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userId := strconv.FormatUint(uint64(user.ID), 10)
	generatedToken := c.jwtService.GenerateToken(userId)
	response := common.BuildResponse(true, "OK", generatedToken)
	ctx.JSON(http.StatusOK, response)
}
