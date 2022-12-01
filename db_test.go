package n2i

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestLoagTags(t *testing.T) {
	tags, err := loadTags("datatest")
	assert.NoError(t, err)
	assert.Equal(t, map[string][]int{
		"Contraception": []int{3},
		"Prevention":    []int{1, 3},
		"Protéction":    []int{1, 3},
		"VIH":           []int{1},
	}, tags)
}

func TestLoadPage(t *testing.T) {
	page, err := loadPage("datatest", 1)
	assert.NoError(t, err)
	assert.Equal(t, &Page{
		Title: "VIH",
		Body:  "# Introduction\n\nBonjour tout le monde j'ai le VIH.",
		PageTags: PageTags{
			Tags: []string{
				"VIH", "Protéction", "Prevention",
			},
		},
		Date: time.Date(2022, time.December, 1, 20, 29, 4, 0, time.UTC),
	}, page)
}

func TestStorePage(t *testing.T) {
	expected := &Page{
		Title: "VIH",
		Body:  "# Introduction\n\nBonjour tout le monde j'ai le VIH.",
		PageTags: PageTags{
			Tags: []string{
				"VIH", "Protéction", "Prevention",
			},
		},
		Date: time.Date(2022, time.December, 1, 20, 29, 4, 0, time.UTC),
	}
	dirDB := "/tmp/n2i/"
	_, err := initDB(dirDB)
	assert.NoError(t, err)
	defer os.RemoveAll(dirDB)

	assert.NoError(t, storePage(dirDB, 42, expected))

	received, err := loadPage(dirDB, 42)
	assert.NoError(t, err)
	assert.Equal(t, expected, received)
}
