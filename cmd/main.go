package main

import (
	"fmt"
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
	g.Run(":8080")
}
