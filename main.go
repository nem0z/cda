package main

import (
	"log"

	"github.com/ostafen/clover"
)

func main() {
	db, _ := clover.Open("storage")
	db.CreateCollection("myCollection")

	doc := clover.NewDocument()
	doc.Set("hello", "clover!")

	docId, _ := db.InsertOne("myCollection", doc)

	doc, _ = db.Query("myCollection").FindById(docId)
	log.Println(doc.Get("hello"))

}
