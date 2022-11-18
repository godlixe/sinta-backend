package controller

import (
	"net/http"
	"sinta-backend/common"
	"sinta-backend/dto"
	"sinta-backend/service"

	"github.com/gin-gonic/gin"
)

type TransaksiController interface {
	CreateTransaksi(ctx *gin.Context)
	GetAllTransaksiByTokoID(ctx *gin.Context)
}

type transaksiController struct {
	transaksiService service.TransaksiService
	jwtService       service.JWTService
}

func NewTransaksiController(ts service.TransaksiService, js service.JWTService) TransaksiController {
	return &transaksiController{
		transaksiService: ts,
		jwtService:       js,
	}
}

func (c *transaksiController) CreateTransaksi(ctx *gin.Context) {

	token := ctx.MustGet("token").(string)
	tokoID, err := c.jwtService.GetTokoIDByToken(token)
	if err != nil {
		res := common.BuildErrorResponse("token invalid", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var transaksi dto.TransaksiCreateDTO
	if err := ctx.ShouldBind(&transaksi); err != nil {
		res := common.BuildErrorResponse("Failed to bind transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	transaksi.TokoID = tokoID

	result, err := c.transaksiService.CreateTransaksi(ctx.Request.Context(), transaksi, tokoID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to insert transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *transaksiController) GetAllTransaksiByTokoID(ctx *gin.Context) {

	token := ctx.MustGet("token").(string)
	tokoID, err := c.jwtService.GetTokoIDByToken(token)
	if err != nil {
		res := common.BuildErrorResponse("token invalid", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.transaksiService.GetAllTransaksiByTokoID(ctx.Request.Context(), tokoID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}
