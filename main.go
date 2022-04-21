// this is the main function which we will do the most important things
// in the end of this file we run two of our server
// server one -- api -- port:3333
// server two -- web server -- port:8000

package main

import (
	"fmt"
	"log"
	"net/http"
)

func apiData(w http.ResponseWriter, r *http.Request) {

	// tell the client that the content is json
	w.Header().Set("Content-Type", "application/json")

	// enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get content from content.json file
	posts := getContent()

	// serve this content to the client
	fmt.Fprint(w, posts)
}

func handleRequests() {

	// return content when /api endpoint is hit
	http.HandleFunc("/api", apiData)
	log.Fatal(http.ListenAndServe("127.0.0.1:3333", nil))
}

func webServer() {

	// get all files from /web folder and serve them
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}

func main() {
	// run both api and cron jobs in background
	// and keep the webserver in the main screen
	fmt.Println(`
	API SERVER: http://127.0.0.1:3333 ✅
	WEB SERVER: http://127.0.0.1:8000 ✅
	Cron Jobs: ✅
	`)
	go handleRequests()
	go runCronJob()
	webServer()
}
