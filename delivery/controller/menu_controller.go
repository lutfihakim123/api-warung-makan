package controller

import (
	"api-warung-makan/model"
	"api-warung-makan/usecase"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	router      *gin.Engine
	menuUseCase usecase.MenuUseCase
}

func (mc *MenuController) CreateNewMenu(ctx *gin.Context) {
	var newMenu *model.Menu
	err := ctx.ShouldBind(&newMenu)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	img, _ := ctx.FormFile("img")
	if img != nil {
		ctx.SaveUploadedFile(img, "assets/img/menu/"+img.Filename)
		newMenu.Img = img.Filename
	} else {
		newMenu.Img = ""
	}
	mc.menuUseCase.CreateNewMenu(newMenu)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    newMenu,
	})
}

func (mc *MenuController) GetAllMenu(ctx *gin.Context) {
	var err error
	page, _ := strconv.Atoi(ctx.Query("page"))
	totalRows, _ := strconv.Atoi(ctx.Query("totalRows"))
	if page == 0 || totalRows == 0 {
		page = 1
		totalRows = 10
	}
	menus, err := mc.menuUseCase.GetAllMenu(page, totalRows)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    menus,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message":   "OK",
			"data":      menus,
			"page":      page,
			"totalRows": totalRows,
		})
	}
}

func (mc *MenuController) UpdateMenu(ctx *gin.Context) {
	var newMenu *model.Menu
	err := ctx.ShouldBind(&newMenu)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		img, _ := ctx.FormFile("img")
		if img != nil {
			ctx.SaveUploadedFile(img, "assets/img/menu/"+img.Filename)
			newMenu.Img = img.Filename
		} else {
			newMenu.Img = ""
		}
		err := mc.menuUseCase.UpdateMenu(*newMenu)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data":    *newMenu,
				"message": "OK",
			})
		}
	}
}

func (m *MenuController) GetMenuById(ctx *gin.Context) {
	idMenu := ctx.Param("id")
	responseUc, _ := m.menuUseCase.GetMenuById(idMenu)
	if (responseUc == model.Menu{}) {
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

func (m *MenuController) DeleteMenu(ctx *gin.Context) {
	id := ctx.Param("id")
	responseUc, _ := m.menuUseCase.GetMenuById(id)
	err := m.menuUseCase.DeleteMenu(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id tidak ditemukan",
		})
	} else {
		path := "assets/img/menu/" + responseUc.Img
		err := os.Remove(path)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"error":   err,
		})
	}
}

func NewMenuController(router *gin.Engine, menuUseCase usecase.MenuUseCase) *MenuController {
	newMenuController := MenuController{
		router:      router,
		menuUseCase: menuUseCase,
	}
	menu := router.Group("warung/menu")
	menu.POST("", newMenuController.CreateNewMenu)
	menu.GET("", newMenuController.GetAllMenu)
	menu.GET("/:id", newMenuController.GetMenuById)
	menu.PUT("", newMenuController.UpdateMenu)
	menu.DELETE("/:id", newMenuController.DeleteMenu)
	return &newMenuController
}
