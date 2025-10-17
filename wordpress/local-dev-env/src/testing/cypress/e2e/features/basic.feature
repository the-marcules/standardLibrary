Feature: Basic Setup Tests

  Scenario: Homepage has expected elements
    Given I open the homepage
    Then I see "WordPress"
    Then I see the top navigation menu
    Then I see the logo image
    Then I see the footer navigation menu

  Scenario: The Shop should be accessible and have one entry
    Given I open the homepage
    When I click on the "Shop" link in the top navigation menu
    Then I see "Shop" in the "main" content area
    Then I see the product "Testprodukt"

  Scenario: The Cart should be accessible and be empty
    Given I open the homepage
    When I click on the cart icon
    Then I see "Warenkorb" in the "main" content area
    Then I see "Dein Warenkorb ist leer" in the "main" content area
