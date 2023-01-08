package pages

import "github.com/nerijusdu/go-portfolio/internal/data"

type HomePageData struct {
	Projects    []data.Project
	Blogs       []data.Blog
	Skills      []data.Skill
	Experiences []data.Experience
}

type PageWithData[T any] struct {
	Version     string
	Title       string
	Description string
	ImageUrl    string
	Data        T
}
