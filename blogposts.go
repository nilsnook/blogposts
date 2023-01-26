package blogposts

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
}

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := getPost(filesystem, f.Name())
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

func getPost(filesystem fs.FS, fileName string) (Post, error) {
	// Opening the file in the given filesystem
	postFile, err := filesystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	// REQUIRED: derfer closing the file
	defer postFile.Close()

	// Reading file contents
	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	// postData, err := io.ReadAll(postFile)
	// if err != nil {
	// 	return Post{}, err
	// }
	//
	// post := Post{Title: string(postData[7:])}
	// return post, nil

	const (
		titleSeparator       = "Title: "
		descriptionSeparator = "Description: "
	)

	scanner := bufio.NewScanner(postFile)

	// // Title
	// scanner.Scan()
	// titleLine := scanner.Text()
	// // Description
	// scanner.Scan()
	// descriptionLine := scanner.Text()

	// readLine := func() string {
	// 	scanner.Scan()
	//  return scanner.Text()
	// }

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	// title := readLine()[len(titleSeparator):]
	// description := readLine()[len(descriptionSeparator):]

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)

	return Post{
		Title:       title,
		Description: description,
	}, nil
}
