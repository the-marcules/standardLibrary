Feature: Navigation

Scenario Outline: Navigate to all sub pages
    Given I am on the landing page
    Given I see the Navigation bar
    When I click on the "<Link>" link
    Then i should see the "<Link>" page

    Examples:
      | Link   | 
      | Second |
      | About  |

