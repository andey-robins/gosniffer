package api

import "github.com/gin-gonic/gin"

// GetRoot is invoked at the / endpoint. Currently it just performs a status check
func GetRoot(c *gin.Context) {
	GetStatus(c)
}

// GetStatus is invoked anytime /status is called. It returns a json struct like {"status": "ok"}
func GetStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// GetSearch is invoked by the /search endpoint. TODO: Write this
func GetSearch(c *gin.Context) {

}

// PostRegister is invoked by the /register endpoint. TODO: Write this
func PostRegister(c *gin.Context) {

}
