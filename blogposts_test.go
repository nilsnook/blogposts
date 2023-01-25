package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/nilsnook/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("Number of posts", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("Title: hi")},
			"hello-world2.md": {Data: []byte("Title: hola")},
		}

		posts, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		got := len(posts)
		want := len(fs)

		if got != want {
			t.Errorf("got %d posts, wanted %d posts", got, want)
		}
	})

	t.Run("Posts title", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("Title: Post 1")},
			"hello-world2.md": {Data: []byte("Title: Post 2")},
		}

		posts, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		got := posts[0]
		want := blogposts.Post{Title: "Post 1"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})
}
