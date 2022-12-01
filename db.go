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
	id    int
	Title string    `json:"title"`
	Body  string    `json:"body"`
	Date  time.Time `json:"date"`
	PageTags
}

type PageTags struct {
	Tags []string `json:"tags"`
}

func loadTags(dirDB string) (map[string][]int, error) {
	dirDB = filepath.Clean(dirDB)
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
