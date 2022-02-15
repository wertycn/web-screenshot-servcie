package util

import (
	"bytes"
	"fmt"
	io "io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func RequestUploadFiles(method, url string, params, headers, files map[string]string) ([]byte, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	for fieldName, filePath := range files {
		file, _ := os.Open(filePath)
		formFile, _ := writer.CreateFormFile(fieldName, filePath)
		io.Copy(formFile, file)
		file.Close()
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(strings.ToUpper(method), url, body)
	if err != nil {
		//log.Error("NewRequest err: %v, url: %s, body: %v", err, url, body)
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	for key, val := range headers {
		if strings.ToUpper(key) == "HOST" {
			req.Host = val
		} else {
			req.Header.Set(key, val)
		}
	}
	resp, err := c.Do(req)
	if err != nil {
		//log.Error("Do Request got err: %v, req: %v, resp: %v", err, req, resp)
		return nil, err
	}

	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	all, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return all, fmt.Errorf("status code is %d", resp.StatusCode)
	}
	return all, err
}
