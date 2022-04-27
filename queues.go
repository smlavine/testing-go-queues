package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("testing slice-based queue now")

	for i := 0; i < 100000; i++ {
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

	fmt.Println("end of slice-based queue")

	fmt.Println("start of list-based queue")

	for i := 0; i < 100000; i++ {
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

	fmt.Println("end of list-based queue")
}
