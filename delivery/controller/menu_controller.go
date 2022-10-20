package controller

import (
	"api-warung-makan/model"
	"api-warung-makan/usecase"
	"api-warung-makan/utils"
	"net/http"
	"os"
	"strconv"
	"strings"

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
		path := utils.Base_url + "menu/image/" + img.Filename
		ctx.SaveUploadedFile(img, "assets/menu/"+img.Filename)
		newMenu.Img = path
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
		responseUc, _ := mc.menuUseCase.GetMenuById(newMenu.Id)
		nameFile := strings.Split(responseUc.Img, "/")
		path := "assets/menu/" + nameFile[5]
		error := os.Remove(path)
		if error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "data not found",
			})
		}
		img, _ := ctx.FormFile("img")
		var tempImg string
		if img != nil {
			path := utils.Base_url + "menu/image/" + img.Filename
			ctx.SaveUploadedFile(img, "assets/menu/"+img.Filename)
			tempImg = path
		} else {
			tempImg = ""
		}
		newMenu.Img = tempImg
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

func (mc *MenuController) GetMenuById(ctx *gin.Context) {
	idMenu := ctx.Param("id")
	responseUc, _ := mc.menuUseCase.GetMenuById(idMenu)
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

func (mc *MenuController) DeleteMenu(ctx *gin.Context) {
	id := ctx.Param("id")
	responseUc, _ := mc.menuUseCase.GetMenuById(id)
	err := mc.menuUseCase.DeleteMenu(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id tidak ditemukan",
		})
	} else {
		nameFile := strings.Split(responseUc.Img, "/")
		path := "assets/menu/" + nameFile[5]
		err := os.Remove(path)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"error":   err,
		})
	}
}

func (mc *MenuController) GetImageMenu(ctx *gin.Context) {
	img := ctx.Param("img")
	image := "assets/menu/" + img
	ctx.File(image)
}

func NewMenuController(router *gin.Engine, menuUseCase usecase.MenuUseCase) *MenuController {
	newMenuController := MenuController{
		router:      router,
		menuUseCase: menuUseCase,
	}
	menu := router.Group("menu")
	menu.POST("", newMenuController.CreateNewMenu)
	menu.GET("", newMenuController.GetAllMenu)
	menu.GET("/:id", newMenuController.GetMenuById)
	menu.GET("/image/:img", newMenuController.GetImageMenu)
	menu.PUT("", newMenuController.UpdateMenu)
	menu.DELETE("/:id", newMenuController.DeleteMenu)
	return &newMenuController
}
