package articles

// Article article entity
type Article struct {
	ID    int
	Title string `json:"title"`
	Body  string `json:"body"`
}
