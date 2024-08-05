package main

import (
	"fmt"
	"sync"
)

type FooBar struct {
	n   int
	ch1 chan bool
	ch2 chan bool
	wg  sync.WaitGroup
}

func main() {
	n := NewFooBar(5)
	n.wg.Add(2)
	go n.Foo()
	go n.Bar()
	n.wg.Wait()

}

func NewFooBar(n int) *FooBar {
	return &FooBar{n: n,
		ch1: make(chan bool),
		ch2: make(chan bool)}
}

func (fb *FooBar) Foo() {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
		// fb.ch1 <- true
		// printFoo()
		// fb.ch2 <- true
		// instead of two channels can also be done with one
		printFoo()
		fb.ch1 <- true
		<-fb.ch1
	}
	fb.wg.Done()
}

func (fb *FooBar) Bar() {
	for i := 0; i < fb.n; i++ {
		// printBar() outputs "bar". Do not change or remove this line.
		// <-fb.ch1
		// printBar()
		// <-fb.ch2
		// instead of two channels can also be done with one
		<-fb.ch1
		printBar()
		fb.ch1 <- true
	}
	fb.wg.Done()

}

func printFoo() {
	fmt.Println("Foo")
}

func printBar() {
	fmt.Println("Bar")
}
