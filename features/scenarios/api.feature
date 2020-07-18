@all-1
Feature: GO Language - Learn and Demo API

  @demo-api
  Scenario: send GET to endpoint-1
    When I send a GET request to "ENV:ENDPOINT_1"
    Then the response status should be "404"

  @api-demo
  Scenario: send GET to endpoint-1
    When I send a GET request to "ENV:ENDPOINT_1"
    Then the response status should be "200"
