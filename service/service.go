package service

import (
	"api-translate/model"
	"api-translate/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tidwall/gjson"

	translate "cloud.google.com/go/translate/apiv3"
	translatepb "cloud.google.com/go/translate/apiv3/translatepb"

	"github.com/logrusorgru/aurora"
)

func TranslateTextAPI(param model.Translate) (response model.TranslateResponse) {
	url := os.Getenv("URL_API_TRANSLATE")

	// bodyParams := map[string]interface{}{
	// 	"contents":           param.Text,
	// 	"targetLanguageCode": param.From,
	// 	"sourceLanguageCode": param.To,
	// 	"mimeType":           "text/plain",
	// 	"model":              "projects/YOUR_PROJECT_ID/locations/global/models/nmt",
	// }

	for i, v := range param.Text {
		param.Text[i] = strings.ReplaceAll(v, "\"", "")
		param.Text[i] = strings.ReplaceAll(v, "'", "`")
	}

	bodyParams := []interface{}{
		[]interface{}{
			param.Text,
			param.From,
			param.To,
		},
		"te",
	}

	headers := map[string]string{
		"accept":             "*/*",
		"accept-language":    "en-US,en;q=0.9,id;q=0.8",
		"content-type":       "application/json+protobuf",
		"origin":             "https://satudata.kemendag.go.id",
		"priority":           "u=1, i",
		"referer":            "https://satudata.kemendag.go.id/",
		"sec-ch-ua":          `"Chromium";v="136", "Microsoft Edge";v="136", "Not.A/Brand";v="99"`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": `"Windows"`,
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "cross-site",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36 Edg/136.0.0.0",
		"x-goog-api-key":     os.Getenv("GOOGLE_API_KEY"),
	}

	NewBodyParams, _ := json.Marshal(bodyParams)
	// log.Println(aurora.Cyan(string(NewBodyParams)))
	bodyReq, err := utils.GetRequestUseHttps("POST", url, headers, bytes.NewBuffer(NewBodyParams))
	if err != nil {
		response.Message = "Error"
		return
	}

	log.Println(aurora.Cyan(string(bodyReq)))
	response.ResultText = []string{string(bodyReq)}
	return
}

func TranslateText(param model.Translate) (response model.TranslateResponse) {
	projectID := "YOUR_PROJECT_ID"
	ctx := context.Background()
	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
		response.Message = "Error"
		return
	}
	defer client.Close()
	parent := fmt.Sprintf("projects/%s/locations/global", projectID)

	req := &translatepb.TranslateTextRequest{
		Parent:             parent,
		Contents:           param.Text,
		TargetLanguageCode: param.From,
		SourceLanguageCode: param.To,
		MimeType:           "text/plain",
		Model:              "projects/YOUR_PROJECT_ID/locations/global/models/nmt",
	}

	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		response.Message = "Error"
		return
	}

	for i, translation := range resp.GetTranslations() {
		originalText := param.Text[i] // Ambil teks asli dari slice
		fmt.Printf("Asli (%s): %s\n", param.From, originalText)
		fmt.Printf("Terjemahan (%s): %s\n", param.To, translation.GetTranslatedText())
		fmt.Printf("Terdeteksi Bahasa Sumber: %s\n", translation.GetDetectedLanguageCode())
		fmt.Println("------------------------")
	}

	return
}

func LibreTranslate(param model.Translate) (response model.TranslateResponse) {
	url := os.Getenv("URL_API_TRANSLATE_LIBRE")
	var ResultText []string
	for i, _ := range param.Text {
		bodyParams := map[string]interface{}{
			"q":      param.Text[i],
			"source": param.From,
			"target": param.To,
			"format": "text",
		}
		headers := map[string]string{
			"Content-Type": "application/json",
		}

		NewBodyParams, _ := json.Marshal(bodyParams)
		bodyReq, err := utils.GetRequestUseHttps("POST", url, headers, bytes.NewBuffer(NewBodyParams))
		if err != nil {
			response.Message = "Error"
			return
		}
		ResultText = append(ResultText, gjson.Get(string(bodyReq), "translatedText").String())
	}

	response.ResultText = ResultText
	return
}
