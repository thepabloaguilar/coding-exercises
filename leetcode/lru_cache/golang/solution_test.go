package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/lru_cache/golang"
)

type command struct {
	method   string
	args     []int
	expected interface{}
}

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		capacity int
		commands []command
	}{
		{
			name:     "Example 1",
			capacity: 2,
			commands: []command{
				{"Put", []int{1, 1}, nil},
				{"Put", []int{2, 2}, nil},
				{"Get", []int{1}, 1},
				{"Put", []int{3, 3}, nil},
				{"Get", []int{2}, -1},
				{"Put", []int{4, 4}, nil},
				{"Get", []int{1}, -1},
				{"Get", []int{3}, 3},
				{"Get", []int{4}, 4},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lruCache := golang.Constructor(tc.capacity)

			for _, cmd := range tc.commands {
				switch cmd.method {
				case "Put":
					require.NotPanics(t, func() {
						lruCache.Put(cmd.args[0], cmd.args[1])
					})
				case "Get":
					result := lruCache.Get(cmd.args[0])
					require.Equal(t, cmd.expected, result)
				}
			}
		})
	}
}
