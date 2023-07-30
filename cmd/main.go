package main

import (
	"log"
	"os"

	"github.com/nilsnook/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", posts)
}
