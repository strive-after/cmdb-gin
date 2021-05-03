package main

import (
	"github.com/gin-gonic/gin"

	"gin-moudle/internal/config"
	"gin-moudle/internal/router"
	"gin-moudle/pkg/log"
	"gin-moudle/pkg/mongo"
)

func main() {
	g := gin.New()

	err := router.Reg(g)
	if err != nil {
		panic(err)
	}

	//加载配置文件
	c, err := config.InitConfig("../configs/config.toml")
	if err != nil {
		panic(err)
	}

	//加载日志配置
	logCall := log.InitLoger(&c.Log)
	logCall()

	//初始化mongodb
	mongo.InitMongo(&c.Mongo)

	//初始化redis

	//初始化携程组

	g.Run(":8080")
}
