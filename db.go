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

func initDB(dirDB string) (string, error) {
	dirDB = filepath.Clean(dirDB)
	if err := os.MkdirAll(dirDB, 0o775); err != nil {
		return "", fmt.Errorf("Init the db %q %w", dirDB, err)
	}
	return dirDB, nil
}

func loadTags(dirDB string) (map[string][]int, error) {
	entrys, err := os.ReadDir(dirDB)
	if err != nil {
		return nil, fmt.Errorf("Load DB %q: %w", dirDB, err)
	}

	m := make(map[string][]int)

	for _, entry := range entrys {
		name := entry.Name()
		if !fileNameRegexp.MatchString(name) {
			return nil, fmt.Errorf("Load DB %q: file %q is not a regular page file name", dirDB, name)
		}
		id, _ := strconv.Atoi(fileNameRegexp.ReplaceAllString(name, "$1"))

		name = filepath.Join(dirDB, entry.Name())
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

func loadPage(dirDB string, id int) (*Page, error) {
	data, err := os.ReadFile(pagePath(dirDB, id))
	if err != nil {
		return nil, fmt.Errorf("Load %d: %w", id, err)
	}

	page := new(Page)
	if err := json.Unmarshal(data, page); err != nil {
		return nil, fmt.Errorf("Load %d: %w", id, err)
	}

	return page, nil
}

func storePage(dirDB string, id int, page *Page) error {
	data, _ := json.Marshal(page)
	err := os.WriteFile(pagePath(dirDB, id), data, 0o664)
	if err != nil {
		return fmt.Errorf("Store %d fail: %w", id, err)
	}
	return nil
}

func pagePath(dirDB string, id int) string {
	return filepath.Join(dirDB, strconv.Itoa(id)+".json")
}
