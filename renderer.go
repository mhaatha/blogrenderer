package blogrenderer

import "io"

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(w io.Writer, content Post) error {
	return nil
}
