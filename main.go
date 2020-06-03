package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/video/get/{id}", handleVideo)
	r.HandleFunc("/article/get/{id}", handleArticle)
	r.HandleFunc("/video/bycategory/{cat}", handleCategory)
	r.HandleFunc("/article/bycategory/{cat}", handleNewsCategory)
	r.HandleFunc("/news/homepage", handleNewsHomepage)

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT must be set")
		port = "8080"
	}

	http.ListenAndServe(":"+port, r)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!!!"))
}

func getData(url string) (responseData []byte) {
	response, err := http.Get(url)
	log.Println(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData
}

func handleNewsCategory(w http.ResponseWriter, r *http.Request) {
	url := "https://ancap.su/api/Article/ByCategory?token=&categ="

	vars := mux.Vars(r)
	cat := vars["cat"]
	url += cat

	responseData := getData(url)
	respond(w, responseData)
}

func handleCategory(w http.ResponseWriter, r *http.Request) {
	url := "https://ancap.su/api/Video/ByCategory?token=&categ="

	vars := mux.Vars(r)
	cat := vars["cat"]
	url += cat

	responseData := getData(url)
	respond(w, responseData)
}

func handleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	responseData := getData("https://ancap.su/api/Article/Get?token=&id=" + id)
	respond(w, responseData)
}

func handleVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	responseData := getData("https://ancap.su/api/Video/Get?token=&id=" + id)
	respond(w, responseData)
}

func handleNewsHomepage(w http.ResponseWriter, r *http.Request) {
	responseData := getData("https://ancap.su/api/News/HomePage")
	respond(w, responseData)
}

func respond(w http.ResponseWriter, responseData []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseData))
}
