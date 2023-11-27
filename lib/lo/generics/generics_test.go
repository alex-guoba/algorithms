package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	m := map[string]int{
		"hello": 3,
		"world": 4,
	}

	v := Values(m)
	r := []int{3, 4}
	assert.Equal(t, v, r, "values")
	//fmt.Println(v)
}
