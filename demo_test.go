package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DATA-DOG/godog"
	. "github.com/logrusorgru/aurora"
)

var url string

func websiteURL() error {
	url = os.Getenv("URL_WEBSITE")

	return nil
}

func accessWebsiteURL() error {
	seleniumWebDriverConnect()
	GetURL(url)

	return nil
}

func validateWebsiteURL() error {
	currentURL, _ := wd.CurrentURL()

	if expectURL := (url); currentURL != expectURL {
		wd.Screenshot()
		log.Fatalln(Bold(Red("URL DOESN'T MATCH")))
	} else {
		fmt.Println(Bold(Magenta("SCENARIO IS SUCCESS")))
	}

	return nil
}

func GodogDemoSteps(s *godog.Suite) {
	s.Step(`^user has url$`, websiteURL)
	s.Step(`^user access the url$`, accessWebsiteURL)
	s.Step(`^user must directed to correct url$`, validateWebsiteURL)
}
