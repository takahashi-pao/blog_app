package article_model

type Article_list_model struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	DateTime  string   `json:"date"`
	Tag       []string `json:"tag"`
	Thumbnail string   `json:"thumbnail"`
}
