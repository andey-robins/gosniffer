package website

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

func Main() {
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	p := filepath.Join(cwd, "website", "index.html")
	t := template.Must(template.ParseFiles(p))

	if err := t.Execute(w, nil); err != nil {
		log.Print(err.Error())
	}
}
