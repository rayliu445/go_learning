package goroutine

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func New(k int) *Tree {
	values := make([]int, 10)
	for i := 0; i < 10; i++ {
		values[i] = (i + 1) * k
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})

	var t *Tree
	for _, v := range values {
		t = insert(t, v)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func Walk(t *Tree, ch chan int) {
	var walker func(t *Tree)
	walker = func(t *Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
	close(ch)
}

func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 != ok2 || v1 != v2 {
			return false
		}

		if !ok1 && !ok2 {
			return true
		}
	}
}

func PracticeTest() {
	ch := make(chan int)
	go Walk(New(1), ch)

	fmt.Println("Walk(New(1),ch) output:")
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Same(New(1),New(1)) output:", Same(New(1), New(1)))
	fmt.Println("Same(New(1),New(2)) output:", Same(New(1), New(2)))
}
