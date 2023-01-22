package main

import "fmt"

func main() {
	myQueue := ConstructorQueue()
	myQueue.Push(1)              // queue is: [1]
	myQueue.Push(2)              // queue is: [1, 2] (leftmost is front of the queue)
	fmt.Println(myQueue.Peek())  // return 1
	fmt.Println(myQueue.Pop())   // return 1, queue is [2]
	fmt.Println(myQueue.Empty()) // return false
}

type MyQueue struct {
	queue []int
}

func ConstructorQueue() MyQueue {
	return MyQueue{queue: make([]int, 0, 0)}
}

func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyQueue) Pop() int {
	var res = this.queue[0]
	this.queue = this.queue[1:]
	return res
}

func (this *MyQueue) Peek() int {

	return this.queue[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
