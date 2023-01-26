package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/nilsnook/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	// t.Run("Number of posts", func(t *testing.T) {
	// 	fs := fstest.MapFS{
	// 		"hello world.md":  {Data: []byte("Title: hi")},
	// 		"hello-world2.md": {Data: []byte("Title: hola")},
	// 	}
	//
	// 	posts, err := blogposts.NewPostsFromFS(fs)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	//
	// 	got := len(posts)
	// 	want := len(fs)
	//
	// 	if got != want {
	// 		t.Errorf("got %d posts, wanted %d posts", got, want)
	// 	}
	// })
	//
	// t.Run("Posts title", func(t *testing.T) {
	// 	fs := fstest.MapFS{
	// 		"hello world.md":  {Data: []byte("Title: Post 1")},
	// 		"hello-world2.md": {Data: []byte("Title: Post 2")},
	// 	}
	//
	// 	posts, err := blogposts.NewPostsFromFS(fs)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	//
	// 	assertPost(t, posts[0], blogposts.Post{Title: "Post 1"})
	// })

	t.Run("Posts title with description", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1`
			secondBody = `Title: Post 2
Description: Description 2`
		)

		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		assertPost(t, posts[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
		})
	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
