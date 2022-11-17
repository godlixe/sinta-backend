package controller

import (
	"net/http"
	"sinta-backend/common"
	"sinta-backend/dto"
	"sinta-backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProdukController interface {
	CreateProduk(ctx *gin.Context)
	GetAllProduk(ctx *gin.Context)
	UpdateProduk(ctx *gin.Context)
	DeleteProduk(ctx *gin.Context)
}

type produkController struct {
	produkService service.ProdukService
}

func NewProdukController(ps service.ProdukService) ProdukController {
	return &produkController{
		produkService: ps,
	}
}

func (c *produkController) CreateProduk(ctx *gin.Context) {
	var produk dto.ProdukCreateDTO
	if err := ctx.ShouldBind(&produk); err != nil {
		res := common.BuildErrorResponse("Failed to bind produk", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.produkService.CreateProduk(ctx.Request.Context(), produk)
	if err != nil {
		res := common.BuildErrorResponse("Failed to insert produk", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *produkController) GetAllProduk(ctx *gin.Context) {
	result, err := c.produkService.GetAllProduk(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to insert produk", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *produkController) UpdateProduk(ctx *gin.Context) {
	var produk dto.ProdukUpdateDTO
	if err := ctx.ShouldBind(&produk); err != nil {
		res := common.BuildErrorResponse("Failed to bind produk", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.produkService.UpdateProduk(ctx.Request.Context(), produk)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update produk", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *produkController) DeleteProduk(ctx *gin.Context) {
	produkID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	err := c.produkService.DeleteProduk(ctx.Request.Context(), produkID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to insert produk", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "DELETED", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
