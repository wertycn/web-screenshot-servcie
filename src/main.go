package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"web-screenshot-service/app/Controller"
	"web-screenshot-service/app/Service"
	conf "web-screenshot-service/lib/conf"
)

func main() {
	log.Info("web screenshots service start ...")

	log.Info("chromeDP init complete")
	conf.LoadConfig("conf/app.ini")
	Service.RegisterConf()

	log.Info("gin route register start ...")
	r := gin.Default()
	r.Use(Cors())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	r.GET("/screen/sync", Controller.ScreenShots)
	r.POST("/screen/plus/sync", Controller.ScreenShotsPlus)
	r.POST("/screen/plus/async", Controller.AsyncScreenShotsPlus)
	r.GET("/screen/plus/async_res", Controller.GetTaskResp)
	r.GET("/screen/device", Controller.GetDeviceList)
	r.StaticFS(conf.Conf.String("gin::output_route"), http.Dir(conf.Conf.String("gin::output_dir")))
	r.Run("0.0.0.0:1920")
}

/*
 * 跨域配置
 */
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 动态跨域支持
		c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("origin"))
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
