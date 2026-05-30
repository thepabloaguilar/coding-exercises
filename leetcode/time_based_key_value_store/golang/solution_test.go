package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/time_based_key_value_store/golang"
)

type command struct {
	setArgs  *setArgs
	getArgs  *getArgs
	expected interface{}
}

type setArgs struct {
	key       string
	value     string
	timestamp int
}

type getArgs struct {
	key       string
	timestamp int
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
				{setArgs: &setArgs{"alice", "happy", 1}, expected: nil},
				{getArgs: &getArgs{"alice", 1}, expected: "happy"},
				{getArgs: &getArgs{"alice", 2}, expected: "happy"},
				{setArgs: &setArgs{"alice", "sad", 3}, expected: nil},
				{getArgs: &getArgs{"alice", 3}, expected: "sad"},
			},
		},
		{
			name: "Example 2",
			commands: []command{
				{setArgs: &setArgs{"key1", "value1", 10}, expected: nil},
				{getArgs: &getArgs{"key1", 1}, expected: ""},
				{getArgs: &getArgs{"key1", 10}, expected: "value1"},
				{getArgs: &getArgs{"key1", 11}, expected: "value1"},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			timeMap := golang.Constructor()

			for _, cmd := range tc.commands {
				switch {
				case cmd.setArgs != nil:
					timeMap.Set(cmd.setArgs.key, cmd.setArgs.value, cmd.setArgs.timestamp)
				case cmd.getArgs != nil:
					result := timeMap.Get(cmd.getArgs.key, cmd.getArgs.timestamp)
					require.Equal(t, cmd.expected, result)
				}
			}
		})
	}
}
