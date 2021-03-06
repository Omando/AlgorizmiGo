package sequential

import (
	"AlgorizmiGo/collections/sequential/LinkedList"
	"AlgorizmiGo/collections/sequential/LinkedList/cicularyLinkedList"
	"AlgorizmiGo/collections/sequential/LinkedList/doublyLinkedList"
	"AlgorizmiGo/collections/sequential/LinkedList/singlyLinkedList"
	"fmt"
	"github.com/cucumber/godog"
	"os"
	"strconv"
	"testing"
)

var list LinkedList.LinkedList // interface implemented by singlyLinkedList and doublyLinkedList
var isItemFound bool

/*  Given */
//s.Step(`^linked list implementation is impName$`, linkedListImplementation)
func linkedListImplementation(arg string) error {
	if arg == "singlyLinkedList" {
		list = singlyLinkedList.New()
	} else if arg == "doublyLinkedList" {
		list = doublyLinkedList.New()
	} else if arg == "CircularlyLinkedList" {
		list = CircularlyLinkedList.New()
	}

	return nil
}

func iAppendItems(items *godog.Table) error {
	for i := 1; i < len(items.Rows); i++ {
		value, err := strconv.Atoi(items.Rows[i].Cells[0].Value)
		if err != nil {
		}
		list.Append(value)
	}
	return nil
}

/* When */
func iCreateANewList() error {
	list = singlyLinkedList.New()
	return nil
}

func iPrependItems(items *godog.Table) error {
	for i := 1; i < len(items.Rows); i++ {
		value, err := strconv.Atoi(items.Rows[i].Cells[0].Value)
		if err != nil {
		}
		list.Prepend(value)
	}
	return nil
}

func iRemove(arg1 int) error {
	_, err := list.Remove(arg1)
	if err != nil {
		isItemFound = false
	}
	return nil
}

func iSearchFor(arg1 int) error {
	isItemFound = list.Contains(arg1)
	return nil
}

/* Then */
func headAndTailAreNilAndSizeIsZero() error {
	if actualValue, err := list.First(); err == nil {
		return fmt.Errorf("expected head to be nil, but actual is: %d", actualValue)
	}

	if actualValue, err := list.Last(); err == nil {
		return fmt.Errorf("expected tail to be nil, but actual is: %d", actualValue)
	}
	if list.Length() != 0 {
		return fmt.Errorf("expected length be 0, but actual is: %d", list.Length())
	}

	return nil
}

func headIs(expectedValue int) error {
	if actualValue, err := list.First(); err != nil && expectedValue != actualValue {
		return fmt.Errorf("expected head to be %d, but actual is: %d", expectedValue, actualValue)
	}
	return nil
}

func itemIsFound() error {
	if isItemFound != true {
		return fmt.Errorf("expected item to be found")
	}
	return nil
}

func itemIsNotFound() error {
	if isItemFound != false {
		return fmt.Errorf("expected item not to be found")
	}
	return nil
}

func tailIs(expectedValue int) error {
	if actualValue, err := list.Last(); err != nil && actualValue != expectedValue {
		return fmt.Errorf("expected tail to be: %d, but actual is: %d", expectedValue, actualValue)
	}
	return nil
}

func InitializeScenario(s *godog.ScenarioContext) {
	s.Step(`^linked list implementation is "([^"]*)"$`, linkedListImplementation)
	s.Step(`^head and tail are nil and size is zero$`, headAndTailAreNilAndSizeIsZero)
	s.Step(`^Head is (\d+)$`, headIs)
	s.Step(`^I append items$`, iAppendItems)
	s.Step(`^I create a new list$`, iCreateANewList)
	s.Step(`^I prepend items$`, iPrependItems)
	s.Step(`^I remove (\d+)$`, iRemove)
	s.Step(`^I search for (\d+)$`, iSearchFor)
	s.Step(`^item is found$`, itemIsFound)
	s.Step(`^item is not found$`, itemIsNotFound)
	s.Step(`^Tail is (\d+)$`, tailIs)
}

func Test_Runner(t *testing.T) {
	opts := godog.Options{
		Format: "progress",
		Paths:  []string{"features"},
	}

	status := godog.TestSuite{
		Name:                "",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	os.Exit(status)
}
