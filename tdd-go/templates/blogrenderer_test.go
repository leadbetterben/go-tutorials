package blogrenderer

import (
	"bytes"
	"io"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1>

<p>This is a description</p>

Tags: <ul><li>go</li><li>tdd</li></ul>`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
