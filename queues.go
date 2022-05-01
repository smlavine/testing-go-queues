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

func listQueue(n int, r bool) {
	q := list.New()
	for i := 1; i <= n; i++ {
		q.PushBack(i)
	}

	if r {
		n++
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

func sliceQueue(n int, r bool) {
	q := []int{}
	for i := 1; i <= n; i++ {
		q = append(q, i)
	}

	if r {
		n++
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

func queueQueue(n int, r bool) {
	q := queue.New()
	for i := 1; i <= n; i++ {
		q.Add(i)
	}

	if r {
		n++
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

func run(s string, f func(int, bool), iterations int, n int, r bool) {
	for i := 1; i <= iterations; i++ {
		fmt.Printf("Start of %s-based queue (%d)\n", s, i)
		f(n, r)
		fmt.Printf("End of %s-based queue (%d)\n", s, i)
	}
}

func main() {
	var (
		iterations int
		n          int
		r          bool
	)
	rand.Seed(time.Now().UnixNano())

	flag.IntVar(&iterations, "i", 1, "amount of iterations")
	flag.IntVar(&n, "n", 10, "amount of initial elements in the queue")
	flag.BoolVar(&r, "r", false, "randomly insert elements during loop")
	flag.Parse()

	if len(flag.Args()) == 0 {
		run("list", sliceQueue, iterations, n, r)
		run("queue", queueQueue, iterations, n, r)
		run("slice", sliceQueue, iterations, n, r)
		return
	}

	for _, arg := range flag.Args() {
		switch arg {
		case "list":
			run(arg, listQueue, iterations, n, r)
		case "queue":
			run(arg, queueQueue, iterations, n, r)
		case "slice":
			run(arg, sliceQueue, iterations, n, r)
		default:
			log.Fatalf("bad argument %s\n", arg)
		}
	}
}
