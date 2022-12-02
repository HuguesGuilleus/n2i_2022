package n2i

import (
	"crypto/sha256"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestLoadUsers(t *testing.T) {
	users, err := database("datatest").loadUsers()
	assert.NoError(t, err)

	assert.Equal(t, map[string][sha256.Size]byte{
		"user1": sha256.Sum256([]byte("123456")),
		"user2": sha256.Sum256([]byte("abcdef")),
	}, users)
}

func TestLoagTags(t *testing.T) {
	tags, err := database("datatest").loadTags()
	assert.NoError(t, err)
	assert.Equal(t, map[string][]int{
		"Contraception": []int{3},
		"Prevention":    []int{1, 3},
		"Protéction":    []int{1, 3},
		"VIH":           []int{1},
	}, tags)
}

func TestLoadPage(t *testing.T) {
	page, err := database("datatest").loadPage(1)
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

	db, err := newDB("/tmp/n2i/")
	assert.NoError(t, err)
	defer os.RemoveAll(string(db))

	assert.NoError(t, db.storePage(42, expected))

	received, err := db.loadPage(42)
	assert.NoError(t, err)
	assert.Equal(t, expected, received)
}
