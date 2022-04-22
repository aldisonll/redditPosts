// just a simple package to get posts from content.json file
// and to get the live data which will be saved to content.json

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func fetchContent() []byte {

	// get http.Client pointer
	client := &http.Client{}

	// create a new request
	req, err := http.NewRequest("GET", "https://www.reddit.com/r/programming/top.json?limit=100&t=month", nil)
	if err != nil {
		log.Fatal(err)
	}

	// set headers for this request
	req.Header.Set("authority", "www.reddit.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	// here. since i'm using a private api, not their api. I need to set my cookies here in order to work
	req.Header.Set("cookie", `loid=0000000000m9gzn1hz.2.1650642814366.Z0FBQUFBQmlZczlfcm9aZUdyN3hZQVNzSENoeVZnSkI1NHdXTnlUZlhPZU1IWHdmLTBoQlBFdE85aFpibmkzOFpITHhEWjQyREFHUFI5MGx2clJTU3hnbncxTnFiVWhIN0ZuRlBjMGpTbGFqTFFiRkFwN1ZJTzJVSE5HcjhDX1FsSzJlSFVMeC1IZ0Q; session_tracker=Z7ERRC6y6vrtOoYbU7.0.1650642814367.Z0FBQUFBQmlZczlfZlRIVGpYd2U5d0pwT3NueTV4TU00MFJ0TjN6bmZFeWMzSEFwWFZXa3A5dFkwa1RSRnRWOUZqcnAxVWcxQlJzSGljd3NYblpUZGtBd2lMU0JuWU1kLU9UUEFBdVhSMFNNNkRNQW90UGZkNTRsTE5IdkYyMG45aTloRDZDdG9yVjk; csv=2; edgebucket=0jabwqTcYOQukrAxXV`)
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")

	// make the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// get the body from response request
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return bodyText
}

func saveFetchedContent(content []byte) {

	// open content.json file as read/write only
	f, err := os.OpenFile("content.json", os.O_WRONLY, 0644)

	// check if something's wrong
	if err != nil {
		panic(err)
	}

	// close the opened file at the end of execution
	defer f.Close()

	// write content to this file
	f.Write([]byte(content))
}

func getContent() string {

	// read content.txt which contains all json content
	data, err := os.ReadFile("content.json")

	// check if something's wrong
	if err != nil {
		panic(err)
	}

	// convert bytes to string and return it
	return string(data)
}
