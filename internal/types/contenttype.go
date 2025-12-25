package types

type ContentType interface {
	WebsiteRootPath() string
	ObsidianRootPath() string
	CreateNew(file ObsidianFile) ContentType // create a new instance from a provided ObsidianFile
}

type BookReview struct {
	Author string
	Title  string
}

func (b BookReview) WebsiteRootPath() string {
	return "library/books"
}

func (b BookReview) ObsidianRootPath() string {
	return "Database/Index/Books"
}

func (b BookReview) CreateNew(file ObsidianFile) ContentType {
	return BookReview{
		Author: "Test",
		Title:  file.Name,
	}
}

type MovieReview struct{}

func (m MovieReview) WebsiteRootPath() string {
	return "library/movies"
}

func (m MovieReview) ObsidianRootPath() string {
	return "Database/Index/Movies" // temp, not in Obsidian yet
}

func (m MovieReview) CreateNew(file ObsidianFile) ContentType {
	return MovieReview{}
}
