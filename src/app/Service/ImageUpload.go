package Service

import (
	"encoding/json"
	"errors"
	"fmt"
	"web-screenshot-service/lib/util"
)

var uploadConfig UploadConfig

type UploadConfig struct {
	Mode    string
	Auth    string
	Ip      string
	Uri     string
	Domain  string
	AppName string
	AppCode string
}

type UploadResponse struct {
	ErrorCode int          `json:errorCode`
	ErrorMsg  string       `json:errorMsg`
	Result    UploadResult `json:result`
}

type UploadResult struct {
	Thurl string `json:thurl`
	Url   string `json:url`
}

func setUploadConfig(config UploadConfig) {
	uploadConfig = config
}

func UploadImage(file string) (string, error) {
	var files = make(map[string]string)
	files["upfile"] = file
	resp, err := util.RequestUploadFiles("POST", getUploadUrl(), buildParam(), buildFormHeader(), files)
	if err != nil {
		return "", err
	}
	uploadResp, status := decodeUploadResp(resp)
	if status != nil {
		return "", err
	}
	if uploadResp.ErrorCode != 0 {
		return "", errors.New(uploadResp.ErrorMsg)
	}
	return uploadResp.Result.Url, nil
}

func getUploadUrl() string {
	return fmt.Sprintf("http://%s%s", uploadConfig.Ip, uploadConfig.Uri)
}

func buildFormHeader() map[string]string {
	var header = make(map[string]string)
	header["X-Arsenal-Auth"] = uploadConfig.Auth
	header["HOST"] = uploadConfig.Domain
	return header
}

func buildParam() map[string]string {
	var param = make(map[string]string)
	param["appCode"] = uploadConfig.AppCode
	param["appName"] = uploadConfig.AppName
	return param
}

func decodeUploadResp(resp []byte) (UploadResponse, error) {
	var uploadResponse UploadResponse
	err := json.Unmarshal(resp, &uploadResponse)
	if err != nil {
		return uploadResponse, err
	}
	return uploadResponse, nil
}
