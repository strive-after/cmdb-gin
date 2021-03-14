package router

import (
	"gin-moudle/internal/api"
	"github.com/gin-gonic/gin"
)

func RegTest(engine *gin.Engine) error {
	g := engine.Group("/rest")
	v1 := g.Group("/test")
	{
		handler := api.Newtest()
		v1.GET("/abc", handler.Add)
	}
	return nil
}
