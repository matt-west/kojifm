package main

import (
	"net/http"
)

type User struct {
	Name string
	Playlist map[string]*Video
}

type Video struct {
	Name string
	URL string
}

const apiKey = "be6887e6212620554f63bd3c8e121b9d"
const apiSecret = "e9d7669333bc7a9815946babfd87f974"

const watchPath = len("/watch/")
const assetPath = len("/assets/")

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pages/index.html")
}

func watchHandler(w http.ResponseWriter, r *http.Request) {
	user := User{r.FormValue("u"), nil}

	http.ServeFile(w, r, "pages/watch.html")
}

func assetHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[assetPath:])
}

func main() {
	http.HandleFunc("/watch/", watchHandler)
	http.HandleFunc("/assets/", assetHandler)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":9980", nil)
}
