package main

import (
	"container/list"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func sliceQueue() {
	slice := []int{}

	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)

	n := 4

	for len(slice) > 0 {
		elem := slice[0]
		fmt.Printf("elem: %v\n", elem)

		if rand.Intn(2) == 0 {
			fmt.Printf("appending %v\n", n)
			slice = append(slice, n)
			n++
		}

		slice = slice[1:]
	}
}

func listQueue() {
	queue := list.New()
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)

	n := 4

	for queue.Len() > 0 {
		elem := queue.Remove(queue.Front())
		fmt.Printf("elem: %v\n", elem)

		if rand.Intn(2) == 0 {
			fmt.Printf("appending %v\n", n)
			queue.PushBack(n)
			n++
		}
	}
}

func run(s string, f func(), n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Start of %s-based queue (%d)\n", s, i)
		f()
		fmt.Printf("End of %s-based queue (%d)\n", s, i)
	}
}

func main() {
	var n int
	rand.Seed(time.Now().UnixNano())

	flag.IntVar(&n, "n", 10, "amount of iterations")
	flag.Parse()

	if len(flag.Args()) == 0 {
		run("slice", sliceQueue, n)
		run("list", sliceQueue, n)
		return
	}

	for _, arg := range flag.Args() {
		switch arg {
		case "slice":
			run(arg, sliceQueue, n)
		case "list":
			run(arg, listQueue, n)
		}
	}
}
