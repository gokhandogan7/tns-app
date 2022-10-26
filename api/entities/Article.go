package entities

// Article represents the model for an article
type Article struct {
	Id      int64  `json:"Id" example: "1"`
	Title   string `json:"Title" example: "Title 1"`
	Desc    string `json:"desc" example: "Description 1"`
	Context string `json:"context" example: "Content 1"`
}

var Articles = []Article{}
