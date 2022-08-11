package data

import "html/template"

type Project struct {
	Id                  int    `yaml:"id"`
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
