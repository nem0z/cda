package types

type Headline struct {
	Status string `bson:"status"`
	Title  string `bson:"title"`
	Price  string `bson:"price"`
}

type Author struct {
	Name string `bson:"name"`
	Date string `bson:"date"`
}

type Post struct {
	Index     int       `bson:"index"`
	Head      *Headline `bson:"head"`
	Author    *Author   `bson:"author"`
	Content   string    `bson:"content"`
	Images    []string  `bson:"images"`
	Signature string    `bson:"signature"`
}
