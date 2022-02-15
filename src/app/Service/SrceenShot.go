package Service

import (
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
	"web-screenshot-service/lib/util"
)

func setOutputDir(dir string) {
	log.Info("register output dir is " + dir)
	outputDir = dir
	_, err := os.Stat(dir)
	if err == nil || os.IsExist(err) {
		return
	}
	log.Infof("The directory [%s] does not exist, it will be created automatically", dir)

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Errorf("Automatic directory [%s] creation failed: %s", dir, err.Error())

		log.Error(err)
		panic(err)
		os.Exit(127)
	}

}

type ScreenshotRes struct {
	ImageUrl string `json:"image_url,omitempty"`
}

type CapQuery struct {
	// 请求地址
	Url string `json:"url"`
	// 设备
	Device string `json:"device"`
	// 截图模式 full/element/normal  默认normal
	CapMode string `json:"cap_mode"`
	// 截图元素选择器
	CapElement string `json:"cap_element"`
	// 渲染策略
	RenderStrategy string `json:"render_strategy"`
	// 渲染元素选择器
	RenderElement string `json:"render_element"`
	// 等待渲染延迟时长
	RenderDelay int64 `json:"render_delay"`
}

func CaptureScreenshot(url string, deviceName string) (ScreenshotRes, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	filename := util.GetMd5(url+timestamp) + ".png"
	savePath := outputDir + filename
	device := GetDevice(deviceName)
	log.WithFields(log.Fields{"url": url, "savePath": savePath, "device_name": deviceName, "device": device.Device().String()}).Info("request Capture Screenshot")
	var res ScreenshotRes
	var captureByte []byte
	if err := chromedp.Run(
		GetChromeContext(),
		chromedp.Emulate(device),
		chromedp.Navigate(url),
		chromedp.CaptureScreenshot(&captureByte),
	); err != nil {
		log.WithFields(log.Fields{"url": url, "savePath": savePath, "device_name": deviceName, "device": device.Device().String()}).Error(err)
		return res, err
	}
	if err := ioutil.WriteFile(savePath, captureByte, 0777); err != nil {
		log.WithFields(log.Fields{"url": url, "savePath": savePath}).Error(err)
		return res, err
	}
	res.ImageUrl = formatUrl(filename)
	return res, nil
}

func CaptureScreenshotPlus(query CapQuery) (ScreenshotRes, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	url := query.Url
	taskId := util.GetMd5(url + timestamp)
	return doScreenshotPlus(query, taskId)
}

// AsyncCaptureScreenshotPlus 异步调用
func AsyncCaptureScreenshotPlus(query CapQuery) (TaskResp, error) {
	task, err := CreateTask(query)
	if err != nil {
		var emptyRes TaskResp
		return emptyRes, err
	}
	return task, nil
}

func doScreenshotPlus(query CapQuery, taskId string) (ScreenshotRes, error) {
	url := query.Url
	filename := taskId + ".png"
	savePath := outputDir + filename
	deviceName := query.Device
	device := GetDevice(deviceName)
	log.WithFields(log.Fields{"url": url, "savePath": savePath, "device_name": deviceName, "device": device.Device().String()}).Info("request Capture Screenshot")
	var res ScreenshotRes
	var captureByte []byte
	context, cancel := GetChromeTimeContext()
	defer cancel()
	if err := chromedp.Run(
		context,
		buildTask(query, &captureByte),
	); err != nil {
		log.WithFields(log.Fields{"url": url, "savePath": savePath, "device_name": deviceName, "device": device.Device().String()}).Error(err)
		return res, err
	}

	if err := ioutil.WriteFile(savePath, captureByte, 0777); err != nil {
		log.WithFields(log.Fields{"url": url, "savePath": savePath}).Error(err)
		return res, err
	}

	image, err := doImageUpload(savePath)
	if err != nil {
		return res, err
	}

	res.ImageUrl = image
	return res, nil
}

func doImageUpload(filePath string) (string, error) {
	// 图片上传资源服务器
	return UploadImage(filePath)
}

func buildTask(query CapQuery, captureByte *[]byte) chromedp.Tasks {
	tasks := chromedp.Tasks{}
	// UA 配置
	tasks = append(tasks, chromedp.Emulate(GetDevice(query.Device)))
	// 渲染策略任务
	tasks = append(tasks, chromedp.Navigate(query.Url))
	if query.RenderStrategy == "element" {
		tasks = append(tasks, chromedp.WaitVisible(query.RenderElement, chromedp.ByQuery))
	}
	if query.RenderStrategy == "delay" {
		log.Infof("delay sleep %s ms ...", query.RenderDelay)
		duration := time.Duration(query.RenderDelay) * time.Millisecond
		tasks = append(tasks, chromedp.Sleep(duration))
	}
	// 截图模式
	if query.CapMode == "element" {
		tasks = append(tasks, chromedp.WaitVisible(query.CapElement, chromedp.ByQuery))
		tasks = append(tasks, chromedp.Screenshot(query.CapElement, captureByte))
	}
	if query.CapMode == "full" {
		tasks = append(tasks, chromedp.FullScreenshot(captureByte, 100))
	}
	if query.CapMode == "normal" || query.CapMode == "default" {
		tasks = append(tasks, chromedp.CaptureScreenshot(captureByte))
	}

	return tasks

}

func formatUrl(path string) string {
	return domain + route + path
}
