package controller

import (
	"api-translate/model"
	"api-translate/service"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type TranslateController struct {
	router *gin.RouterGroup
}

func NewTranslateTextController(router *gin.RouterGroup) *TranslateController {
	o := &TranslateController{router: router}
	// router.POST("/translate-text", o.TranslateTextController)
	router.POST("/translate-text-libre", o.TranslateLibreController)
	return o
}

// // @Tags Translate
// // @Accept json
// // @Param parameter body model.Translate true "PARAM"
// // @Produce json
// // @Router /translate-text [post]
// func (o TranslateController) TranslateTextController(ctx *gin.Context) {
// 	var resp model.TranslateResponse
// 	defer model.SetMetadataResponse(ctx, time.Now(), &resp)

// 	var param model.Translate
// 	if err := ctx.BindJSON(&param); err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	resp = service.TranslateTextAPI(param)
// }

// @Tags Translate
// @Accept json
// @Param parameter body model.Translate true "PARAM"
// @Produce json
// @Router /translate-text-libre [post]
func (o TranslateController) TranslateLibreController(ctx *gin.Context) {
	var resp model.TranslateResponse
	defer model.SetMetadataResponse(ctx, time.Now(), &resp)

	var param model.Translate
	if err := ctx.BindJSON(&param); err != nil {
		log.Println(err)
		return
	}
	resp = service.LibreTranslate(param)
}
