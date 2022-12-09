package controller

import (
	"net/http"
	"sinta-backend/common"
	"sinta-backend/dto"
	"sinta-backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AjuanController interface {
	GetAllAjuan(ctx *gin.Context)
	GetAjuanByID(ctx *gin.Context)
	CreateAjuan(ctx *gin.Context)
	AcceptAjuan(ctx *gin.Context)
	DeclineAjuan(ctx *gin.Context)
}

type ajuanController struct {
	ajuanService service.AjuanService
	jwtService   service.JWTService
}

func NewAjuanController(as service.AjuanService, js service.JWTService) AjuanController {
	return &ajuanController{
		ajuanService: as,
		jwtService:   js,
	}
}

func (c *ajuanController) GetAllAjuan(ctx *gin.Context) {
	result, err := c.ajuanService.GetAllAjuan(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to get ajuan", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *ajuanController) GetAjuanByID(ctx *gin.Context) {
	ajuanID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	result, err := c.ajuanService.GetAjuanByID(ctx.Request.Context(), ajuanID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get ajuan", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *ajuanController) CreateAjuan(ctx *gin.Context) {

	token := ctx.MustGet("token").(string)
	tokoID, err := c.jwtService.GetTokoIDByToken(token)
	if err != nil {
		res := common.BuildErrorResponse("token invalid", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var ajuan dto.AjuanCreateDTO
	if err := ctx.ShouldBind(&ajuan); err != nil {
		res := common.BuildErrorResponse("Failed to bind ajuan", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ajuan.TokoID = tokoID

	result, err := c.ajuanService.CreateAjuan(ctx.Request.Context(), tokoID, ajuan)
	if err != nil {
		res := common.BuildErrorResponse("Failed to create ajuan", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *ajuanController) AcceptAjuan(ctx *gin.Context) {
	ajuanID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	err := c.ajuanService.AcceptAjuan(ctx.Request.Context(), ajuanID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to accept ajuan", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (c *ajuanController) DeclineAjuan(ctx *gin.Context) {
	ajuanID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	err := c.ajuanService.DeclineAjuan(ctx.Request.Context(), ajuanID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to accept ajuan", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
