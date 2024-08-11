package main

import "fmt"

func main() {

	castingMapWithInterfaces()
}

func castingMapWithInterfaces() {

	// create a slice of posts
	posts := []Post{
		{
			ID:       1,
			AuthorID: 1,
			Title:    "Title 1",
			Body:     "Some text 1",
		},
		{
			ID:       2,
			AuthorID: 2,
			Title:    "Title 2",
			Body:     "Some text 2",
		},
	}

	// add it to a map
	data := make(map[string]interface{})
	data["Posts"] = posts

	// range through it casting to the interface type
	for _, post := range data["Posts"].([]Post) {
		fmt.Println(post)
	}
}
