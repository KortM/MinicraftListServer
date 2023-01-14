package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	//Create routing
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Server started")

	fileServer := http.FileServer(http.Dir("../../ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	err := http.ListenAndServe(":60000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

//HOME page
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{

		"../../ui/templates/index.page.html",
		"../../ui/templates/servers.page.html",
		"../../ui/templates/base.page.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	//w.Write([]byte("Test"))
}
