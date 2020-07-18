package main

import (
	"fmt"
	"os"

	"github.com/DATA-DOG/godog"
	. "github.com/logrusorgru/aurora"
)

var userName string

func setUserName(name string) error {
	seleniumWebDriverConnect()
	userName = name

	return nil
}

func accessURL() error {
	GetURL(fmt.Sprintf(os.Getenv("URL_LOCAL")+"%s", userName))
	fmt.Println(wd.CurrentURL())

	return nil
}

func validateWindowTitle(name string) error {
	currentTitle, _ := wd.Title()

	if expectTitle := (fmt.Sprintf("Welcome %s", name)); currentTitle != expectTitle {
		wd.Screenshot()
		fmt.Println(Bold(Red("NAME DOESN'T MATCH")))
	} else {
		fmt.Println(Bold(Magenta("SCENARIO IS SUCCESS")))
	}

	return nil
}

func thereAreGodogs(available int) error {
	Godogs = available

	return nil
}

func iEat(num int) error {
	if Godogs < num {
		return fmt.Errorf("you cannot eat %d godogs, there are %d available", num, Godogs)
	}
	Godogs -= num

	return nil
}

func thereShouldBeRemaining(remaining int) error {
	if Godogs != remaining {
		return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, Godogs)
	}

	return nil
}

func GodogExampleSteps(s *godog.Suite) {
	s.Step(`^there are (\d+) godogs$`, thereAreGodogs)
	s.Step(`^I eat (\d+)$`, iEat)
	s.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
	s.Step(`^user with name "([^\"]*)"$`, setUserName)
	s.Step(`^user access url with given name$`, accessURL)
	s.Step(`^user must get window title Welcome "([^\"]*)"$`, validateWindowTitle)
}
