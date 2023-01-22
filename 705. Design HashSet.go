package main

type MyHashSet struct {
	next  *MyHashSet
	value int
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	newNode := &MyHashSet{
		next:  nil,
		value: key,
	}
	if this == nil {
		this = newNode
	}
	currNode := this
	for currNode != nil {
		if currNode.next == nil {
			currNode.next = newNode
		}
		currNode = currNode.next
	}
}

func (this *MyHashSet) Remove(key int) {
	if this == nil {
		return
	}
	currNode := this
	for currNode != nil {
		currNode = currNode.next
	}
}

func (this *MyHashSet) Contains(key int) bool {
	return true
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
