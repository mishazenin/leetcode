package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a = 1
	fmt.Println(uintptr(unsafe.Pointer(&a)))
}

// LinkedList
type MyHashMap struct {
	List *MapNode
}
type MapNode struct {
	next  *MapNode
	value int
}

func ConstructorHashMap() MyHashMap {
	return MyHashMap{&MapNode{next: nil, value: 0}}
}

func (this *MyHashMap) Put(key int, value int) {

}

func (this *MyHashMap) Get(key int) int {
	return 0
}

func (this *MyHashMap) Remove(key int) {

}

// Array solution
// type MyHashMap struct {
// 	hashmap []*bucket
// }
//
// type bucket struct {
// 	val int
// }
//
// func Constructor() MyHashMap {
// 	return MyHashMap{hashmap: make([]*bucket, 1024)}
// }
//
// func (this *MyHashMap) Put(key int, value int) {
// 	if len(this.hashmap) < key {
// 		ext := make([]*bucket, key-len(this.hashmap)+2)
// 		this.hashmap = append(this.hashmap, ext...)
// 	}
// 	//as value use uinptr
// 	newBucket := &bucket{val: value}
// 	this.hashmap[key] = newBucket
// }
//
// func (this *MyHashMap) Get(key int) int {
// 	if key > len(this.hashmap) {
// 		return -1
// 	}
// 	res := this.hashmap[key]
// 	if res == nil {
// 		return -1
// 	}
// 	return res.val
// }
//
// func (this *MyHashMap) Remove(key int) {
// 	if key > len(this.hashmap) {
// 		return
// 	}
// 	this.hashmap[key] = nil
// }

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
