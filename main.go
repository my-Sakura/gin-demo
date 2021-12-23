package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/my-Sakura/gin-demo/model_name/controller"
	"github.com/my-Sakura/gin-demo/utils"
)

const (
	_raftGroup  = "/api/v1"
	_serverAddr = "0.0.0.0:5000"
)

func main() {
	gin.SetMode(gin.DebugMode)

	engine := gin.Default()
	engine.Use(utils.Cors())

	c := controller.New()
	if err := c.Regist(engine.Group(_raftGroup)); err != nil {
		panic(err)
	}

	log.Printf("Error run server: %v\n", engine.Run(_serverAddr))
}
