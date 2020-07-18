package main

import (
	"net/http"

	"github.com/DATA-DOG/godog"
	"github.com/go-cucumber/features/step_definitions"
)

var req *http.Response

func sendRequest(request string) error {
	req = step_definitions.SendRequest(request)

	return nil
}

func validateResponse(response int) error {
	step_definitions.ValidateResponse(response, req)

	return nil
}

func GodogAPIDemoSteps(s *godog.Suite) {
	s.Step(`^I send a GET request to "([^\"]*)"$`, sendRequest)
	s.Step(`^the response status should be "([^\"]*)"$`, validateResponse)
}
