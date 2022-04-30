package main

import (
	"container/list"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func sliceQueue(r bool) {
	slice := []int{}

	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)

	n := 4

	if r {
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
	} else {
		for len(slice) > 0 {
			elem := slice[0]
			fmt.Printf("elem: %v\n", elem)
			slice = slice[1:]
		}
	}
}

func listQueue(r bool) {
	queue := list.New()
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)

	n := 4

	if r {
		for queue.Len() > 0 {
			elem := queue.Remove(queue.Front())
			fmt.Printf("elem: %v\n", elem)

			if rand.Intn(2) == 0 {
				fmt.Printf("appending %v\n", n)
				queue.PushBack(n)
				n++
			}
		}
	} else {
		for queue.Len() > 0 {
			elem := queue.Remove(queue.Front())
			fmt.Printf("elem: %v\n", elem)
		}
	}
}

func run(s string, f func(bool), n int, r bool) {
	for i := 0; i < n; i++ {
		fmt.Printf("Start of %s-based queue (%d)\n", s, i)
		f(r)
		fmt.Printf("End of %s-based queue (%d)\n", s, i)
	}
}

func main() {
	var n int
	var r bool
	rand.Seed(time.Now().UnixNano())

	flag.IntVar(&n, "n", 1, "amount of iterations")
	flag.BoolVar(&r, "r", false, "randomly insert elements during loop")
	flag.Parse()

	if len(flag.Args()) == 0 {
		run("slice", sliceQueue, n, r)
		run("list", sliceQueue, n, r)
		return
	}

	for _, arg := range flag.Args() {
		switch arg {
		case "slice":
			run(arg, sliceQueue, n, r)
		case "list":
			run(arg, listQueue, n, r)
		default:
			log.Fatalf("bad argument %s\n", arg)
		}
	}
}
