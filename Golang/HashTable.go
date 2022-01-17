package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

// Insert -> Pointer receiver
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search
func (h *HashTable) Search(key string) bool {
	return false
}

// Delete
func (h *HashTable) Delete(key string) {

}

// insert
func (b *bucket) insert(k string) {
	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

// search
func (b *bucket) search(k string) bool {
	curNode := b.head
	for curNode != nil {
		if curNode.key == k {
			return true
		}
		curNode = curNode.next
	}
	return false
}

// delete

func hash(key string) int {
	var sum int
	for _, v := range key {
		sum += int(v)
	}
	fmt.Println(sum)
	return sum % ArraySize
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	// HashMap, HashTable, HashSet
	hashTable := Init()
	fmt.Println(hashTable)

	testBucket := &bucket{}
	testBucket.insert("RANDY")
	fmt.Println("_", testBucket.head.key)

	fmt.Println(testBucket.search("RANDY"))
	fmt.Println(testBucket.search("ERIC"))
}
