package en //nolint:testpackage // package should be `en_test` instead of `en`

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kuoss/namegen/util"
)

func TestLeft(t *testing.T) {
	t.Parallel()

	wantMinLen := 3
	wantMaxLen := 13

	minLen := util.Len(left[0])
	maxLen := util.Len(left[0])

	for _, word := range left {
		l := util.Len(word)

		if l < minLen {
			minLen = l
		}

		if l > maxLen {
			maxLen = l
		}
	}

	require.Equal(t, wantMinLen, minLen)
	require.Equal(t, wantMaxLen, maxLen)
}

func TestRight(t *testing.T) {
	t.Parallel()

	wantMinLen := 2
	wantMaxLen := 13

	minLen := util.Len(right[0])
	maxLen := util.Len(right[0])

	for _, word := range right {
		l := util.Len(word)

		if l < minLen {
			minLen = l
		}

		if l > maxLen {
			maxLen = l
		}
	}

	require.Equal(t, wantMinLen, minLen)
	require.Equal(t, wantMaxLen, maxLen)
}

func TestGet(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		sep         string
		wantPattern string
	}{
		{"", "[a-z]{3,13}[a-z]{2,13}"},
		{"-", "[a-z]{3,13}-[a-z]{2,13}"},
		{"_", "[a-z]{3,13}_[a-z]{2,13}"},
	}
	for _, tc := range testCases {
		rx := regexp.MustCompile(tc.wantPattern)

		t.Run("", func(t *testing.T) {
			t.Parallel()

			for range 10000 {
				got := Get(tc.sep)
				require.Regexp(t, rx, got)
			}
		})
	}
}
