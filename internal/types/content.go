package types

import "path/filepath"

type Content interface {
	PathInWebsite() string
	PathInObsidian() string
	IsSinglePage() bool                    // indicates if the content type a non-list file (e.g., now page)
	IndexPathInObsidian() string           // single pages don't have an index path
	NewFromFile(file ObsidianFile) Content // create a new instance from a provided ObsidianFile
	TypeName() string                      // returns the type name for printing
}

const ObsidianWebsiteContentDir = "Notes/matkv.dev"

// index file only needs to be created if IndexPathInObsidian is not empty
// The index file then needs to be renamed into _index.md for Hugo

// Standalone content types

var StandalonePages = []Content{
	NowPage{},
	HomePage{},
}

type NowPage struct {
	Content string
}

func (n NowPage) PathInWebsite() string {
	return "/now"
}

func (n NowPage) PathInObsidian() string {
	return filepath.Join(ObsidianWebsiteContentDir, "Now.md")
}

func (n NowPage) IsSinglePage() bool {
	return true
}

func (n NowPage) IndexPathInObsidian() string {
	return ""
}

func (n NowPage) NewFromFile(file ObsidianFile) Content {
	return NowPage{
		Content: file.Content,
	}
}

func (n NowPage) TypeName() string {
	return "NowPage"
}

type HomePage struct {
	Content string
}

func (h HomePage) PathInWebsite() string {
	return "/"
}

func (h HomePage) PathInObsidian() string {
	return filepath.Join(ObsidianWebsiteContentDir, "Index.md")
}

func (h HomePage) IsSinglePage() bool {
	return true
}

func (h HomePage) IndexPathInObsidian() string {
	return ""
}

func (h HomePage) NewFromFile(file ObsidianFile) Content {
	return HomePage{
		Content: file.Content,
	}
}

func (h HomePage) TypeName() string {
	return "HomePage"
}

// List content types

var ListPages = []Content{
	BookReview{},
	MovieReview{},
	Project{},
}

type BookReview struct {
	Author string
	Title  string
}

func (b BookReview) PathInWebsite() string {
	return "library/books"
}

func (b BookReview) PathInObsidian() string {
	return "Database/Index/Books"
}

func (b BookReview) IsSinglePage() bool {
	return false
}

func (b BookReview) IndexPathInObsidian() string {
	return filepath.Join(ObsidianWebsiteContentDir, "Library", "Books", "Index.md")
}

func (b BookReview) NewFromFile(file ObsidianFile) Content {
	return BookReview{
		Author: "Test",
		Title:  file.Name,
	}
}

func (b BookReview) TypeName() string {
	return "BookReview"
}

type MovieReview struct {
	Title string
}

func (m MovieReview) PathInWebsite() string {
	return "library/movies"
}

func (m MovieReview) PathInObsidian() string {
	return "Database/Index/Movies" // not in Obsidian yet
}

func (m MovieReview) IsSinglePage() bool {
	return false
}

func (m MovieReview) IndexPathInObsidian() string {
	return filepath.Join(ObsidianWebsiteContentDir, "Library", "Movies", "Index.md")
}

func (m MovieReview) NewFromFile(file ObsidianFile) Content {
	return MovieReview{
		Title: file.Name,
	}
}

func (m MovieReview) TypeName() string {
	return "MovieReview"
}

type Project struct {
	Title string
}

func (p Project) PathInWebsite() string {
	return "projects"
}

func (p Project) PathInObsidian() string {
	return filepath.Join(ObsidianWebsiteContentDir, "Projects")
}

func (p Project) IsSinglePage() bool {
	return false
}

func (p Project) IndexPathInObsidian() string {
	return filepath.Join(ObsidianWebsiteContentDir, "Projects", "Index.md")
}

func (p Project) NewFromFile(file ObsidianFile) Content {
	return Project{
		Title: file.Name,
	}
}

func (p Project) TypeName() string {
	return "Project"
}

type Log struct {
	Content string
}

type Garden struct {
	Content string
}

func (g Garden) PathInWebsite() string {
	return "garden"
}

func (g Garden) PathInObsidian() string {
	return filepath.Join(ObsidianWebsiteContentDir, "Garden")
}

func (g Garden) IsSinglePage() bool {
	return false
}

func (g Garden) IndexPathInObsidian() string {
	return filepath.Join(ObsidianWebsiteContentDir, "Garden", "Index.md")
}

// won't be used directlly
func (g Garden) NewFromFile(file ObsidianFile) Content {
	return Garden{
		Content: file.Content,
	}
}

func (g Garden) TypeName() string {
	return "Garden"
}

type Library struct {
	Content string
}

func (l Library) PathInWebsite() string {
	return "library"
}

func (l Library) PathInObsidian() string {
	return filepath.Join(ObsidianWebsiteContentDir, "Library")
}

func (l Library) IsSinglePage() bool {
	return false
}

func (l Library) IndexPathInObsidian() string {
	return filepath.Join(ObsidianWebsiteContentDir, "Library", "Index.md")
}

// won't be used directlly
func (l Library) NewFromFile(file ObsidianFile) Content {
	return Library{
		Content: file.Content,
	}
}

func (l Library) TypeName() string {
	return "Library"
}
