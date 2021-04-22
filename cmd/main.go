package main

import (
	"fmt"
	"gin-moudle/internal/config"
	"gin-moudle/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()
	err := router.Reg(g)
	if err != nil {
		fmt.Println(err)
		return
	}
	//加载配置文件
	_, err = config.InitConfig("../configs/config.toml")
	if err != nil {
		fmt.Println(err)
		return
	}

	g.Run(":8080")
}
