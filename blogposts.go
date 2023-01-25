package blogposts

import (
	"io"
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := getPost(filesystem, f)
		if err != nil {
			// TODO: needs clarification
			// Should we totally fail if one file fails?
			// Or just ignore it
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(filesystem fs.FS, f fs.DirEntry) (Post, error) {
	// Opening the file in the given filesystem
	postFile, err := filesystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	// REQUIRED: derfer closing the file
	defer postFile.Close()

	// Reading file contents
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}

	// Very trivial parsing,
	// just slicing the `Title: ` which is of length 7
	post := Post{Title: string(postData[7:])}
	return post, nil
}
