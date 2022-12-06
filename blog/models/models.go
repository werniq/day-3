package models

type Post struct {
	Id       int        `json:"id"`
	Author   *User      `json:"user"`
	Likes    int        `json:"likes"`
	Comments []*Comment `json:"Comments"`
}

type Comment struct {
	Id     int   `json:"id"`
	Post   *Post `json:"post"`
	Author *User `json:"author"`
	// Complain ?
	Category []*PostCategory `json:"category"`
}

type User struct {
	userId       int        `json:"userId"`
	username     string     `json:"username"`
	userLikes    []*Comment `json:"userLikes"`
	userComments []*Comment `json:"comments"`
	userPosts    []*Post    `json:"post"`
}

type PostCategory struct {
	Tech       string
	Business   string
	Cooking    string
	Music      string
	Education  string
	Films      string
	Theater    string
	Sport      string
	Adventures string
	Tourism    string
}
