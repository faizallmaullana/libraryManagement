package books

type Books struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Year      string `json:"year"`
	Quantity  int    `json:"quantity"`
	Image     []byte `json:"image"`
}
