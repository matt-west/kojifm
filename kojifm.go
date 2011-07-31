package main

import (
	"http"
)

const apiKey = "be6887e6212620554f63bd3c8e121b9d"
const apiSecret = "e9d7669333bc7a9815946babfd87f974"

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/watch/", watchHandler)
	http.HandleFunc("/assets/", assetHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func watchHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "watch.html")
}

const assetsPath = len("/assets/")

func assetHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[assetsPath:])
}
