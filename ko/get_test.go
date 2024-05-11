package ko //nolint:testpackage // package should be `ko_test` instead of `ko`

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kuoss/namegen/util"
)

func TestLeft(t *testing.T) {
	t.Parallel()

	wantMinLen := 1
	wantMaxLen := 5

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

	wantMinLen := 1
	wantMaxLen := 5

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
		{"", "[가-힣]{1,6}[가-힣]{1,5}"},
		{"-", "[가-힣]{1,6}-[가-힣]{1,5}"},
		{"_", "[가-힣]{1,6}_[가-힣]{1,5}"},
		{"hello", "[가-힣]{1,6}hello[가-힣]{1,5}"},
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
