package controller

import (
	"api-warung-makan/model"
	"api-warung-makan/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RincianController struct {
	router         *gin.Engine
	rincianUseCase usecase.RincianUseCase
}

func (rc *RincianController) CreateNewRincian(ctx *gin.Context) {
	var newRincian *model.Rincian
	err := ctx.ShouldBind(&newRincian)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	responseUc, _ := rc.rincianUseCase.GetMenuById(newRincian.MenuId)
	if newRincian.Kuantitas > responseUc.Stock {
		error_value := fmt.Sprintf("Stock tidak mencukupi saat ini hanya tersedia %d", responseUc.Stock)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": error_value,
		})
	} else {
		newRincian.Total = responseUc.Harga * newRincian.Kuantitas
		newRincian.TempStock = responseUc.Stock - newRincian.Kuantitas
		rc.rincianUseCase.CreateNewRincian(newRincian)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    newRincian,
		})
	}
}

func (rc *RincianController) GetAllRincian(ctx *gin.Context) {
	var err error
	page, _ := strconv.Atoi(ctx.Query("page"))
	totalRows, _ := strconv.Atoi(ctx.Query("totalRows"))
	if page == 0 || totalRows == 0 {
		page = 1
		totalRows = 10
	}
	rincian, err := rc.rincianUseCase.GetAllRincian(page, totalRows)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    rincian,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message":   "OK",
			"data":      rincian,
			"page":      page,
			"totalRows": totalRows,
		})
	}
}

func (rc *RincianController) UpdateRincian(ctx *gin.Context) {
	var newRincian *model.Rincian
	err := ctx.ShouldBind(&newRincian)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		responseUc, _ := rc.rincianUseCase.GetMenuById(newRincian.MenuId)
		if newRincian.Kuantitas > responseUc.Stock {
			error_value := fmt.Sprintf("Stock tidak mencukupi saat ini hanya tersedia %d", responseUc.Stock)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": error_value,
			})
		} else {
			newRincian.Total = responseUc.Harga * newRincian.Kuantitas
			newRincian.TempStock = responseUc.Stock - newRincian.Kuantitas
			rc.rincianUseCase.CreateNewRincian(newRincian)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
				"data":    newRincian,
			})
		}
	}
}

func (rc *RincianController) GetRincianById(ctx *gin.Context) {
	id := ctx.Param("id")
	responseUc, _ := rc.rincianUseCase.GetRincianById(id)
	if (responseUc == model.Rincian{}) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Data tidak ditemukan",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    responseUc,
		})
	}
}

func (rc *RincianController) DeleteRincian(ctx *gin.Context) {
	id := ctx.Param("id")
	err := rc.rincianUseCase.DeleteRincian(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}

func NewRincianController(router *gin.Engine, rincianUseCase usecase.RincianUseCase) *RincianController {
	newRincianController := RincianController{
		router:         router,
		rincianUseCase: rincianUseCase,
	}
	rincian := router.Group("rincian")
	rincian.POST("", newRincianController.CreateNewRincian)
	rincian.GET("", newRincianController.GetAllRincian)
	rincian.GET("/:id", newRincianController.GetRincianById)
	rincian.PUT("", newRincianController.UpdateRincian)
	rincian.DELETE("/:id", newRincianController.DeleteRincian)
	return &newRincianController
}
