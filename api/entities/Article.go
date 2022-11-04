package entities

// Article represents the model for an article
type FullArticle struct {
	Id         int64  `json:"Id"`
	Title      string `json:"Title"`
	Desc       string `json:"desc"`
	Content_Id int64  `json:"Content_Id"`
	Author_Id  int64  `json:"Author_Id"`
	Name       string `json:"Name"`
	Email      string `json:"Email"`
	Text       string `json:"Text"`
	Image      string `json:"Image"`
	Short_Text string `json:"Short_Text"`
	Date       string `json:"Date"`
}

type Author struct {
	Id    int64  `json:"Id"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type Content struct {
	Id    int64  `json:"Id"`
	Text  string `json:"Text"`
	Image string `json:"Image"`
}

type Highlight struct {
	Id         int64  `json:"Id"`
	Article_Id int64  `json:"Article_Id"`
	Short_Text string `json:"Short_Text"`
	Date       string `json:"Date"`
}

var Authors = []Author{}
var Contents = []Content{}
var Highlights = []Highlight{}
var FullArticles = []FullArticle{}
