package entities

type Article struct {
	Id      int64  `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Context string `json:"context"`
}

var Articles = []Article{}
