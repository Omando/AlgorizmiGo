Feature: Singly linked list
  @SLLEmpty
  Scenario Outline: New list is empty
    Given linked list implementation is "<implementation>"
    When I create a new list
    Then head and tail are nil and size is zero
    Examples:
    |implementation  |
    |SinglyLinkedList|
    |DoublyLinkedList|
    |CircularlyLinkedList|

  Scenario Outline: Appending
    Given linked list implementation is "<implementation>"
    When I append items
      |value|
      |1    |
      |2    |
      |3    |
    Then Head is 1
    And Tail is 3
    Examples:
      |implementation|
      |SinglyLinkedList|
      |DoublyLinkedList|
      |CircularlyLinkedList|

  Scenario Outline: Prepending
    Given linked list implementation is "<implementation>"
    When I prepend items
      |value|
      |1    |
      |2    |
      |3    |
    Then Head is 3
    And Tail is 1
    Examples:
      |implementation|
      |SinglyLinkedList|
      |DoublyLinkedList|
      |CircularlyLinkedList|

  Scenario Outline: Removing head
    Given linked list implementation is "<implementation>"
    And  I append items
      |value|
      |1    |
      |2    |
      |3    |
    When I remove 1
    Then Head is 2
    And Tail is 3
    Examples:
      |implementation|
      |SinglyLinkedList|
      |DoublyLinkedList|
      |CircularlyLinkedList|

  Scenario Outline: Removing tail
    Given linked list implementation is "<implementation>"
    And  I append items
      |value|
      |1    |
      |2    |
      |3    |
    When I remove 3
    Then Head is 1
    And Tail is 2
    Examples:
      |implementation|
      |SinglyLinkedList|
      |DoublyLinkedList|
      |CircularlyLinkedList|

  Scenario Outline: Removing inner item
    Given linked list implementation is "<implementation>"
    And  I append items
      |value|
      |1    |
      |2    |
      |3    |
    When I remove 2
    Then Head is 1
    And Tail is 3
    Examples:
      |implementation|
      |SinglyLinkedList|
      |DoublyLinkedList|
      |CircularlyLinkedList|

  Scenario Outline: Removing non-existing item
    Given linked list implementation is "<implementation>"
    And  I append items
      |value|
      |1    |
      |2    |
      |3    |
    When I remove 100
    Then item is not found
    Examples:
      |implementation|
      |SinglyLinkedList|
      |DoublyLinkedList|
      |CircularlyLinkedList|

  Scenario Outline: Finding an existing item
    Given linked list implementation is "<implementation>"
    And  I append items
      |value|
      |1    |
      |2    |
      |3    |
    When I search for 2
    Then item is found
    Examples:
      |implementation|
      |SinglyLinkedList|
      |DoublyLinkedList|
      |CircularlyLinkedList|

  Scenario Outline: Finding a non-existing item
    Given linked list implementation is "<implementation>"
    And  I append items
      |value|
      |1    |
      |2    |
      |3    |
    When I search for 200
    Then item is not found
    Examples:
      |implementation|
      |SinglyLinkedList|
      |DoublyLinkedList|
      |CircularlyLinkedList|