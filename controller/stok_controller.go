package controller

import (
	"net/http"
	"sinta-backend/common"
	"sinta-backend/dto"
	"sinta-backend/service"

	"github.com/gin-gonic/gin"
)

type StokController interface {
	GetStokByTokoID(ctx *gin.Context)
	InsertStok(ctx *gin.Context)
	UpdateStok(ctx *gin.Context)
}

type stokController struct {
	stokService service.StokService
	jwtService  service.JWTService
}

func NewStokController(ss service.StokService, js service.JWTService) StokController {
	return &stokController{
		stokService: ss,
		jwtService:  js,
	}
}

func (c *stokController) GetStokByTokoID(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	tokoID, err := c.jwtService.GetTokoIDByToken(token)
	if err != nil {
		res := common.BuildErrorResponse("token invalid", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.stokService.GetStokByTokoID(ctx.Request.Context(), tokoID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get stok", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *stokController) InsertStok(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	tokoID, err := c.jwtService.GetTokoIDByToken(token)
	if err != nil {
		res := common.BuildErrorResponse("token invalid", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var stokDTO dto.StokBatchCreateDTO
	if err := ctx.ShouldBind(&stokDTO); err != nil {
		res := common.BuildErrorResponse("Failed to bind stok", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.stokService.InsertStok(ctx.Request.Context(), stokDTO, tokoID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to create stok", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *stokController) UpdateStok(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	tokoID, err := c.jwtService.GetTokoIDByToken(token)
	if err != nil {
		res := common.BuildErrorResponse("token invalid", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var stokDTO dto.StokBatchUpdateDTO
	if err := ctx.ShouldBind(&stokDTO); err != nil {
		res := common.BuildErrorResponse("Failed to bind stok", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.stokService.UpdateStok(ctx.Request.Context(), stokDTO, tokoID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update stok", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}
