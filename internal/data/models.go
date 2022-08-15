package data

import "html/template"

type Project struct {
	Slug                string `yaml:"slug"`
	Name                string `yaml:"name"`
	Description         string `yaml:"description"`
	LongDescription     string `yaml:"longDescription"`
	LongDescriptionHTML template.HTML
	LongDescriptionPath string   `yaml:"longDescriptionPath"`
	Tags                []string `yaml:"tags"`
	Url                 string   `yaml:"url"`
	ImageUrl            string   `yaml:"imageUrl"`
	RepositoryUrl       string   `yaml:"repositoryUrl"`
	Order               int      `yaml:"order"`
}

type Blog struct {
	Slug        string `yaml:"slug"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Date        string `yaml:"date"`
	ImageUrl    string `yaml:"imageUrl"`
	Path        string `yaml:"path"`
	HTML        template.HTML
}
