package app

import (
	"github.com/gin-gonic/gin"
	"newbtc-ws/app/controller"
)
func Router(r *gin.Engine){
	new(controller.PublicController).Router(r)
	new(controller.TestController).Router(r)
	new(controller.WebsocketController).Router(r)
}