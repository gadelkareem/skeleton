package paginator_test

import (
	"testing"

	"backend/utils/paginator"
	"github.com/stretchr/testify/assert"
)

func makeModels(n int) []interface{} {
	ms := make([]interface{}, n)
	for i := 0; i < n; i++ {
		ms[i] = i
	}
	return ms
}

func TestPaginator_Slice(t *testing.T) {
	ms := makeModels(10)

	// page 1 limit 5
	p := paginator.NewPaginator(map[string]int{"after": 1, "size": 5}, "", "", "/items")
	p.Models = ms
	p.Slice()
	assert.Equal(t, 5, len(p.Models))
	assert.Equal(t, 0, p.Models[0].(int))
	assert.Equal(t, 4, p.Models[4].(int))

	// page 2 limit 5
	p = paginator.NewPaginator(map[string]int{"after": 2, "size": 5}, "", "", "/items")
	p.Models = ms
	p.Slice()
	assert.Equal(t, 5, len(p.Models))
	assert.Equal(t, 5, p.Models[0].(int))
	assert.Equal(t, 9, p.Models[4].(int))

	// page 4 limit 3 => last page should have 1 item
	p = paginator.NewPaginator(map[string]int{"after": 4, "size": 3}, "", "", "/items")
	p.Models = ms
	p.Slice()
	assert.Equal(t, 1, len(p.Models))
	assert.Equal(t, 9, p.Models[0].(int))

	// limit -1 should return all models
	p = paginator.NewPaginator(map[string]int{"after": 1, "size": -1}, "", "", "/items")
	p.Models = ms
	p.Slice()
	assert.Equal(t, 10, len(p.Models))
}
