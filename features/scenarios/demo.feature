@all-3
Feature: GO Language - Learn and Demo Godog

  @demo
  Scenario: Open website and search
    Given user has url
    When user access the url
    Then user must directed to correct url
