package api

import (
	"encoding/json"
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

// GetStartup is invoked by the /startup endpoint.
func GetStartup(c *gin.Context) {
	db := db.Connect()
	rows, err := db.Db.Query("SELECT * FROM networks;")
	defer rows.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp := make([]sniffer.NetworkNode, 0)

	for rows.Next() {
		var node sniffer.NetworkNode
		var uid int
		err := rows.Scan(&uid, &node.StationMac, &node.Power, &node.PacketCount, &node.BSSID, &node.ESSID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		resp = append(resp, node)
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Data(http.StatusOK, "application/json", data)
}

// PostRegister is invoked by the /register endpoint. Sending a json representation of a data struct
// will dump it into the database
func PostRegister(c *gin.Context) {
	var registration sniffer.NetworkNode
	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
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
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "received",
	})
}
