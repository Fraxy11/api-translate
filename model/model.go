package model

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Translate struct {
	From string   `json:"from"example:"id"`
	To   string   `json:"to" example:"en"`
	Text []string `json:"text"`
}

type TranslateResponse struct {
	Message    string   `json:"message"`
	Status     bool     `json:"status"`
	Time       string   `json:"time"`
	ResultText []string `json:"resultText"`
}

func SetMetadataResponse(ctx *gin.Context, startTime time.Time, resp *TranslateResponse) {
	code := http.StatusOK
	resp.Status = true
	if resp.Message == "" {
		resp.Message = "OK"
	}

	if resp.Message != "OK" {
		resp.Status = false
		code = 400
	}

	resp.Time = time.Since(startTime).String()
	ctx.JSON(code, resp)
}
