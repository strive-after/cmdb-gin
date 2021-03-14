package api

import "github.com/gin-gonic/gin"

type testhandler struct{}

func Newtest() *testhandler {
	return &testhandler{}
}

func (t *testhandler) Add(ctx *gin.Context) {
	ctx.JSON(200, "ok")
}
