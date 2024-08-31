package main

import "time"

type Post struct {
	ID       int
	AuthorID int
	Created  time.Time
	Title    string
	Body     string
}

func NewPost() Post {
	return Post{
		ID:       1,
		AuthorID: 1,
		Title:    "Post 1",
		Body:     "Body 1",
	}
}

func NewPointerPost() *Post {
	return &Post{
		ID:       1,
		AuthorID: 1,
		Title:    "Post 1",
		Body:     "Body 1",
	}
}
