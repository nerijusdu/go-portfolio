package data

import "html/template"

type BaseItem struct {
	Slug        string `yaml:"slug"`
	Hidden      bool   `yaml:"hidden"`
	Highlighted bool   `yaml:"highlighted"`
	Order       int    `yaml:"order"`
}

func (b BaseItem) IsHidden() bool      { return b.Hidden }
func (b BaseItem) IsHighlighted() bool { return b.Highlighted }

type Project struct {
	BaseItem            `yaml:",inline"`
	Name                string `yaml:"name"`
	Description         string `yaml:"description"`
	LongDescription     string `yaml:"longDescription"`
	LongDescriptionHTML template.HTML
	LongDescriptionPath string   `yaml:"longDescriptionPath"`
	Tags                []string `yaml:"tags"`
	Url                 string   `yaml:"url"`
	ImageUrl            string   `yaml:"imageUrl"`
	RepositoryUrl       string   `yaml:"repositoryUrl"`
}

type Blog struct {
	BaseItem    `yaml:",inline"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Date        string `yaml:"date"`
	ImageUrl    string `yaml:"imageUrl"`
	Path        string `yaml:"path"`
	HTML        template.HTML
}

type Skill struct {
	Name  string `yaml:"name"`
	Level int    `yaml:"level"`
}

type Experience struct {
	Name        string   `yaml:"name"`
	Stack       []string `yaml:"stack"`
	Description string   `yaml:"description"`
	Current     bool     `yaml:"current"`
	First       bool     `yaml:"first"`
	Start       string   `yaml:"start"`
	End         string   `yaml:"end"`
}
