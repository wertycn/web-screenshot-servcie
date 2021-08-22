package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"web-srceenshot-service/app/Controller"
	"web-srceenshot-service/app/Service"
	conf "web-srceenshot-service/lib/conf"
)

func main() {
	log.Info("web screenshots service start ...")
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	Service.RegisterContext(ctx)
	log.Info("chromeDP init complete")
	conf.LoadConfig("conf/app.ini")
	Service.RegisterConf()

	log.Info("gin route register start ...")
	r := gin.Default()
	r.Use(Cors())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	r.GET("/screen", Controller.ScreenShots)
	r.StaticFS(conf.Conf.String("gin::output_route"), http.Dir(conf.Conf.String("gin::output_dir")))
	r.Run("0.0.0.0:1920")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		// 可将将* 替换为指定的域名
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
