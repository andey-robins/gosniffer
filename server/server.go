package server

import (
	"github.com/andey-robins/gosniffer/api"
	"github.com/gin-gonic/gin"
)

func Main() {
	r := gin.Default()

	r.GET("/", api.GetRoot)
	r.GET("/status", api.GetStatus)
	r.GET("/search", api.GetSearch)
	r.POST("/register", api.PostRegister)

	r.Run()
}
