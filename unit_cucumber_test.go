package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/go-cucumber/features/step_definitions"
)

func UnitTestCucumber(s *godog.Suite) {
	s.Step(`^user has a name "([^\"]*)"$`, step_definitions.GivenUserName)
	s.Step(`^Testivus meet user$`, step_definitions.MeetUserName)
	s.Step(`^Testivus say "([^\"]*)"$`, step_definitions.SayHelloName)
}
