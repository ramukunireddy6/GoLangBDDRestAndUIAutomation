@all-2
Feature: Bike - Landing Page

  @test-1
  Scenario: Fill bike location suggestion form valid message
    Given user access bike url
    When user fill suggestion form
    Then user must not see window "Usulkan Lokasi Sepeda Selanjutnya"

  @test-2
  Scenario: Fill bike location suggestion form invalid data
    Given user access bike url
    When user does not fill suggestion form
    Then user must see empty error message
