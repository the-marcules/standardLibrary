Feature: Voucher Add as admin, use as user

  Background:
    Given I open the homepage
    And I log in to the wordpress admin area
    Given I am on the Woody Products Settings page "wopro_admin_submenu_vouchers"

  Scenario: Add a voucher to the Shop using wordpress
    Then I see the input field with id "code" to be "empty"
    When I click on the "Generieren" button
    Then I see the input field with id "code" to be "not empty"
    When I set the valid until date to "01.01.9999" by clicking
    And I set the voucher type to "voucher-absolute"
    And I set the voucher value to "5"
    And I submit the voucher form
    Then I see the success message that the voucher was created
    And I see the new voucher in the voucher list
    When I delete the created voucher
    Then I see the success message that the voucher was deleted

  Scenario Outline: Add a voucher to the Shop using wordpress
    Then I see the input field with id "code" to be "empty"
    When I set the voucher code to "<voucher_code>"
    Then I see the input field with id "code" to be "not empty"
    When I set the valid until date to "<date_valid>" by clicking
    And I set the voucher type to "<voucher_type>"
    And I set the voucher value to "<voucher_value>"
    And I submit the voucher form
    Then I see the success message that the voucher was created
    And I see the new voucher with code "<voucher_code>" in the voucher list
    When I delete the voucher with code "<voucher_code>"
    Then I see the success message that the voucher was deleted

    Examples:
      | voucher_code | date_valid | voucher_type     | voucher_value |
      | TEST10       | tomorrow   | voucher-absolute |            10 |
      | TESTPERC20   | tomorrow   | voucher-relative |            20 |
      | SHIPPING     | tomorrow   | voucher-shipping |               |


  Scenario Outline: Add a shipping voucher and use it in the shop
    Then I see the input field with id "code" to be "empty"
    When I set the voucher code to "<voucher_code>"
    Then I see the input field with id "code" to be "not empty"
    When I set the valid until date to "tomorrow" by clicking
    And I set the voucher type to "<voucher_type>"
    And I set the voucher value to "<voucher_value>"
    And I submit the voucher form
    Then I see the success message that the voucher was created
    And I see the new voucher in the voucher list
    When I open the homepage
    And I click on the "Shop" link in the top navigation menu
    When I click on the "Testprodukt" button
    Then I see the product details for "Testprodukt"
    When I click on the "in den Warenkorb" button
    Then I see the element "#banner"
    Then I see "Produkt wurde zum Warenkorb hinzugefügt." in the "#banner" content area
    Then I see "1" in the cart icon
    And I click on the cart icon
    Then I see "Warenkorb" in the "main" content area
    And I see "Testprodukt" in the cart summary
    When I enter the voucher code "<voucher_code>" in the voucher input field
    And I click on the "Einlösen" button
    Then I see "Du hast den Gutschein <voucher_code> eingelöst" in the "main" content area
    Then I see "Versand888,00" in the "table.wopro-sum-table" content area
    Then I see "Gutschein<voucher_calculated_value>" in the "table.wopro-sum-table" content area
    Then I see "Gesamtsumme<sum>" in the "table.wopro-sum-table" content area
    When I click on the "Gutschein entfernen" button
    Then I see "Versand888,00" in the "table.wopro-sum-table" content area
    Then I do not see "Gutschein<voucher_calculated_value>" in the "table.wopro-sum-table" content area
    Then I see "Gesamtsumme913,99" in the "table.wopro-sum-table" content area

 Examples:
      | voucher_code | date_valid | voucher_type     | voucher_value | sum     | voucher_calculated_value |
      | TEST10       | tomorrow   | voucher-absolute |            10 | 903,99  | 10,00                     |
      | TESTPERC20   | tomorrow   | voucher-relative |            20 | 731,19  | 182,80                     |
      | SHIPPING     | tomorrow   | voucher-shipping |            888 | 25,99  | 888,00                     |
