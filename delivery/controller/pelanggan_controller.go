package controller

import (
	"api-warung-makan/model"
	"api-warung-makan/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PelangganController struct {
	router           *gin.Engine
	pelangganUseCase usecase.PelangganUseCase
}

func (pc *PelangganController) CreateNewPelanggan(ctx *gin.Context) {
	var newPelanggan *model.Pelanggan
	err := ctx.ShouldBind(&newPelanggan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	pc.pelangganUseCase.CreateNewPelanggan(newPelanggan)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    newPelanggan,
	})
}

func (pc *PelangganController) GetAllPelanggan(ctx *gin.Context) {
	var err error
	page, _ := strconv.Atoi(ctx.Query("page"))
	totalRows, _ := strconv.Atoi(ctx.Query("totalRows"))
	if page == 0 || totalRows == 0 {
		page = 1
		totalRows = 10
	}
	pelanggan, err := pc.pelangganUseCase.GetAllPelanggan(page, totalRows)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    pelanggan,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message":   "OK",
			"data":      pelanggan,
			"page":      page,
			"totalRows": totalRows,
		})
	}
}

func (pc *PelangganController) UpdatePelanggan(ctx *gin.Context) {
	var newPelanggan *model.Pelanggan
	err := ctx.ShouldBind(&newPelanggan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := pc.pelangganUseCase.UpdatePelanggan(*newPelanggan)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data":    newPelanggan,
				"message": "OK",
			})
		}
	}
}

func (pc *PelangganController) GetPelangganById(ctx *gin.Context) {
	id := ctx.Param("id")
	responseUc, _ := pc.pelangganUseCase.GetPelangganById(id)
	if (responseUc == model.Pelanggan{}) {
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

func (pc *PelangganController) DeletePelanggan(ctx *gin.Context) {
	id := ctx.Param("id")
	err := pc.pelangganUseCase.DeletePelanggan(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id tidak ditemukan",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}

func NewPelangganController(router *gin.Engine, pelangganUseCase usecase.PelangganUseCase) *PelangganController {
	newPelangganController := PelangganController{
		router:           router,
		pelangganUseCase: pelangganUseCase,
	}
	pelanggan := router.Group("pelanggan")
	pelanggan.POST("", newPelangganController.CreateNewPelanggan)
	pelanggan.GET("", newPelangganController.GetAllPelanggan)
	pelanggan.GET("/:id", newPelangganController.GetPelangganById)
	pelanggan.PUT("", newPelangganController.UpdatePelanggan)
	pelanggan.DELETE("/:id", newPelangganController.DeletePelanggan)
	return &newPelangganController
}
