package Service

import (
	"fmt"
	"os"
	"strings"
	"web-screenshot-service/lib"
	"web-screenshot-service/lib/conf"
)

var outputDir string
var domain string
var route string

func RegisterConf() {
	setOutputDir(conf.Conf.String("gin::output_dir"))
	domain = conf.Conf.String("gin::domain")
	route = conf.Conf.String("gin::output_route")

	initRedis()
	initUpload()

	InitTask(
		conf.Conf.DefaultInt("task::consumer_number", 30),
		conf.Conf.DefaultInt("task::wait_queue_size", 10000),
	)
}

func initRedis() {
	var redisConfig lib.ClientConfig
	redisConfig.Ip = conf.Conf.String("redis::ip")
	redisConfig.Mode = conf.Conf.String("redis::mode")
	redisConfig.Port = conf.Conf.String("redis::port")
	redisConfig.Passwd = conf.Conf.String("redis::auth")
	lib.CreateRedisClientByConfig(redisConfig)
}

func initUpload() {
	var uploadConfig UploadConfig
	uploadConfig.Mode = conf.Conf.String("upload::mode")
	uploadConfig.Ip = conf.Conf.String("upload::ip")
	uploadConfig.Auth = conf.Conf.String("upload::auth")
	uploadConfig.Domain = conf.Conf.String("upload::domain")
	uploadConfig.Uri = conf.Conf.String("upload::uri")
	uploadConfig.AppCode = conf.Conf.String("upload::app_code")
	uploadConfig.AppName = conf.Conf.String("upload::app_name")
	if uploadConfig.Mode == "sidecar" {
		uploadConfig.Ip = getServerIp(uploadConfig.Domain)
	}
	setUploadConfig(uploadConfig)
}

func getServerIp(domain string) string {
	domain = strings.Replace(domain, ".", "_", -1)
	domain = strings.Replace(domain, "-", "_", -1)
	domain = fmt.Sprintf("ARSENAL_SVC_%s_HTTP_HOST", strings.ToUpper(domain))
	return os.Getenv(domain)
}
