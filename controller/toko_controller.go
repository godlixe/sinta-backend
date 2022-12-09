package controller

import (
	"net/http"
	"sinta-backend/common"
	"sinta-backend/dto"
	"sinta-backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TokoController interface {
	GetAllToko(ctx *gin.Context)
	UpdateToko(ctx *gin.Context)
	DeleteToko(ctx *gin.Context)
}

type tokoController struct {
	tokoService service.TokoService
	jwtService  service.JWTService
}

func NewTokoController(ts service.TokoService, js service.JWTService) TokoController {
	return &tokoController{
		tokoService: ts,
		jwtService:  js,
	}
}

func (c *tokoController) GetAllToko(ctx *gin.Context) {
	result, err := c.tokoService.GetAllToko(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to get toko", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *tokoController) UpdateToko(ctx *gin.Context) {
	tokoID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	var toko dto.TokoUpdateDTO
	if err := ctx.ShouldBind(&toko); err != nil {
		res := common.BuildErrorResponse("Failed to bind produk", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	toko.ID = tokoID
	result, err := c.tokoService.UpdateToko(ctx.Request.Context(), toko)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update toko", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *tokoController) DeleteToko(ctx *gin.Context) {
	tokoID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	err := c.tokoService.DeleteToko(ctx.Request.Context(), tokoID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to delete toko", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "DELETED", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
