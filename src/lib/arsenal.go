package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
	"web-screenshot-service/lib/util"
)

type ArsenalApi struct {
	Domain string
	Uri    string
	Auth   string
	Host   string
}

type RequestOption struct {
	Timeout time.Duration
	Header  map[string]string
}

type ArsenalResponse struct {
	StatusCode int             `json:"status_code"`
	StatusMsg  string          `json:"status_msg"`
	Data       json.RawMessage `json:"data"`
	DataList   json.RawMessage `json:"datalist"`
	Result     json.RawMessage `json:"result"`
	Datas      json.RawMessage `json:"datas"`
}

func GetServerIp(domain string) string {
	domain = strings.Replace(domain, ".", "_", -1)
	domain = strings.Replace(domain, "-", "_", -1)
	domain = fmt.Sprintf("ARSENAL_SVC_%s_HTTP_HOST", strings.ToUpper(domain))
	return os.Getenv(domain)
}

func BuildHeader(domain string, auth string) map[string]string {
	var header = make(map[string]string)
	header["X-Arsenal-Auth"] = auth
	header["HOST"] = domain
	return header
}

func RequestPost(url string, param, header map[string]string) ([]byte, error) {
	return util.NewRequest("post", url, header, param)
}
