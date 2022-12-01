package n2i

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

var fileNameRegexp = regexp.MustCompile(`^([1-9]\d*).json$`)

type Page struct {
	Title string    `json:"title"`
	Body  string    `json:"body"`
	Date  time.Time `json:"date"`
	PageTags
}

type PageTags struct {
	Tags []string `json:"tags"`
}

type database string

func newDB(dirDB string) (database, error) {
	dirDB = filepath.Clean(dirDB)
	if err := os.MkdirAll(dirDB, 0o775); err != nil {
		return "", fmt.Errorf("Init the db %q %w", dirDB, err)
	}
	return database(dirDB), nil
}

func (db database) loadTags() (map[string][]int, error) {
	entrys, err := os.ReadDir(string(db))
	if err != nil {
		return nil, fmt.Errorf("Load DB %q: %w", db, err)
	}

	m := make(map[string][]int)

	for _, entry := range entrys {
		name := entry.Name()
		if !fileNameRegexp.MatchString(name) {
			return nil, fmt.Errorf("Load DB %q: file %q is not a regular page file name", db, name)
		}
		id, _ := strconv.Atoi(fileNameRegexp.ReplaceAllString(name, "$1"))

		name = filepath.Join(string(db), entry.Name())
		data, err := os.ReadFile(name)
		if err != nil {
			return nil, fmt.Errorf("DB load %q: %w", name, err)
		}
		tags := PageTags{}
		if err := json.Unmarshal(data, &tags); err != nil {
			return nil, fmt.Errorf("DB load %q: %w", name, err)
		}
		for _, tag := range tags.Tags {
			m[tag] = append(m[tag], id)
		}
	}

	return m, nil
}

func (db database) loadPage(id int) (*Page, error) {
	data, err := os.ReadFile(db.pagePath(id))
	if err != nil {
		return nil, fmt.Errorf("Load %d: %w", id, err)
	}

	page := new(Page)
	if err := json.Unmarshal(data, page); err != nil {
		return nil, fmt.Errorf("Load %d: %w", id, err)
	}

	return page, nil
}

func (db database) storePage(id int, page *Page) error {
	data, _ := json.Marshal(page)
	err := os.WriteFile(db.pagePath(id), data, 0o664)
	if err != nil {
		return fmt.Errorf("Store %d fail: %w", id, err)
	}
	return nil
}

func (db database) pagePath(id int) string {
	return filepath.Join(string(db), strconv.Itoa(id)+".json")
}
