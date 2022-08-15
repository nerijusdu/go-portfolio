package data

var (
	Projects []Project
	Blogs    []Blog
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

	Projects = p
	Blogs = b
}
