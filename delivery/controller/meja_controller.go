package controller

import (
	"api-warung-makan/model"
	"api-warung-makan/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MejaController struct {
	router      *gin.Engine
	mejaUseCase usecase.MejaUseCase
}

func (mc *MejaController) CreateNewMeja(ctx *gin.Context) {
	var newMeja *model.Meja
	err := ctx.ShouldBind(&newMeja)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	mc.mejaUseCase.CreateNewMeja(newMeja)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    newMeja,
	})
}

func (mc *MejaController) GetAllMeja(ctx *gin.Context) {
	var err error
	page, _ := strconv.Atoi(ctx.Query("page"))
	totalRows, _ := strconv.Atoi(ctx.Query("totalRows"))
	if page == 0 || totalRows == 0 {
		page = 1
		totalRows = 10
	}
	meja, err := mc.mejaUseCase.GetAllMeja(page, totalRows)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    meja,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message":   "OK",
			"data":      meja,
			"page":      page,
			"totalRows": totalRows,
		})
	}
}

func (mc *MejaController) UpdateMeja(ctx *gin.Context) {
	var newMeja *model.Meja
	err := ctx.ShouldBind(&newMeja)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := mc.mejaUseCase.UpdateMeja(*newMeja)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data":    newMeja,
				"message": "OK",
			})
		}
	}
}

func (mc *MejaController) GetMejaById(ctx *gin.Context) {
	id := ctx.Param("id")
	responseUc, _ := mc.mejaUseCase.GetMejaById(id)
	if (responseUc == model.Meja{}) {
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

func (mc *MejaController) DeleteMeja(ctx *gin.Context) {
	id := ctx.Param("id")
	err := mc.mejaUseCase.DeleteMeja(id)
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

func NewMejaController(router *gin.Engine, mejaUseCase usecase.MejaUseCase) *MejaController {
	newMejaController := MejaController{
		router:      router,
		mejaUseCase: mejaUseCase,
	}
	meja := router.Group("warung/meja")
	meja.POST("", newMejaController.CreateNewMeja)
	meja.GET("", newMejaController.GetAllMeja)
	meja.GET("/:id", newMejaController.GetMejaById)
	meja.PUT("", newMejaController.UpdateMeja)
	meja.DELETE("/:id", newMejaController.DeleteMeja)
	return &newMejaController
}
