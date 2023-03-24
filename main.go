package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nem0z/cda/database"
	"github.com/nem0z/cda/scrapper"
)

func main() {
	start := 286089
	n := 100

	db, err := database.Init()
	if err != nil {
		log.Fatal(err)
	}

	timeStart := time.Now()

	fmt.Println("Start scrapping")
	posts := scrapper.Process(start, n)
	fmt.Println(posts)

	cpt := 0
	for _, post := range posts {
		if post != nil {
			cpt++
			db.Insert(post)
		}
	}

	fmt.Printf("Found %v posts in %v!\n", cpt, time.Since(timeStart))

	post, err := db.Find(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("First post: %+v\n", post)
}
