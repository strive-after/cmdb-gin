package api

import (
	"encoding/json"
	"gin-moudle/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type testhandler struct{}

func Newtest() *testhandler {
	return &testhandler{}
}

type Res struct {
	Code   int
	Msg    string
	Status int
}

func (t *testhandler) Add(ctx *gin.Context) {
	a := &Res{
		Code:   1,
		Msg:    "test",
		Status: 200,
	}
	d, err := json.Marshal(a)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}
	ctx.JSON(http.StatusOK, string(d))
}
