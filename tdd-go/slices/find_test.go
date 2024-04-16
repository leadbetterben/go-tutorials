package main

import (
	"strings"
	"testing"
)

type Person struct {
	Name string
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})

	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			Person{Name: "Kent Beck"},
			Person{Name: "Martin Fowler"},
			Person{Name: "Chris James"},
			Person{Name: "Ben Leadbetter"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Ben")
		})

		AssertTrue(t, found)
		AssertEqual(t, king, Person{Name: "Ben Leadbetter"})
	})
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}
