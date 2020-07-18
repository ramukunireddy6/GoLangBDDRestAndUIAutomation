package main

import (
	"fmt"
	"os"

	"github.com/DATA-DOG/godog"
	"github.com/go-cucumber/features/object_abstractions"
)

func accessBukaBikeLandingPage() error {
	seleniumWebDriverConnect()
	GetURL(os.Getenv("URL_BIKE"))
	fmt.Println(wd.CurrentURL())

	return nil
}

func suggestionForm() error {
	clickYesIWantButton()
	inputSuggestionForm()

	return nil
}

func invalidSuggestionForm() error {
	clickYesIWantButton()
	clickSuggestionButton()

	return nil
}

func clickSuggestionButton() error {
	FindElementByCss(object_abstractions.BtnSubmit).Click()

	return nil
}

func validateEmptyErrorMessage() error {
	FindElementByText("Nama harus diisi").IsDisplayed()
	FindElementByText("Umur harus diisi").IsDisplayed()
	FindElementByText("Pekerjaan harus diisi").IsDisplayed()
	FindElementByText("Lokasi yang diinginkan harus diisi").IsDisplayed()

	return nil
}

func clickYesIWantButton() error {
	FindElementByCss(object_abstractions.BtnSuggestion).Click()

	return nil
}

func inputSuggestionForm() error {
	FindElementByXpath(object_abstractions.FieldName).SendKeys(os.Getenv("NAME"))
	FindElementByXpath(object_abstractions.RadioAge).Click()
	FindElementByXpath(object_abstractions.OptionAge).Click()
	FindElementByXpath(object_abstractions.FieldOccupation).SendKeys(os.Getenv("OCCUPATION"))
	FindElementByXpath(object_abstractions.FieldLocation).SendKeys(os.Getenv("LOCATION"))
	clickSuggestionButton()

	return nil
}

func validateSuggestionInputSuccess(msg string) error {
	FindElementByText(os.Getenv("VALIDATION")).IsDisplayed()

	return nil
}

func BukaBikeLandingPageSteps(s *godog.Suite) {
	s.Step(`^user access bike url$`, accessBukaBikeLandingPage)
	s.Step(`^user fill suggestion form$`, suggestionForm)
	s.Step(`^user does not fill suggestion form$`, invalidSuggestionForm)
	s.Step(`^user must not see window "([^\"]*)"$`, validateSuggestionInputSuccess)
	s.Step(`^user must see empty error message$`, validateEmptyErrorMessage)
}
