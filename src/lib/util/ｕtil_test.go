package util

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestUpload(t *testing.T) {
	var files = make(map[string]string)
	var param = make(map[string]string)
	var header = make(map[string]string)
	files["upfile"] = "/home/hanjinxiang@myhexin.com/image/2021-08-28_14-12.png"
	header["X-Arsenal-Auth"] = "arsenal-tools"
	header["HOST"] = "sns-space.base"
	param["appCode"] = "0"
	param["appName"] = "sns"
	upload, err := RequestUploadFiles("post", "http://10.0.7.249/newupload/upload/", param, header, files)
	if err != nil {
		t.Error(err)
	}
	log.Info(string(upload))
}
