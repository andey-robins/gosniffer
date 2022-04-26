package website

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/andey-robins/gosniffer/sniffer"
)

func Main() {
	fileServer := http.FileServer(http.Dir("./website/css"))
	http.Handle("/resources/css/cyberpunk.css", http.StripPrefix("/resources/css", fileServer))
	http.HandleFunc("/", mainHandler)
	log.Println("Started listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	p := filepath.Join(cwd, "website", "index-tmpl.html")
	t := template.Must(template.ParseFiles(p))

	Networks := []sniffer.NetworkNode{
		{
			StationMac:  "01:23:45:67:89:cd",
			FirstSeen:   time.Unix(10, 10),
			LastSeen:    time.Unix(10, 10),
			Power:       -8,
			PacketCount: 12,
			BSSID:       "(not associated)",
			ESSID:       "NSA Van #8",
		},
		{
			StationMac:  "01:23:45:67:89:ab",
			FirstSeen:   time.Unix(10, 10),
			LastSeen:    time.Unix(10, 10),
			Power:       -10,
			PacketCount: 10,
			BSSID:       "(not associated)",
			ESSID:       "NSA Van #9",
		},
		{
			StationMac:  "01:23:d5:67:89:cd",
			FirstSeen:   time.Unix(10, 10),
			LastSeen:    time.Unix(10, 10),
			Power:       -8,
			PacketCount: 12,
			BSSID:       "(not associated)",
			ESSID:       "NSA Van #8",
		},
		{
			StationMac:  "01:f3:45:67:89:ab",
			FirstSeen:   time.Unix(10, 10),
			LastSeen:    time.Unix(10, 10),
			Power:       -10,
			PacketCount: 10,
			BSSID:       "(not associated)",
			ESSID:       "NSA Van #9",
		},
		{
			StationMac:  "01:23:45:c7:89:cd",
			FirstSeen:   time.Unix(10, 10),
			LastSeen:    time.Unix(10, 10),
			Power:       -8,
			PacketCount: 12,
			BSSID:       "(not associated)",
			ESSID:       "NSA Van #8",
		},
		{
			StationMac:  "01:23:45:a7:89:ab",
			FirstSeen:   time.Unix(10, 10),
			LastSeen:    time.Unix(10, 10),
			Power:       -10,
			PacketCount: 10,
			BSSID:       "(not associated)",
			ESSID:       "NSA Van #9",
		},
	}

	mStruct := make(map[string]any)
	mStruct["Networks"] = Networks

	// w.Write([]byte(p))

	if err := t.Execute(w, mStruct); err != nil {
		log.Print(err.Error())
	}
}

func styleSheet(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	p := filepath.Join(cwd, "website", "cyberpunk.css")
	t := template.Must(template.ParseFiles(p))

	w.Header().Set("content-type", "text/css")
	if err := t.Execute(w, nil); err != nil {
		log.Print(err.Error())
	}
}
