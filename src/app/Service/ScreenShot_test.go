package Service

import (
	"context"
	"encoding/json"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestCaptureScreenshotPlus(t *testing.T) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	RegisterContext(ctx)
	var query CapQuery
	content := "{\"url\":\"https://basic.10jqka.com.cn/basicph/industryComparison.html?code=003035&marketid=33&gphonepredraw=true&fontzoom=no\",\"cap_mode\":\"normal\",\"render_strategy\":\"delay\",\"device\":\"IPhoneX\",\"delay\":\"5000\"}"
	json.Unmarshal([]byte(content), &query)
	//var captureByte []byte
	//
	//task := buildTaskParam(query, &captureByte)
	//
	plus, err := CaptureScreenshotPlus(query)
	if err != nil {
		log.Error("cap failed")
	}
	log.Info(plus.ImageUrl)

}
