package data

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"sort"

	"github.com/gomarkdown/markdown"
	"gopkg.in/yaml.v2"
)

func readProjects() ([]Project, error) {
	projects, err := readYaml[[]Project]("projects.yaml")
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(projects); i++ {
		if p := projects[i]; p.LongDescriptionPath != "" {
			html, err := mdPathToHtml(p.LongDescriptionPath)
			if err != nil {
				return nil, fmt.Errorf("Failed to parse MD for project: %s. %v", p.Slug, err)
			}

			projects[i].LongDescriptionHTML = html
		}
	}

	sort.SliceStable(projects, func(i, j int) bool {
		return projects[i].Priority > projects[j].Priority
	})

	return projects, nil
}

func readBlogs() ([]Blog, error) {
	blogs, err := readYaml[[]Blog]("blogs.yaml")
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(blogs); i++ {
		if b := blogs[i]; b.Path != "" {
			html, err := mdPathToHtml(b.Path)
			if err != nil {
				return nil, fmt.Errorf("Failed to parse MD for blog: %s. %v", b.Slug, err)
			}

			blogs[i].HTML = html
		}
	}

	return blogs, nil
}

func readSkills() ([]Skill, error) {
	return readYaml[[]Skill]("skills.yaml")
}

func readExperiences() ([]Experience, error) {
	return readYaml[[]Experience]("experience.yaml")
}

func readYaml[T any](path string) (T, error) {
	var result T
	data, err := ioutil.ReadFile("data/" + path)
	if err != nil {
		return result, err
	}

	err = yaml.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func mdPathToHtml(path string) (template.HTML, error) {
	content, err := ioutil.ReadFile("data/" + path)
	if err != nil {
		return "", err
	}

	content = markdown.NormalizeNewlines(content)
	output := markdown.ToHTML(content, nil, nil)

	return template.HTML(output), nil
}
