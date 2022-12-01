package n2i

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
