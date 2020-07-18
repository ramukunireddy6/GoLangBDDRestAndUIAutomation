@all-4
Feature: GO Language - Try Godog

  @test-a
  Scenario: Eat 5 out of 12
    Given there are 12 godogs
    When I eat 5
    Then there should be 7 remaining

  @test-b
  Scenario: Open localhost valid welcome message
    Given user with name "mpermperpisang"
    When user access url with given name
    Then user must get window title Welcome "mpermperpisang"

  @test-c
  Scenario: Open localhost invalid welcome message
    Given user with name "mpermperpisang"
    When user access url with given name
    Then user must get window title Welcome "ferawati"
