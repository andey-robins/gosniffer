package sniffer

type NetworkNode struct {
	StationMac  string `json:"mac"`
	Power       string `json:"power"`
	PacketCount string `json:"packetCount"`
	BSSID       string `json:"bssid"`
	ESSID       string `json:"essid"`
}
