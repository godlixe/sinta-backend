package controller

import (
	"net/http"
	"sinta-backend/common"
	"sinta-backend/dto"
	"sinta-backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransaksiController interface {
	CreateTransaksi(ctx *gin.Context)
	GetAllTransaksiByTokoID(ctx *gin.Context)
	GetHarianTransaksiByTokoID(ctx *gin.Context)
	GetMingguanTransaksiByTokoID(ctx *gin.Context)
	GetBulananTransaksiByTokoID(ctx *gin.Context)
	GetHarianTotal(ctx *gin.Context)
	GetMingguanTotal(ctx *gin.Context)
	GetBulananTotal(ctx *gin.Context)
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

func (c *transaksiController) GetHarianTransaksiByTokoID(ctx *gin.Context) {
	tokoID, err := strconv.ParseUint(ctx.Param("tokoID"), 10, 64)
	if err != nil {
		res := common.BuildErrorResponse("failed to process request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, err := c.transaksiService.GetHarianTransaksiByTokoID(ctx.Request.Context(), tokoID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *transaksiController) GetMingguanTransaksiByTokoID(ctx *gin.Context) {
	tokoID, err := strconv.ParseUint(ctx.Param("tokoID"), 10, 64)
	if err != nil {
		res := common.BuildErrorResponse("failed to process request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, err := c.transaksiService.GetMingguanTransaksiByTokoID(ctx.Request.Context(), tokoID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *transaksiController) GetBulananTransaksiByTokoID(ctx *gin.Context) {
	tokoID, err := strconv.ParseUint(ctx.Param("tokoID"), 10, 64)
	if err != nil {
		res := common.BuildErrorResponse("failed to process request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, err := c.transaksiService.GetBulananTransaksiByTokoID(ctx.Request.Context(), tokoID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *transaksiController) GetHarianTotal(ctx *gin.Context) {
	result, err := c.transaksiService.GetHarianTotal(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to get transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *transaksiController) GetMingguanTotal(ctx *gin.Context) {
	result, err := c.transaksiService.GetMingguanTotal(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to get transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *transaksiController) GetBulananTotal(ctx *gin.Context) {
	result, err := c.transaksiService.GetBulananTotal(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to get transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}
