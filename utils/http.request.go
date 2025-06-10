package utils

import (
	"io"
	"log"
	"net/http"

	"github.com/logrusorgru/aurora"
)

func GetRequestUseHttps(method string, url string, headers map[string]string, bodyParams io.Reader) (bodyResp []byte, err error) {
	// log.Println(aurora.Cyan(bodyParams))
	log.Println(aurora.Cyan(method), aurora.Yellow(url), "\n", aurora.Magenta(bodyParams))

	req, err := http.NewRequest(method, url, bodyParams)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyResp, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return
}
