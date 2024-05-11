package en

import (
	"regexp"
	"testing"

	"github.com/kuoss/namegen/util"
	"github.com/stretchr/testify/require"
)

var (
	wantLeftMinLen  = 3
	wantLeftMaxLen  = 13
	wantRightMinLen = 2
	wantRightMaxLen = 13
)

func TestLeft(t *testing.T) {
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
	require.Equal(t, wantLeftMinLen, minLen)
	require.Equal(t, wantLeftMaxLen, maxLen)
}

func TestRight(t *testing.T) {
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
	require.Equal(t, wantRightMinLen, minLen)
	require.Equal(t, wantRightMaxLen, maxLen)
}

func TestGet(t *testing.T) {
	testCases := []struct {
		sep         string
		wantPattern string
	}{
		{"", "[a-z]{3,13}[a-z]{2,13}"},
		{"-", "[a-z]{3,13}-[a-z]{2,13}"},
		{"_", "[a-z]{3,13}_[a-z]{2,13}"},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			rx := regexp.MustCompile(tc.wantPattern)
			for i := 0; i < 10000; i++ {
				got := Get(tc.sep)
				require.Regexp(t, rx, got)
			}
		})
	}
}
