package main

import (
	"fmt"
	"time"

	"github.com/nem0z/cda/scrapper"
)

func main() {
	start := 285086
	n := 1000

	timeStart := time.Now()

	posts := scrapper.Process(start, n)

	cpt := 0
	for _, ok := range posts {
		if ok != nil {
			cpt++
		}
	}

	fmt.Printf("Found %v posts in %v!\n", cpt, time.Since(timeStart))
}
