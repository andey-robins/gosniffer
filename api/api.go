package api

import (
	"log"
	"net/http"

	"github.com/andey-robins/gosniffer/db"
	"github.com/andey-robins/gosniffer/sniffer"
	"github.com/gin-gonic/gin"
)

// GetRoot is invoked at the / endpoint. Currently it just performs a status check
func GetRoot(c *gin.Context) {
	GetStatus(c)
}

// GetStatus is invoked anytime /status is called. It returns a json struct like {"status": "ok"}
func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetSearch is invoked by the /search endpoint. TODO: Write this
func GetSearch(c *gin.Context) {

}

// PostRegister is invoked by the /register endpoint. TODO: Write this
func PostRegister(c *gin.Context) {
	var registration sniffer.NetworkNode
	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	db := db.Connect()
	if _, err := db.Db.Exec(`INSERT INTO networks (mac, power, packetCount, bssid, essid) VALUES ('` +
		registration.StationMac + `', '` +
		registration.Power + `', '` +
		registration.PacketCount + `', '` +
		registration.BSSID + `', '` +
		registration.ESSID + `');`,
	); err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "received",
	})
}
