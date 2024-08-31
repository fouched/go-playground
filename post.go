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
	c, _ := time.Parse(time.DateOnly, "2024-03-25")
	return Post{
		ID:       1,
		AuthorID: 1,
		Title:    "Post 1",
		Body:     "Body 1",
		Created:  c,
	}
}

func NewPointerPost() *Post {
	c, _ := time.Parse(time.DateOnly, "2024-08-31")
	return &Post{
		ID:       1,
		AuthorID: 1,
		Title:    "Post 1",
		Body:     "Body 1",
		Created:  c,
	}
}
