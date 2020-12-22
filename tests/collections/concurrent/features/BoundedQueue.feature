Feature: Bounded queue
  As a bounded queue, no items should be added if full
  and no items can be be removed if empty

  @BoundedQueue
  Scenario: Equal number of producers and consumers leaves the queue empty
    Given there are 10 producers
    And there are 10 consumers
    And bounded queue capacity is 3
    When when producers and consumers run
    Then there should be 0 items remaining