package recursive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	SearchResult int
	SearchFor    int
	SearchList   []int
}

func TestChop(t *testing.T) {
	assert := assert.New(t)

	testsData := []TestStruct{
		{0, 1, []int{1, 5, 6}},
		{2, 19, []int{1, 5, 19, 150}},
		// {0, 0, []int{0, 6, 8, 150}},
		// {1, 6, []int{0, 6, 8, 150}},
		{2, 8, []int{0, 6, 8, 150}},
		{3, 150, []int{0, 6, 8, 150}},
		{0, -1, []int{-1, 6, 8, 150}},
		// {-1, 5, []int{15, 26, 99}},
		{-1, 3, []int{}},
		// {-1, 3, []int{1}},
		// {0, 1, []int{1}},
		// {1, 6, []int{0, 6, 8, 56, 150}},
	}

	for _, test := range testsData {

		assert.Equal(test.SearchResult, Chop(test.SearchFor, test.SearchList), "Error: Search for %d inside the list: %d ", test.SearchFor, test.SearchList)
	}
}
