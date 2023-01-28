package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
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
	const (
		titleSeparator       = "Title: "
		descriptionSeparator = "Description: "
		tagsSeparator        = "Tags: "
	)

	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	// Title
	title := readMetaLine(titleSeparator)
	// Description
	description := readMetaLine(descriptionSeparator)
	// Tags
	tags := strings.Split(readMetaLine(tagsSeparator), ",")
	// Body
	body := readBody(scanner)

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	// Ignore a line, for separtor `---`
	scanner.Scan()

	// Body
	// Scan until there is nothing left to scan
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	// Trim the last newline character added by `Fprintln`
	// in the for loop and return the string
	return strings.TrimSuffix(buf.String(), "\n")
}
