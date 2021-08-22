package Service

import (
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
	"web-srceenshot-service/lib/util"
)

var outputDir string = "screenshot/"

func CaptureScreenshot(url string) (string, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	path := outputDir + util.GetMd5(url+timestamp) + ".png"
	log.WithFields(log.Fields{"url": url, "path": path}).Info("request Capture Screenshot")

	var captureByte []byte
	if err := chromedp.Run(
		GetChromeContext(),
		chromedp.Emulate(device.IPhone8Plus),
		chromedp.Navigate(url),
		chromedp.CaptureScreenshot(&captureByte),
	); err != nil {
		log.WithFields(log.Fields{"url": url, "path": path}).Error(err)
		return "", err
	}
	if err := ioutil.WriteFile(path, captureByte, 0777); err != nil {
		log.WithFields(log.Fields{"url": url, "path": path}).Error(err)
		return "", err

	}

	return path, nil
}
