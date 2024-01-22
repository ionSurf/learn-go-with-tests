package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/ionSurf/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("Retrieves details from the files", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
	Description: Description 1
	Tags: tdd, go
	---
	Hello
	World`
			secondBody = `Title: Post 2
	Description: Description 2
	Tags: rust, borrow-checker
	---
	B
	L
	M`
		)
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, _ := blogposts.NewPostsFromFS(fs)
		got := posts[0]
		want := blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		}

		assertPosts(t, got, want)

	})
}

func assertPosts(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v posts, want %v posts", got, want)
	}
}
