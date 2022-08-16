package pages

import "github.com/nerijusdu/go-portfolio/internal/data"

type HomePageData struct {
	Projects []data.Project
	Blogs    []data.Blog
}

type PageWithData[T any] struct {
	Version     string
	Title       string
	Description string
	Data        T
}
