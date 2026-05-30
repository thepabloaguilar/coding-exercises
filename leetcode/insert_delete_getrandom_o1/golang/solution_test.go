package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/insert_delete_getrandom_o1/golang"
)

type command struct {
	method   string
	args     []int
	expected interface{}
}

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		commands []command
	}{
		{
			name: "Example 1",
			commands: []command{
				{"Insert", []int{1}, true},
				{"Remove", []int{2}, false},
				{"Insert", []int{2}, true},
				{"GetRandom", []int{}, 2},
				{"Remove", []int{1}, true},
				{"Insert", []int{2}, false},
				{"GetRandom", []int{}, 2},
			},
		},
		{
			name: "Example 2",
			commands: []command{
				{"Insert", []int{1}, true},
				{"Remove", []int{1}, true},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			randomizedSet := golang.Constructor()

			for _, cmd := range tc.commands {
				switch cmd.method {
				case "Insert":
					result := randomizedSet.Insert(cmd.args[0])
					require.Equal(t, cmd.expected, result)
				case "Remove":
					result := randomizedSet.Remove(cmd.args[0])
					require.Equal(t, cmd.expected, result)
				case "GetRandom":
					result := randomizedSet.GetRandom()
					require.Equal(t, cmd.expected, result)
				}
			}
		})
	}
}
