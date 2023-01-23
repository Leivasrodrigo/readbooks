package entities

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
	Read   bool   `json:"read"`
	Pages  int    `json:"pages"`
}
