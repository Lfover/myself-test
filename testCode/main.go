package main

import (
	"fmt"
	"math"
	"sync"
)

type Task func() error

type Pool struct {
	capacity int
	queue    chan Task
	mu       sync.Mutex
	wg       sync.WaitGroup
	closed   bool
}

func NewPool(capacity int, queuelen int) *Pool {
	p := &Pool{
		capacity: capacity,
		queue:    make(chan Task, queuelen),
		closed:   false,
	}
	p.wg.Add(capacity)
	for i := 0; i < capacity; i++ {
		defer p.wg.Done()

	}
	return p

}

type ListNode struct {
	Val  int
	Node *ListNode
}

func main1() {

	ch := make(chan int)
	wg := sync.WaitGroup{}
	s := []string{"a", "b", "c"}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 10; i < 16; i += 2 {
			fmt.Println(i)
			ch <- 1
			<-ch
		}
	}()
	go func() {
		defer wg.Done()
		for _, v := range s {
			<-ch
			fmt.Println(v)
			ch <- 1
		}
	}()

	wg.Wait()
	return
}

var container = 1

type MyStack []int

func (s *MyStack) Push(v int) {
	*s = append(*s, v)
}

func (s *MyStack) Pop() int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *MyStack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *MyStack) Size() int {
	return len(*s)
}

func (s *MyStack) Empty() bool {
	return s.Size() == 0
}

// ---------- 分界线 ----------

type MyQueue struct {
	stackIn  *MyStack
	stackOut *MyStack
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  &MyStack{},
		stackOut: &MyStack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.stackIn.Push(x)
}

func (this *MyQueue) Pop() int {
	this.fillStackOut()
	return this.stackOut.Pop()
}

func (this *MyQueue) Peek() int {
	this.fillStackOut()
	return this.stackOut.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.stackIn.Empty() && this.stackOut.Empty()
}

// fillStackOut 填充输出栈
func (this *MyQueue) fillStackOut() {
	if this.stackOut.Empty() {
		for !this.stackIn.Empty() {
			val := this.stackIn.Pop()
			this.stackOut.Push(val)
		}
	}
}
func main() {
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
	x := math.Sqrt(10)
	fmt.Println(x)
}
