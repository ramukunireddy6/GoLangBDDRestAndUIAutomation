package main

import (
	"github.com/tebeka/selenium"
)

func GetURL(url string) error {
	website := wd.Get(baseUrl + url)
	return website
}

func HTTPRequest(url string) error {
	request := wd.Get(baseUrl + url)
	return request
}

// Single element
// ===================================================================================================================
func FindElementByCss(locator string) selenium.WebElement {
	element, _ := wd.FindElement(selenium.ByCSSSelector, locator)
	return element
}

func FindElementById(locator string) selenium.WebElement {
	element, _ := wd.FindElement(selenium.ByID, locator)
	return element
}

func FindElementByXpath(locator string) selenium.WebElement {
	element, _ := wd.FindElement(selenium.ByXPATH, locator)
	return element
}

func FindElementByLinkText(locator string) selenium.WebElement {
	element, _ := wd.FindElement(selenium.ByLinkText, locator)
	return element
}

func FindElementByPartialLink(locator string) selenium.WebElement {
	element, _ := wd.FindElement(selenium.ByPartialLinkText, locator)
	return element
}

func FindElementByName(locator string) selenium.WebElement {
	element, _ := wd.FindElement(selenium.ByName, locator)
	return element
}

func FindElementByTag(locator string) selenium.WebElement {
	element, _ := wd.FindElement(selenium.ByTagName, locator)
	return element
}

func FindElementByClass(locator string) selenium.WebElement {
	element, _ := wd.FindElement(selenium.ByClassName, locator)
	return element
}

func FindElementByText(locator string) selenium.WebElement {
	element, _ := wd.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	return element
}

// Multi element
// ===================================================================================================================
func FindElementsById(locator string) []selenium.WebElement {
	element, _ := wd.FindElements(selenium.ByID, locator)
	return element
}

func FindElementsByXpath(locator string) []selenium.WebElement {
	element, _ := wd.FindElements(selenium.ByXPATH, locator)
	return element
}

func FindElementsByText(locator string) []selenium.WebElement {
	element, _ := wd.FindElements(selenium.ByXPATH, "//*[contains(text(), "+locator+")]")
	return element
}

func FindElementsByLinkText(locator string) []selenium.WebElement {
	element, _ := wd.FindElements(selenium.ByLinkText, locator)
	return element
}

func FindElementsByPartialLink(locator string) []selenium.WebElement {
	element, _ := wd.FindElements(selenium.ByPartialLinkText, locator)
	return element
}

func FindElementsByName(locator string) []selenium.WebElement {
	element, _ := wd.FindElements(selenium.ByName, locator)
	return element
}

func FindElementsByTag(locator string) []selenium.WebElement {
	element, _ := wd.FindElements(selenium.ByTagName, locator)
	return element
}

func FindElementsByClass(locator string) []selenium.WebElement {
	element, _ := wd.FindElements(selenium.ByClassName, locator)
	return element
}

func FindElementsByCss(locator string) []selenium.WebElement {
	element, _ := wd.FindElements(selenium.ByCSSSelector, locator)
	return element
}

func MouseHoverToElement(locator string) selenium.WebElement {
	element, _ := wd.FindElement(selenium.ByCSSSelector, locator)
	element.MoveTo(0, 0)
	return element
}
