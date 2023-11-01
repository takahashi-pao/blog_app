package article_model

// 記事情報モデル
type Article_list_model struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	DateTime  string   `json:"date"`
	Tag       []string `json:"tag"`
	Thumbnail string   `json:"thumbnail"`
}
