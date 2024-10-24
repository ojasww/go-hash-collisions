// hashmap_test.go
package main

import (
	"testing"
)

func TestHashMap_InsertAndGet(t *testing.T) {
	hm := NewHashMap(10)

	// Test inserting and retrieving values
	hm.Insert(1, "foo")
	hm.Insert(2, "bar")
	hm.Insert(3, "baz")

	val, found := hm.Get(1)
	if !found || val != "foo" {
		t.Errorf("expected value 1 for key 'foo', got %s", val)
	}

	val, found = hm.Get(2)
	if !found || val != "bar" {
		t.Errorf("expected value 2 for key 'bar', got %s", val)
	}

	val, found = hm.Get(3)
	if !found || val != "baz" {
		t.Errorf("expected value 3 for key 'baz', got %s", val)
	}

	// Test retrieving non-existent key
	_, found = hm.Get(4)
	if found {
		t.Error("expected key 4 not to be found")
	}
}

func TestHashMap_HandleCollisions(t *testing.T) {
	hm := NewHashMap(5) // Small size to force collisions

	// These keys will likely collide in a small map
	hm.Insert(10, "cat")
	hm.Insert(20, "tac")

	val, found := hm.Get(10)
	if !found || val != "cat" {
		t.Errorf("expected value 10 for key 'cat', got %s", val)
	}

	val, found = hm.Get(20)
	if !found || val != "tac" {
		t.Errorf("expected value 20 for key 'tac', got %s", val)
	}
}

func TestHashMap_UpdateExistingKey(t *testing.T) {
	hm := NewHashMap(10)

	// Insert and then update the value for the same key
	hm.Insert(1, "orange")
	hm.Insert(1, "kiwi") // Update the value

	val, found := hm.Get(1)
	if !found || val != "kiwi" {
		t.Errorf("expected value 'kiwi' for key 1 after update, got %s", val)
	}
}

func TestHashMap_ChainMultipleKeys(t *testing.T) {
	hm := NewHashMap(1) // Size 1 to force all keys into the same bucket

	// Multiple keys will be chained in the same bucket
	hm.Insert(100, "apple")
	hm.Insert(200, "cherry")
	hm.Insert(300, "banana")

	val, found := hm.Get(100)
	if !found || val != "apple" {
		t.Errorf("expected value 100 for key 'apple', got %s", val)
	}

	val, found = hm.Get(200)
	if !found || val != "cherry" {
		t.Errorf("expected value 200 for key 'banana', got %s", val)
	}

	val, found = hm.Get(300)
	if !found || val != "banana" {
		t.Errorf("expected value 300 for key 'cherry', got %s", val)
	}
}

func Test_RehashMap(t *testing.T) {
	hm := NewHashMap(4)

	hm.Insert(1, "apples")
	hm.Insert(2, "oranges")
	hm.Insert(3, "bananas")

	if hm.size != 4 {
		t.Errorf("size of hash map is not as expected")
	}

	hm.ReHash()

	// Size should be doubled by now

	// Checking if existing keys are present
	value, exists := hm.Get(1)
	if !exists {
		t.Errorf("value of '1' does not exist in the map")
	} else {
		if value != "apples" {
			t.Errorf("value of '1' does not match")
		}
	}

	if hm.size != 8 {
		t.Errorf("rehash did not double the hash map.")
	}
}
