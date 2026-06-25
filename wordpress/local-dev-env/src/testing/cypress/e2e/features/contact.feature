Feature: Contact and revoke order form

  Background:
    Given I open the homepage
    
#   Scenario: I fill out the contact form and submit it
#     Given I see "Kontakt" in the "main" content area
#     When I fill out the contact form with name "Max Mustermann", email "max.mustermann@example.com", and message "Hello, this is a test message"
#     And I submit the contact form
#     Then I see "Vielen Dank für Ihre Nachricht. Wir werden uns so schnell wie möglich bei Ihnen melden." in the "main" content area



    Scenario: I fill out the revoke order form and submit it
    Given I see "Widerruf" in the "footer" content area
    And I click on the "Widerruf" link in the footer menu
    Then I see "Widerruf" in the "main" content area
    When I fill out the "kunde" field with "Max Mustermann"
    And I fill out the "email" field with "max.mustermann@example.com"
    And I fill out the "orderid" field with "123456"
    And I fill out the "msg" field with "Hello, this is a test message for revoking my order"
  When I click on the "Widerruf bestätigen" button
    Then I see "Der Widerruf konnte nicht gesendet werden. Bitte versuche es später erneut oder sende mir eine Email an koelnerholzschmied@gmail.com" in the "main" content area