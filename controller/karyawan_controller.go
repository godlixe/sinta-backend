package controller

import (
	"net/http"
	"sinta-backend/common"
	"sinta-backend/dto"
	"sinta-backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KaryawanController interface {
	CreateKaryawan(ctx *gin.Context)
	GetAllKaryawan(ctx *gin.Context)
	UpdateKaryawan(ctx *gin.Context)
	DeleteKaryawan(ctx *gin.Context)
}

type karyawanController struct {
	karyawanService service.KaryawanService
}

func NewKaryawanController(ks service.KaryawanService) KaryawanController {
	return &karyawanController{
		karyawanService: ks,
	}
}

func (c *karyawanController) CreateKaryawan(ctx *gin.Context) {
	var karyawan dto.KaryawanCreateDTO
	if err := ctx.ShouldBind(&karyawan); err != nil {
		res := common.BuildErrorResponse("Failed to bind karyawan", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.karyawanService.CreateKaryawan(ctx.Request.Context(), karyawan)
	if err != nil {
		res := common.BuildErrorResponse("Failed to insert karyawan", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *karyawanController) GetAllKaryawan(ctx *gin.Context) {
	result, err := c.karyawanService.GetAllKaryawan(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to get karyawan", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *karyawanController) UpdateKaryawan(ctx *gin.Context) {
	karyawanID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	var karyawan dto.KaryawanUpdateDTO
	if err := ctx.ShouldBind(&karyawan); err != nil {
		res := common.BuildErrorResponse("Failed to bind karyawan", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	karyawan.ID = karyawanID
	result, err := c.karyawanService.UpdateKaryawan(ctx.Request.Context(), karyawan)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update karyawan", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *karyawanController) DeleteKaryawan(ctx *gin.Context) {
	karyawanID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	err := c.karyawanService.DeleteKaryawan(ctx.Request.Context(), karyawanID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to insert karyawan", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "DELETED", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
