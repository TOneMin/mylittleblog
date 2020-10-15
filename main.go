package main

import (
	"littleblog/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	routers.InitializeRouters(app)
	app.Run()
}
