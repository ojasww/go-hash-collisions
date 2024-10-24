package main

import "fmt"

type Node struct {
	key   int
	value string
	next  *Node
}

// HashMap with chaining
type HashMap struct {
	buckets []*Node
	size    int
}

func NewHashMap(size int) *HashMap {
	return &HashMap{
		buckets: make([]*Node, size),
		size:    size,
	}
}

func (hm *HashMap) hashFunction(key int) int {
	return key % hm.size
}

func (hm *HashMap) Insert(key int, value string) {
	index := hm.hashFunction(key)
	head := hm.buckets[index]

	for head != nil {
		if head.key == key {
			head.value = value
			return
		}
		head = head.next
	}

	newNode := &Node{
		key:   key,
		value: value,
		next:  hm.buckets[index],
	}

	hm.buckets[index] = newNode
}

func (hm *HashMap) Get(key int) (string, bool) {
	index := hm.hashFunction(key)
	head := hm.buckets[index]

	for head != nil {
		if head.key == key {
			return head.value, true
		}

		head = head.next
	}

	return "", false
}

func (hm *HashMap) Delete(key int) bool {
	index := hm.hashFunction(key)
	head := hm.buckets[index]

	if head == nil {
		return false
	}

	if head.key == key {
		hm.buckets[index] = head.next
		return true
	}

	prev := head
	for head != nil {
		if head.key == key {
			prev.next = head.next
			return true
		}
		prev = head
		head = head.next
	}

	return false
}

func (hm *HashMap) ReHash() {
	newSize := hm.size * 2
	newBuckets := make([]*Node, newSize)

	for _, bucket := range hm.buckets {
		node := bucket
		for node != nil {
			newIndex := node.key % newSize

			newNode := &Node{key: node.key, value: node.value, next: newBuckets[newIndex]}
			newBuckets[newIndex] = newNode
			node = node.next
		}
	}

	hm.buckets = newBuckets
	hm.size = newSize
}

// PrintHashMap prints the contents of the hash map
func (hm *HashMap) PrintHashMap() {
	for i, node := range hm.buckets {
		fmt.Printf("Bucket %d: ", i)
		for node != nil {
			fmt.Printf("(%d, %s) -> ", node.key, node.value)
			node = node.next
		}
		fmt.Println("nil")
	}
}

func main() {
	hm := NewHashMap(10)

	hm.Insert(10, "A")
	hm.Insert(20, "B")
	hm.Insert(30, "C")
	hm.Insert(11, "D")

	hm.PrintHashMap()

	value, found := hm.Get(20)
	if found {
		fmt.Println("Found Value: ", value)
	} else {
		fmt.Println("Value not found")
	}

	isDeleted := hm.Delete(20)
	if !isDeleted {
		fmt.Println("Error deleting Value")
	} else {
		fmt.Println("Key deleted.")
	}

	hm.PrintHashMap()

	hm.ReHash()

	hm.PrintHashMap()
}

/**
Load factor: Number of elements in the table / Size of the table ( number of buckets )
*/
