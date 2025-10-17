Feature: Shop

  Background:
    Given I open the homepage
    And I click on the "Shop" link in the top navigation menu

  Scenario: I put an available product into the cart
    Given I see "Shop" in the "main" content area
    And I see the product "Testprodukt"
    When I click on the "Testprodukt" button
    Then I see the product details for "Testprodukt"
    When I click on the "in den Warenkorb" button
    Then I see the element "#banner"
    Then I see "Produkt wurde zum Warenkorb hinzugefügt." in the "#banner" content area
    Then I see "1" in the cart icon
    # Then I see that the element "#banner" is gone
    When I click on the cart icon
    Then I see "Warenkorb" in the "main" content area
    And I see "Testprodukt" in the cart summary
    And I see the element "#no-shipping"
    When I check the checkbox with id "#no-shipping"
    Then I see "0,00" in the shipping cost summary
    When I see the element "#agb-check"
    And I check the checkbox with id "#agb-check"
    Then I see the element "#paypal-button-container"
