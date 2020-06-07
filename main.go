package main

import (
	"encoding/json"
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
	r.HandleFunc("/target-list", handleTarget)
	r.HandleFunc("/video/get/{id}", handleVideo)
	r.HandleFunc("/article/get/{id}", handleArticle)
	r.HandleFunc("/video/bycategory/{cat}", handleCategory)
	r.HandleFunc("/article/bycategory/{cat}", handleNewsCategory)
	r.HandleFunc("/news/homepage", handleNewsHomepage)

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("* env variable PORT does not exist, using 8080")
		fmt.Println("* using 8080")
		fmt.Println("* try: http://192.168.0.7:8080")
		fmt.Println("* or use 'ip iddr show' on shell")
		port = "8080"
	}

	http.ListenAndServe(":"+port, r)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!!!"))
}

func getData(url string) (responseData []byte) {
	response, err := http.Get(url)
	// log.Println(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(string(responseData))
	return responseData
}

func unmarshalTarget(b []byte) {
	var data interface{}
	json.Unmarshal(b, &data)
	mapData := data.(map[string]interface{})
	target := mapData["Targets"]
	log.Println(target)
}

func handleTarget(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	ini := vars["ini"][0]
	max := vars["max"][0]

	url := "https://ancap.su/api/Target/List?token=&ini=" + ini + "&max=" + max

	responseData := getData(url)
	// unmarshalTarget(responseData)

	respond(w, r, responseData)
}

func handleNewsCategory(w http.ResponseWriter, r *http.Request) {
	url := "https://ancap.su/api/Article/ByCategory?token=&categ="

	vars := mux.Vars(r)
	cat := vars["cat"]
	url += cat

	responseData := getData(url)
	respond(w, r, responseData)
}

func handleCategory(w http.ResponseWriter, r *http.Request) {
	url := "https://ancap.su/api/Video/ByCategory?token=&categ="

	vars := mux.Vars(r)
	cat := vars["cat"]
	url += cat

	responseData := getData(url)
	respond(w, r, responseData)
}

func handleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	responseData := getData("https://ancap.su/api/Article/Get?token=&id=" + id)
	respond(w, r, responseData)
}

func handleVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	responseData := getData("https://ancap.su/api/Video/Get?token=&id=" + id)
	respond(w, r, responseData)
}

func handleNewsHomepage(w http.ResponseWriter, r *http.Request) {
	responseData := getData("https://ancap.su/api/News/HomePage")
	respond(w, r, responseData)
}

func respond(w http.ResponseWriter, r *http.Request, responseData []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseData))
}
