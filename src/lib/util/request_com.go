package util

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func equestUploadFile(method, url string, params, headers, files map[string]string) ([]byte, error) {
	//bodyBuffer := &bytes.Buffer{}
	bodyBuffer := new(bytes.Buffer)
	bodyWriter := multipart.NewWriter(bodyBuffer)
	// build param
	for key, value := range params {
		_ = bodyWriter.WriteField(key, value)
	}
	// build file
	//for key, value := range files {
	fileWriter, _ := bodyWriter.CreateFormFile("upfile", files["upfile"])
	file, err := os.Open(files["upfile"])
	if err != nil {
		return nil, err
	}
	io.Copy(fileWriter, file)
	defer file.Close()
	//}

	contentType := bodyWriter.FormDataContentType()
	defer bodyWriter.Close()
	req, err := http.NewRequest("POST", url, bodyBuffer)
	if err != nil {
		return nil, err
	}
	// build header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", contentType)
	var client = &http.Client{}
	resp, err := client.Do(req)
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
