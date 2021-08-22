package Service

import (
	"context"
)

var ctx context.Context

func RegisterContext(c context.Context) {
	ctx = c
}

func GetChromeContext() context.Context {
	return ctx
}
