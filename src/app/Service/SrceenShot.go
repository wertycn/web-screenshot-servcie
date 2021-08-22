package Service

import (
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
	"web-srceenshot-service/lib/conf"
	"web-srceenshot-service/lib/util"
)

var outputDir string
var domain string
var route string

func RegisterConf() {
	setOutputDir(conf.Conf.String("gin::output_dir"))
	domain = conf.Conf.String("gin::domain")
	route = conf.Conf.String("gin::output_route")
}
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
	Device   string `json:"device,omitempty"`
}

func CaptureScreenshot(url string) (ScreenshotRes, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	filename := util.GetMd5(url+timestamp) + ".png"
	savePath := outputDir + filename
	log.WithFields(log.Fields{"url": url, "savePath": savePath}).Info("request Capture Screenshot")
	var res ScreenshotRes

	var captureByte []byte
	if err := chromedp.Run(
		GetChromeContext(),
		chromedp.Emulate(device.IPhone8Plus),
		chromedp.Navigate(url),
		chromedp.CaptureScreenshot(&captureByte),
	); err != nil {
		log.WithFields(log.Fields{"url": url, "savePath": savePath}).Error(err)
		return res, err
	}
	if err := ioutil.WriteFile(savePath, captureByte, 0777); err != nil {
		log.WithFields(log.Fields{"url": url, "savePath": savePath}).Error(err)
		return res, err

	}
	res.ImageUrl = formatUrl(filename)

	res.Device = device.IPhone8Plus.String()
	return res, nil
}

func formatUrl(path string) string {
	return domain + route + path
}
