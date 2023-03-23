package types

type Headline struct {
	Status string `json:"status"`
	Title  string `json:"title"`
	Price  string `json:"price"`
}

type Author struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type Post struct {
	Index     int       `json:"index"`
	Head      *Headline `json:"head"`
	Author    *Author   `json:"author"`
	Content   string    `json:"content"`
	Images    []string  `json:"images"`
	Signature string    `json:"signature"`
}
