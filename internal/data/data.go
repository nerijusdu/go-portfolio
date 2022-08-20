package data

var (
	Projects    []Project
	Blogs       []Blog
	Skills      []Skill
	Experiences []Experience
)

func init() {
	p, err := readProjects()
	if err != nil {
		panic(err)
	}

	b, err := readBlogs()
	if err != nil {
		panic(err)
	}

	s, err := readSkills()
	if err != nil {
		panic(err)
	}

	e, err := readExperiences()
	if err != nil {
		panic(err)
	}

	Projects = p
	Blogs = b
	Skills = s
	Experiences = e
}
