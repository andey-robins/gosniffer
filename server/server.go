package server

import (
	"github.com/andey-robins/gosniffer/api"
	"github.com/gin-gonic/gin"
)

func Main() {
	r := gin.Default()

	r.GET("/status", api.GetStatus)

	r.Run()
}
