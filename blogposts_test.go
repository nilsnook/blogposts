package blogposts_test

import (
	"testing"
	"testing/fstest"

	"github.com/nilsnook/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
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
}
