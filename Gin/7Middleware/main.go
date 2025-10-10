package main

import (
	"Gin/7Middleware/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routers.DefaultRoutersInit(router)
	routers.ApiRoutersInit(router)

	router.Run()
}
