package data

type Project struct {
	Id              string
	Name            string
	Description     string
	LongDescription string
	Tags            []string
	Url             string
	ImageUrl        string
	RepositoryUrl   string
	Order           int
}
