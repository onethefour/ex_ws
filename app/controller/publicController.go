package controller

import (
	"github.com/gin-gonic/gin"
)

type PublicController struct {
}

func (this *PublicController)Router(r *gin.Engine){
	r.Static("/public",  "./public")
}
