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
	content := "{\"url\":\"https://basic.10jqka.com.cn/basicph/briefinfo.html?fromshare=1&code=300033&market_id=17&fontzoom=no&client_userid=xn8rh&share_hxapp=gsc&share_action=webpage_share.1&back_source=wxhy#/company\",\"cap_mode\":\"element\",\"render_strategy\":\"delay\",\"devic\":\"IPhoneX\",\"cap_element\":\"#scrollDom > div.business-analysis.basic-bgColor\",\"delay\":\"10000\",\"render_element\":\"#scrollDom > div.business-analysis.basic-bgColor\"}"
	json.Unmarshal([]byte(content), &query)
	//var captureByte []byte
	//
	//task := buildTask(query, &captureByte)
	//
	plus, err := CaptureScreenshotPlus(query)
	if err != nil {
		log.Error("cap failed")
	}
	log.Info(plus.ImageUrl)

}
