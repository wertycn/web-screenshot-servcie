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

```