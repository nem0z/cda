package main

import (
	"fmt"
	"time"

	"github.com/nem0z/cda/database"
	"github.com/nem0z/cda/scrapper"
	"github.com/nem0z/cda/utils"
)

func main() {
	start := 286089
	n := 25

	db, err := database.Init()
	utils.Handle(err)

	err = db.Clear()
	utils.Handle(err)

	timeStart := time.Now()

	posts := scrapper.Process(start, n)

	cpt := 0
	for _, post := range posts {
		if post != nil {
			cpt++
			db.InsertOne(post)
		}
	}

	fmt.Printf("Found %v/%v posts in %v!\n", cpt, n, time.Since(timeStart))

	post, err := db.FindAll()
	utils.Handle(err)

	fmt.Printf("First post: %+v\n", post)
}
