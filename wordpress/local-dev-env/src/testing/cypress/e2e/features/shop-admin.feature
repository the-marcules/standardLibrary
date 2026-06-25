Feature: Shop Admin - Add/Edit/Delete Products

  Background:
    Given I open the homepage
    And I log in to the wordpress admin area
    Given I am on the "wopro_product" admin page with title "Produkte"

  Scenario Outline: Add a product to the Shop using wordpress
    Then I see "Testprodukt" in the products list
    When I click on the "Beitrag hinzufügen" link within the element "#wpbody-content"
    Then I see "Beitrag hinzufügen" in the "#wpcontent" content area
    And I see no error or warning within the element "#wpcontent"
    When I create a new product with the title "<name>", price "<price>", stock "<stock>" and details "<details>"
    Then I see "Der Beitrag wurde veröffentlicht. Beitrag anzeigen" in the "#wpcontent" content area
    When I click on the "Beitrag anzeigen" link
    Then I see the product details for "<name>"
    And I see "<price>" in the "main" content area
    And I see "<stock>" in the "main" content area
    And I see "<details>" in the ".wopro-details" content area

    Examples:
      | name          | price | stock | details             |
      | Testprodukt 1 |  9,99 |    10 | Erstes Produkt|
      | Testprodukt 2 | 19,99 |     5 | Zweites Produkt|
      | Testprodukt 3 | 29,99 |     0 | Drittes Produkt|  

  Scenario: Delete all products from the Shop using wordpress
    When I select all products except "Testprodukt"
    And I select "In den Papierkorb verschieben" in the bulk actions dropdown
    Then I see "3 Beiträge wurden in den Papierkorb verschoben" in the "#wpcontent" content area
    And I see one product in the products list
