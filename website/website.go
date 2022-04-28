package website

import (
	"encoding/json"
	"io/ioutil"
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
	log.Println("Started listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	p := filepath.Join(cwd, "website", "index-tmpl.html")
	t := template.Must(template.ParseFiles(p))

	resp, err := http.Get("http://localhost:8080/startup")
	if err != nil {
		log.Printf("Error in request to home server: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error in reading body: %v\n", err)
		return
	}

	log.Println(string(body))

	Networks := make([]sniffer.NetworkNode, 0)
	err = json.Unmarshal(body, &Networks)
	if err != nil {
		log.Printf("Error in unmarshalling json: %v\n", err)
		return
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
