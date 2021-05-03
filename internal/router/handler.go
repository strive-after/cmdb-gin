package router

import (
	"github.com/gin-gonic/gin"
)

type Action func(engine *gin.Engine) error

var (
	routerAction = []Action{
		RegTest,
	}
)

func Reg(engine *gin.Engine) error {
	for _, action := range routerAction {
		if err := action(engine); err != nil {
			return err
		}
	}
	return nil
}
