package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	//castingMapWithInterfaces()

	post := NewPost()

	fmt.Println("---- normal ----")
	fmt.Println(post.Title)
	fmt.Println(post.Body)
	modifyPost(post)
	fmt.Println("---- modified ----")
	fmt.Println(post.Title)
	fmt.Println(post.Body)

	pPost := NewPointerPost()
	fmt.Println("---- pointer ----")
	fmt.Println(pPost.Title)
	fmt.Println(pPost.Body)
	modifyPointerPost(pPost)
	fmt.Println("---- modified ----")
	fmt.Println(pPost.Title)
	fmt.Println(pPost.Body)

}

func modifyPost(post Post) {
	post.Title = "Modified " + post.Title
	post.Body = "Modified " + post.Body
}
func modifyPointerPost(post *Post) {
	post.Title = "Modified " + post.Title
	post.Body = "Modified " + post.Body
}

func castingMapWithInterfaces() {

	//// create a slice of posts
	//posts := []Post{
	//	{
	//		ID:       1,
	//		AuthorID: 1,
	//		Title:    "Title 1",
	//		Body:     "Some text 1",
	//	},
	//	{
	//		ID:       2,
	//		AuthorID: 2,
	//		Title:    "Title 2",
	//		Body:     "Some text 2",
	//	},
	//}
	//
	//// add it to a map
	//data := make(map[string]interface{})
	//data["Posts"] = posts
	//
	//// range through it casting to the interface type
	//for _, post := range data["Posts"].([]Post) {
	//	fmt.Println(post)
	//}

	fmt.Println(rand.IntN(2))
	fmt.Println(rand.IntN(2))
	fmt.Println(rand.IntN(2))
}
