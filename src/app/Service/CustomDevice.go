package Service

import "github.com/chromedp/chromedp/device"

var devices = [...]device.Info{
	{"", "", 0, 0, 0.000000, false, false, false},
	{"同花顺安卓普通版", "Mozilla/5.0 (PlayBook; U; RIM Tablet OS 2.1.0; en-US) AppleWebKit/536.2+ (KHTML like Gecko) Version/7.2.1.0 Safari/536.2+", 600, 1024, 1.000000, false, true, true},
	{"同花顺安卓至尊版", "Mozilla/5.0 (PlayBook; U; RIM Tablet OS 2.1.0; en-US) AppleWebKit/536.2+ (KHTML like Gecko) Version/7.2.1.0 Safari/536.2+", 1024, 600, 1.000000, true, true, true},
	{"同花顺IOS至尊版", "Mozilla/5.0 (BB10; Touch) AppleWebKit/537.10+ (KHTML, like Gecko) Version/10.0.9.2372 Mobile Safari/537.10+", 360, 640, 2.000000, false, true, true},
	{"同花顺IOS普通版", "Mozilla/5.0 (BB10; Touch) AppleWebKit/537.10+ (KHTML, like Gecko) Version/10.0.9.2372 Mobile Safari/537.10+", 640, 360, 2.000000, true, true, true},
}
