package main

import "fmt"

type Node struct {
	Key   string
	Value string
	Next  *Node
}

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

func hashFunction(key string, size int) uint {
	return uint(len(key) % size)

}

func (hm *HashMap) Insert(key string, value string) {
	index := hashFunction(key, hm.size)
	node := &Node{
		Key:   key,
		Value: value,
	}
	if hm.buckets[index] == nil {
		hm.buckets[index] = node
	} else {
		current := hm.buckets[index]
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

func (hm *HashMap) Get(key string) string {
	index := hashFunction(key, hm.size)
	current := hm.buckets[index]
	for current != nil {
		if current.Key == key {
			return current.Value
		}
		current = current.Next
	}
	return ""
}

func (hm *HashMap) Delete(key string) {
	index := hashFunction(key, hm.size)
	current := hm.buckets[index]
	var prev *Node
	for current != nil {
		if current.Key == key {
			if prev == nil {
				hm.buckets[index] = current.Next
			} else {
				prev.Next = current.Next
			}
			return
		}
		prev = current
		current = current.Next
	}
}

func main() {
	hashMap := NewHashMap(10)
	hashMap.Insert("firstname", "amitesh")
	hashMap.Insert("surname", "srivastava")
	value1 := hashMap.Get("firstname")
	hashMap.Delete("surname")
	value2 := hashMap.Get("surname")
	fmt.Println(value1, value2)
}

// h(k) = floor (m * frac (k * c))
