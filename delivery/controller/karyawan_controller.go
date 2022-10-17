package controller

import (
	"api-warung-makan/model"
	"api-warung-makan/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KaryawanController struct {
	router          *gin.Engine
	karyawanUseCase usecase.KaryawanUseCase
}

func (kc *KaryawanController) CreateNewKaryawan(ctx *gin.Context) {
	var newKaryawan *model.Karyawan
	err := ctx.ShouldBind(&newKaryawan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	kc.karyawanUseCase.CreateNewKaryawan(newKaryawan)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    newKaryawan,
	})
}

func (kc *KaryawanController) GetAllKaryawan(ctx *gin.Context) {
	var err error
	page, _ := strconv.Atoi(ctx.Query("page"))
	totalRows, _ := strconv.Atoi(ctx.Query("totalRows"))
	if page == 0 || totalRows == 0 {
		page = 1
		totalRows = 10
	}
	karyawan, err := kc.karyawanUseCase.GetAllKaryawan(page, totalRows)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    karyawan,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message":   "OK",
			"data":      karyawan,
			"page":      page,
			"totalRows": totalRows,
		})
	}
}

func (kc *KaryawanController) UpdateKaryawan(ctx *gin.Context) {
	var newKaryawan *model.Karyawan
	err := ctx.ShouldBind(&newKaryawan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := kc.karyawanUseCase.UpdateKaryawan(*newKaryawan)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data":    newKaryawan,
				"message": "OK",
			})
		}
	}
}

func (kc *KaryawanController) GetKaryawanById(ctx *gin.Context) {
	id := ctx.Param("id")
	responseUc, _ := kc.karyawanUseCase.GetKaryawanById(id)
	if (responseUc == model.Karyawan{}) {
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

func (kc *KaryawanController) DeleteKaryawan(ctx *gin.Context) {
	id := ctx.Param("id")
	err := kc.karyawanUseCase.DeleteKaryawan(id)
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

func NewKaryawanController(router *gin.Engine, karyawanUseCase usecase.KaryawanUseCase) *KaryawanController {
	newKaryawanController := KaryawanController{
		router:          router,
		karyawanUseCase: karyawanUseCase,
	}
	karyawan := router.Group("warung/karyawan")
	karyawan.POST("", newKaryawanController.CreateNewKaryawan)
	karyawan.GET("", newKaryawanController.GetAllKaryawan)
	karyawan.GET("/:id", newKaryawanController.GetKaryawanById)
	karyawan.PUT("", newKaryawanController.UpdateKaryawan)
	karyawan.DELETE("/:id", newKaryawanController.DeleteKaryawan)
	return &newKaryawanController
}
