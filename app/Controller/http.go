package Controller

import (
	"github.com/gin-gonic/gin"
	"web-srceenshot-service/app/Service"
)

/**
 * 定义Http接口JOSN响应数据结构
 */
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func ScreenShots(c *gin.Context) {
	url := c.Request.FormValue("url")
	var resp Resp
	resp.Code = -1

	if url == "" {
		resp.Msg = "url cannot be empty..."
		c.JSON(200, resp)
		return
	}

	screenshot, err := Service.CaptureScreenshot(url)
	if err != nil {
		resp.Msg = "capture screenshot failed:" + err.Error()
		c.JSON(200, resp)
		return
	}

	resp.Code = 0
	resp.Msg = "success"
	resp.Data = screenshot
	c.JSON(200, resp)

}
