package data

var Projects []Project

func init() {
	p, err := readProjects()
	if err != nil {
		panic(err)
	}

	Projects = p
}
