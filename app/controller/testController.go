package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestController struct {
}

func (this *TestController)Router(r *gin.Engine){
	group := r.Group("/test")
	{
		group.GET("/hello", this.Hello)
		group.GET("/create", this.Create)
		group.GET("/list", this.List)
		group.GET("/delete", this.Delete)
		group.GET("/update", this.Update)
	}
}

func (this *TestController)Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "hello")
	return
}
func (this *TestController)Create(ctx *gin.Context) {

	return
}
func (this *TestController)List(ctx *gin.Context) {

	return
}
func (this *TestController)Delete(ctx *gin.Context) {

	return
}
func (this *TestController)Update(ctx *gin.Context) {

	return
}
