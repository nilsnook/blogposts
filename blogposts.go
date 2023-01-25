package blogposts

import (
	"io/fs"
)

type Post struct {
}

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}
	return posts, nil
}
