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
	req.Header.Set("cookie", `loid=0000000000arra1r4s.2.1615163100000.Z0FBQUFBQmdSVzdjV3J1Q2MtX2tIdzMtZFJRNHUzaHFWRG1DQUlIZkg0SG9mbGdjZ3hXYURsTlpYcV83Q1JLVFdSSW1UYU5QclBZRjIteTZHMWVNX21sbFpMQzlKaVl0MDJtZzBVNlVkYzY4OFBlWXF3RjF5ZnE2c1R4bjc1S2RSaGFVV2l5VjFiSEE; edgebucket=ed1Ff4docfz3pGptNu; over18=1; G_ENABLED_IDPS=google; pc=tk; csv=2; g_state={"i_l":0}; __stripe_mid=b9d840ff-c7a5-4b05-b9e2-6295afdeddeca6373b; reddit_session=844064228764%2C2021-12-17T18%3A42%3A19%2C92cdabce95d9787746d3481de76b0c5ac7d7991a; USER=eyJwcmVmcyI6eyJ0b3BDb250ZW50VGltZXNEaXNtaXNzZWQiOjAsInJwYW5EdURpc21pc3NhbFRpbWUiOm51bGwsImNvbGxhcHNlZFRyYXlTZWN0aW9ucyI6eyJmYXZvcml0ZXMiOmZhbHNlLCJtdWx0aXMiOmZhbHNlLCJtb2RlcmF0aW5nIjpmYWxzZSwic3Vic2NyaXB0aW9ucyI6ZmFsc2UsInByb2ZpbGVzIjpmYWxzZX0sImxheW91dCI6ImNhcmQiLCJnbG9iYWxUaGVtZSI6Ik5JR0hUIiwibmlnaHRtb2RlIjp0cnVlLCJ0b3BDb250ZW50RGlzbWlzc2FsVGltZSI6bnVsbH19; _rdt_uuid=1648581078028.e460e829-c76b-44bb-a4ba-86022a7ad1ce; token_v2=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTA2NTM2MDAsInN1YiI6Ijg0NDA2NDIyODc2NC1heGtqUFZ3TE4xWGpDbVJCQlNkVm9YX2J3V1A5ZmciLCJsb2dnZWRJbiI6dHJ1ZSwic2NvcGVzIjpbIioiLCJlbWFpbCIsInBpaSJdfQ.CmcthkLZe4MGVFHwtf_JkGhnhBmmal-TzsuO9qlVWPc; recent_srs=t5_2fwo%2Ct5_2rc7j%2Ct5_2qlyu%2Ct5_2qh4a%2Ct5_37z6f%2Ct5_2qs0k%2Ct5_2qizd%2Ct5_2bf7r5%2Ct5_2sxhs%2Ct5_3mcbrc; datadome=.5YV6.g8sPJ2g0ZPp58WLX.953swy3D5-_MEDz6YvN8od~0bTXOgcBFhMxw.~kLfRg_xRJ3aGwty7qh7RAkTZ26bkYW.f85jwk1l4a0b5rQ-Tn-GPzmVaY4qF44boA5Y; session_tracker=qjpiccqarcjrnejeeo.0.1650569278059.Z0FBQUFBQmlZYkEtT3VJSWNiSHZNVTVQbmhKOWVaU3R2cTVSUFlxa3FyWUhXSjFnMWJHMGFDbjZkVVBEUjhmekp1SnFMN29PLXNDbHdTQ3VueWJ4ZHJKUzl1bzJIWHg1OFN1aF9aSE9DczhUOXRBVjRpVUpCLUNwMW1kM0ZiVUhSQURyaVppQ1NPcGE`)
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
