package model

type RequestCreateArticleInfo struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	Author string `json:"author"`
}
