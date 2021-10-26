# web-screenshot-service 

web-screenshot-service 一款基于go语言及chromedp开发的服务端网页截图服务

支持截图模式：
* 可视界面截图
* 网页全屏
* 元素截图


完成标识：
* 延迟时长
* 特定元素加载
* 默认判断

TODO:
高级特性：
* 行为
* 数据断言
* UA定制 
* Cookie
* 浏览器缓存

异步任务队列



高级模式请求示例:


```shell
curl 'http://localhost:1920/screen' -H 'Content-Type: application/json' --data-raw '{"url":"https://basic.10jqka.com.cn/basicph/industryComparison.html?code=003035&marketid=33&gphonepredraw=true&fontzoom=no","cap_mode":"normal","render_strategy":"default","device":"IPhoneX"}'  --compressed


curl "https://localhost:1920/screen/plus" ^
  -X "OPTIONS" ^
  -H "Accept: */*" ^
  -H "Access-Control-Request-Method: POST" ^
  -H "Access-Control-Request-Headers: content-type" ^
  -H "Origin: https://app.werty.cn" ^
  -H "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 Edg/92.0.902.84" ^
  -H "Sec-Fetch-Mode: cors" ^
  --compressed



curl 'https://app.werty.cn/screen/plus/async' \
  -H 'Connection: keep-alive' \
  -H 'sec-ch-ua: "Chromium";v="92", " Not A;Brand";v="99", "Microsoft Edge";v="92"' \
  -H 'Accept: application/json, text/plain, */*' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 Edg/92.0.902.84' \
  -H 'Content-Type: application/json' \
  -H 'Origin: https://app.werty.cn' \
  -H 'Sec-Fetch-Site: same-origin' \
  -H 'Sec-Fetch-Mode: cors' \
  -H 'Sec-Fetch-Dest: empty' \
  -H 'Referer: https://app.werty.cn/app/v2/index.html?app=screenshot_async' \
  -H 'Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6' \
  --data-raw '{"url":"https://basic.10jqka.com.cn/basicph/industryComparison.html?code=003035&marketid=33&gphonepredraw=true&fontzoom=no","cap_mode":"normal","render_strategy":"delay","device":"IPhoneX","delay":"5000"}' \
  --compressed
```