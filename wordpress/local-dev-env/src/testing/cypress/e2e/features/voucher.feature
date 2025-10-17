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
    | voucher_code | date_valid | voucher_type | voucher_value |
    | TEST10       | tomorrow   | voucher-absolute     | 10            |
    | TESTPERC20   | tomorrow   | voucher-relative      | 20            |
    | SHIPPING   | tomorrow   | voucher-shipping      |             |