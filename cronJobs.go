package main

import (
	"gopkg.in/robfig/cron.v2"
)

func getAndSaveContent() {

	// fetch content and save it
	content := fetchContent()
	saveFetchedContent(content)
}

func runCronJob() {

	// run cron job every 5 minutes
	c := cron.New()
	c.AddFunc("@every 5m", getAndSaveContent)
	c.Start()
}
