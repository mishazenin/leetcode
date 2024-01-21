package main

import "fmt"

var input = []struct {
	id  int
	val string
}{
	{
		id:  5,
		val: "",
	},
	{
		id:  3,
		val: "ccccc",
	},
	{
		id:  1,
		val: "aaaaa",
	},
	{
		id:  2,
		val: "bbbbb",
	},
	{
		id:  5,
		val: "eeeee",
	},
	{
		id:  4,
		val: "ddddd",
	},
}

func main() {
	stream := Constructor(5)
	for _, t := range input {
		res := stream.Insert(t.id, t.val)
		fmt.Println(res)
	}
}

type OrderedStream struct {
	hm     map[int]string
	n      int
	posPtr int
	ans    []string
}

// func Constructor(n int) OrderedStream {
// 	return OrderedStream{
// 		hm:     make(map[int]string),
// 		n:      n,
// 		posPtr: 1,
// 	}
// }

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.hm[idKey] = value
	this.ans = []string{}
	for this.posPtr <= this.n && this.hm[this.posPtr] != "" {
		this.ans = append(this.ans, this.hm[this.posPtr])
		this.posPtr++
	}
	return this.ans
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
