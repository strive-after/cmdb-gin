package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"gin-moudle/internal/config"
	"gin-moudle/internal/router"
	"gin-moudle/pkg/log"
	"gin-moudle/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type Tes struct {
	Name string
	Age  int
	Addr string
}

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

	m := mongo.Newm()
	data := make([]Tes, 0, 0)
	err = m.FindMany("testaa", 1, 0, bson.M{}, &data)
	fmt.Println(err, data)

	//初始化携程组

	g.Run(":8080")
}
