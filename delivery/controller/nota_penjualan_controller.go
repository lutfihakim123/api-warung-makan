package controller

import (
	"api-warung-makan/model"
	"api-warung-makan/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type NotaController struct {
	router      *gin.Engine
	notaUseCase usecase.NotaUseCase
}

func (nc *NotaController) CreateNewNota(ctx *gin.Context) {
	var newNota *model.Nota
	err := ctx.ShouldBind(&newNota)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	responseUc, _ := nc.notaUseCase.GetMejaById(newNota.MejaId)
	if responseUc.Status != "availabe" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "meja sedang tidak dapat di gunakan",
			"data":    responseUc.Status,
		})
	} else {
		nc.notaUseCase.CreateNewNota(newNota)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    newNota,
		})
	}
}

func (nc *NotaController) GetAllNota(ctx *gin.Context) {
	var err error
	page, _ := strconv.Atoi(ctx.Query("page"))
	totalRows, _ := strconv.Atoi(ctx.Query("totalRows"))
	if page == 0 || totalRows == 0 {
		page = 1
		totalRows = 10
	}
	nota, err := nc.notaUseCase.GetAllNota(page, totalRows)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nota,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message":   "OK",
			"data":      nota,
			"page":      page,
			"totalRows": totalRows,
		})
	}
}

func (nc *NotaController) UpdateNota(ctx *gin.Context) {
	var newNota *model.Nota
	err := ctx.ShouldBind(&newNota)
	newNota.WaktuPesan = time.Now()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := nc.notaUseCase.UpdateNota(*newNota)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data":    newNota,
				"message": "OK",
			})
		}
	}
}

func (nc *NotaController) GetNotaById(ctx *gin.Context) {
	id := ctx.Param("id")
	responseUc, _ := nc.notaUseCase.GetNotaById(id)
	if (responseUc == model.Nota{}) {
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

func (nc *NotaController) DeleteNota(ctx *gin.Context) {
	id := ctx.Param("id")
	err := nc.notaUseCase.DeleteNota(id)
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

func NewNotaController(router *gin.Engine, notaUseCase usecase.NotaUseCase) *NotaController {
	newNotaController := NotaController{
		router:      router,
		notaUseCase: notaUseCase,
	}
	nota := router.Group("nota")
	nota.POST("", newNotaController.CreateNewNota)
	nota.GET("", newNotaController.GetAllNota)
	nota.GET("/:id", newNotaController.GetNotaById)
	nota.PUT("", newNotaController.UpdateNota)
	nota.DELETE("/:id", newNotaController.DeleteNota)
	return &newNotaController
}
