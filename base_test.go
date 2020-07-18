package main

import (
	"github.com/DATA-DOG/godog"
)

func GodogMainTest(s *godog.Suite) {
	s.BeforeScenario(func(interface{}) {
		Godogs = 0 // clean the state before every scenario
	})

	s.AfterScenario(func(interface{}, error) {
		if wd != nil {
			wd.Quit()
		}
	})
}
