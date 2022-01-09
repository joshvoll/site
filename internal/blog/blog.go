package blog

import (
	"fmt"
	"html/template"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"site/internal/front"
	"sort"
	"strings"
	"time"

	blackfriday "github.com/russross/blackfriday/v2"
)

// Post defenition
type Post struct {
	Title      string        `json:"title"`
	Link       string        `json:"link"`
	Summary    string        `json:"summary"`
	Body       string        `json:"-"`
	BodyHTML   template.HTML `json:"body"`
	Date       time.Time
	DateString string `json:"date"`
}

type Posts []Post

func (p Posts) Len() int      { return len(p) }
func (p Posts) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Posts) Less(i, j int) bool {
	iDate := p[i].Date
	jDate := p[j].Date
	return iDate.Unix() < jDate.Unix()
}

// LoadPosts for a given directory
func LoadPosts(path string) (Posts, error) {
	type postFM struct {
		Title string
		Date  string
	}
	var result Posts
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("prevent panic by handler failure accesing a path: %q: %v ", path, err)
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		content, err := ioutil.ReadAll(file)
		if err != nil {
			return nil
		}
		var fm postFM
		remainder, err := front.Unmarshal(content, &fm)
		if err != nil {
			return err
		}
		output := blackfriday.Run(remainder)
		const timeFormat = `2006-01-02`
		date, err := time.Parse(timeFormat, fm.Date)
		if err != nil {
			return nil
		}
		p := Post{
			Title:      fm.Title,
			Date:       date,
			DateString: fm.Date,
			Link:       strings.Split(path, ".")[0],
			Body:       string(remainder),
			BodyHTML:   template.HTML(output),
		}
		result = append(result, p)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("cannot find the blogs: %v ", err)
	}
	sort.Sort(sort.Reverse(result))
	return result, nil
}
