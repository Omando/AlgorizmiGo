Feature: Singly linked list

  @SLLEmpty
  Scenario: New list is empty
    When I create a new list
    Then head and tail are nil and size is zero

  Scenario:Appending
    Given A new linked list
    When I append items
    |value|
    |1    |
    |2    |
    |3    |
    Then Head is 1
    And Tail is 3

  Scenario: Prepending
    Given A new linked list
    When I prepend items
      |value|
      |1    |
      |2    |
      |3    |
    Then Head is 3
    And Tail is 1

  Scenario: Removing head
    Given A new linked list
    And  I append items
      |value|
      |1    |
      |2    |
      |3    |
    When I remove 1
    Then Head is 2
    And Tail is 3

  Scenario: Removing tail
    Given A new linked list
    And  I append items
      |value|
      |1    |
      |2    |
      |3    |
    When I remove 3
    Then Head is 1
    And Tail is 2

  Scenario: Removing inner item
    Given A new linked list
    And  I append items
      |value|
      |1    |
      |2    |
      |3    |
    When I remove 2
    Then Head is 1
    And Tail is 3

  Scenario: Finding an existing item
    Given A new linked list
    And  I append items
      |value|
      |1    |
      |2    |
      |3    |
    When I search for 2
    Then item is found

  Scenario: Finding a non-existing item
    Given A new linked list
    And  I append items
      |value|
      |1    |
      |2    |
      |3    |
    When I search for 200
    Then item is not found
