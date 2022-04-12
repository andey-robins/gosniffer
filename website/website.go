package website

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/andey-robins/gosniffer/sniffer"
)

func Main() {
	fileServer := http.FileServer(http.Dir("./website/css"))
	http.Handle("/resources/css/cyberpunk.css", http.StripPrefix("/resources/css", fileServer))
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	p := filepath.Join(cwd, "website", "index-tmpl.html")
	t := template.Must(template.ParseFiles(p))

	Networks := []sniffer.NetworkNode{
		{
			Speed:          5.0,
			SignalStrength: -10,
			SSID:           "NSA Van #8",
		},
		{
			Speed:          51.0,
			SignalStrength: -20,
			SSID:           "CEDAR_IOT",
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
