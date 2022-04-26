package sniffer

import "time"

type NetworkNode struct {
	StationMac  string    `csv:"Station MAC"`
	FirstSeen   time.Time `csv:"First time seen"`
	LastSeen    time.Time `csv:"Last time seen"`
	Power       int       `csv:"Power"`
	PacketCount int       `csv:"# packets"`
	BSSID       string    `csv:"BSSID"`
	ESSID       string    `csv:"Probed ESSIDs"`
}
