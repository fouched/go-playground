package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand/v2"
	"strings"
)

func main() {

	//castingMapWithInterfaces()
	//pointerSample()
	//forSample()
	//rangeSample()
	//randomExample()
	//hashPassword("password")
	//sliceSplit()
	getExtension("/uploads/jobseeker/1/hfx3hT7ZjjVVzzJtvbjz5LFHv.pdf")
}

func getExtension(s string) {
	extIdx := strings.LastIndex(s, ".")
	fmt.Println(s[extIdx:])
}

func sliceSplit() {
	exploded := strings.Split("/admin/2/3/4", "/")
	fmt.Println(exploded[4])
}

func hashPassword(p string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(p), 12)
	fmt.Println(string(hashedPassword))
}

func forSample() {
	for i := 5; i >= 5*-1; i-- {
		fmt.Println(i)
	}
	//for i := 5 * -1; i <= 5; i++ {
	//	fmt.Println(i)
	//}
}

func rangeSample() {

	for i := range 10 {
		fmt.Println(i)
	}
}

func pointerSample() {
	post := NewPost()

	fmt.Println("---- normal ----")
	fmt.Println(post.Created)
	fmt.Println(post.Title)
	fmt.Println(post.Body)
	modifyPost(post)
	fmt.Println("---- modified ----")
	fmt.Println(post.Created)
	fmt.Println(post.Title)
	fmt.Println(post.Body)

	pPost := NewPointerPost()
	fmt.Println("---- pointer ----")
	fmt.Println(pPost.Created)
	fmt.Println(pPost.Title)
	fmt.Println(pPost.Body)
	modifyPointerPost(pPost)
	fmt.Println("---- modified ----")
	fmt.Println(pPost.Created)
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

func randomExample() {
	fmt.Println(rand.IntN(2))
	fmt.Println(rand.IntN(2))
	fmt.Println(rand.IntN(2))
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

}
