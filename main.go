package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nem0z/cda/scrapper"
)

func worker(indexs <-chan int, result chan *scrapper.Post) {
	for i := range indexs {
		res, err := scrapper.Process(i)
		if err != nil {
			log.Fatal(err)
		}
		result <- res
	}
}

func main() {
	start := 286049
	n := 1000

	jobs := make(chan int, n)
	result := make(chan *scrapper.Post, n)

	s := time.Now()

	for i := 0; i < 10; i++ {
		go worker(jobs, result)
	}

	for i := start; i >= start-n; i-- {
		jobs <- i
	}

	ok := 0
	done := 0
	go func() {
		for {
			time.Sleep(time.Second / 1)
			fmt.Printf("%v/%v pages parsed out of %v\n", ok, done, n)
		}
	}()

	for i := 0; i < n; i++ {
		select {
		case post := <-result:
			done++
			if post.Content != "" {
				ok++
			}
		}
	}

	fmt.Printf("%v/%v pages parsed out of %v in %v\n", ok, done, n, time.Since(s))
}
