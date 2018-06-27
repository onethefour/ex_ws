package main

import (
	"github.com/gin-gonic/gin"
	"newbtc-ws/app"
	"fmt"
)

func main(){

	router := gin.Default()
	app.Router(router)
	router.Run(fmt.Sprintf(":%d", 8000))

}