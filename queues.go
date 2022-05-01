package main

import (
	"container/list"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/eapache/queue"
)

func listQueue(r bool) {
	q := list.New()
	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(3)

	n := 4

	if r {
		for q.Len() > 0 {
			elem := q.Remove(q.Front())
			fmt.Printf("elem: %v\n", elem)

			if rand.Intn(2) == 0 {
				fmt.Printf("appending %v\n", n)
				q.PushBack(n)
				n++
			}
		}
	} else {
		for q.Len() > 0 {
			elem := q.Remove(q.Front())
			fmt.Printf("elem: %v\n", elem)
		}
	}
}

func sliceQueue(r bool) {
	q := []int{}
	q = append(q, 1)
	q = append(q, 2)
	q = append(q, 3)

	n := 4

	if r {
		for len(q) > 0 {
			elem := q[0]
			fmt.Printf("elem: %v\n", elem)

			if rand.Intn(2) == 0 {
				fmt.Printf("appending %v\n", n)
				q = append(q, n)
				n++
			}

			q = q[1:]
		}
	} else {
		for len(q) > 0 {
			elem := q[0]
			fmt.Printf("elem: %v\n", elem)
			q = q[1:]
		}
	}
}

func queueQueue(r bool) {
	q := queue.New()
	q.Add(1)
	q.Add(2)
	q.Add(3)

	n := 4

	if r {
		for q.Length() > 0 {
			elem := q.Remove()
			fmt.Printf("elem: %v\n", elem)

			if rand.Intn(2) == 0 {
				fmt.Printf("appending %v\n", n)
				q.Add(n)
				n++
			}
		}
	} else {
		for q.Length() > 0 {
			elem := q.Remove()
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
		run("list", sliceQueue, n, r)
		run("queue", queueQueue, n, r)
		run("slice", sliceQueue, n, r)
		return
	}

	for _, arg := range flag.Args() {
		switch arg {
		case "list":
			run(arg, listQueue, n, r)
		case "queue":
			run(arg, queueQueue, n, r)
		case "slice":
			run(arg, sliceQueue, n, r)
		default:
			log.Fatalf("bad argument %s\n", arg)
		}
	}
}
