package main

import (
	"Gin/9FileUpload/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/**/*")

	routers.DefaultRoutersInit(router)
	routers.ApiRoutersInit(router)

	router.Run()
}
