package n2i

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

const usersFileName = "users.json"

var fileNameRegexp = regexp.MustCompile(`^([1-9]\d*).json$`)

type Page struct {
	Title string    `json:"title"`
	Body  string    `json:"body"`
	Date  time.Time `json:"date"`
	Tags  []string  `json:"tags"`
}

type database string

func newDB(dirDB string) (database, error) {
	dirDB = filepath.Clean(dirDB)
	if err := os.MkdirAll(dirDB, 0o775); err != nil {
		return "", fmt.Errorf("Init the db %q %w", dirDB, err)
	}
	return database(dirDB), nil
}

func (db database) loadUsers() (map[string][sha256.Size]byte, error) {
	fileName := filepath.Join(string(db), usersFileName)
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Read users credencials fail from %q: %w", fileName, err)
	}

	m := make(map[string][sha256.Size]byte)
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("Load users credencials fail from %q: %w", fileName, err)
	}

	return m, nil
}

// Return
//  1. A map indexed by tags and the values is the index of page
//  2. A map indexed by index, and the value is the title page.
//  3. Error if occure
func (db database) loadAllMetadata() (map[string][]int, map[int]string, error) {
	entrys, err := os.ReadDir(string(db))
	if err != nil {
		return nil, nil, fmt.Errorf("Load DB %q: %w", db, err)
	}

	tags := make(map[string][]int)
	titles := make(map[int]string)

	for _, entry := range entrys {
		name := entry.Name()
		if name == usersFileName {
			continue
		}
		if !fileNameRegexp.MatchString(name) {
			return nil, nil, fmt.Errorf("Load DB %q: file %q is not a regular page file name", db, name)
		}
		id, _ := strconv.Atoi(fileNameRegexp.ReplaceAllString(name, "$1"))

		name = filepath.Join(string(db), entry.Name())
		data, err := os.ReadFile(name)
		if err != nil {
			return nil, nil, fmt.Errorf("DB load %q: %w", name, err)
		}
		page := Page{}
		if err := json.Unmarshal(data, &page); err != nil {
			return nil, nil, fmt.Errorf("DB load %q: %w", name, err)
		}
		for _, tag := range page.Tags {
			tags[tag] = append(tags[tag], id)
		}
		titles[id] = page.Title
	}

	return tags, titles, nil
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
