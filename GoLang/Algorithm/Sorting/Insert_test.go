package sorting_test

import (
	sorting "Algorithm/Sorting"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestInsert(t *testing.T) {
	rand.NewSource(time.Now().UnixMicro())
	slice_len := rand.Intn(100)
	test_case := make([]int, slice_len)
	for i := 0; i < slice_len; i++ {
		test_case[i] = rand.Intn(300)
	}
	answer := make([]int, slice_len)
	copy(answer, test_case)
	sort.Slice(answer, func(i, j int) bool {
		return answer[i] < answer[j]
	})
	sorting.Insert_Sort(test_case)
	assert.Equal(t, test_case, answer)
}
