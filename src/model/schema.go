package model

type CacheArticleInfoSchema struct {
	ID     int    `json:"id" redis:"id"`
	Title  string `json:"title" redis:"title"`
	Text   string `json:"text" redis:"text"`
	Author string `json:"author" redis:"author"`
}
