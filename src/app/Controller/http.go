package Controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"web-screenshot-service/app/Service"
)

// Resp 接口响应对象
type Resp struct {
	Code int         `json:"status_code"`
	Msg  string      `json:"status_msg,omitempty"`
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
	device := c.Request.FormValue("device")
	if device == "DEFAULT" || device == "" {
		device = "default"
	}

	screenshot, err := Service.CaptureScreenshot(url, device)
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

func ScreenShotsPlus(c *gin.Context) {
	var resp Resp
	resp.Code = -1

	var query Service.CapQuery
	err := c.ShouldBindJSON(&query)
	log.Info(query)
	if err != nil {
		resp.Msg = "参数解析错误:" + err.Error()
		c.JSON(200, resp)
		return
	}

	screenshot, err := Service.CaptureScreenshotPlus(query)
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

func AsyncScreenShotsPlus(c *gin.Context) {
	var resp Resp
	resp.Code = -1

	var query Service.CapQuery
	err := c.ShouldBindJSON(&query)
	log.Info(query)
	if err != nil {
		resp.Msg = "参数解析错误:" + err.Error()
		c.JSON(200, resp)
		return
	}

	screenshot, err := Service.AsyncCaptureScreenshotPlus(query)
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

func GetTaskResp(c *gin.Context) {
	taskId := c.Request.FormValue("task_id")
	var resp Resp
	resp.Code = -1

	if taskId == "" {
		resp.Msg = "taskId cannot be empty..."
		c.JSON(200, resp)
		return
	}

	taskResp, exist := Service.GetTaskResp(taskId)
	if exist == false {
		resp.Msg = "taskId is not exist"
		c.JSON(200, resp)
		return
	}
	resp.Code = 0
	status := taskResp.Status
	resp.Data = taskResp
	if status == "wait" || status == "process" {
		c.JSON(202, resp)
		return
	}
	c.JSON(200, resp)

}

func GetDeviceList(c *gin.Context) {
	var resp Resp

	resp.Code = 0
	resp.Msg = "success"
	resp.Data = Service.GetDeviceList()
	c.JSON(200, resp)
}
