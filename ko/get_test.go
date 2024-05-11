package ko

import (
	"regexp"
	"testing"

	"github.com/kuoss/namegen/util"
	"github.com/stretchr/testify/require"
)

var (
	wantLeftMinLen  = 1
	wantLeftMaxLen  = 6
	wantRightMinLen = 1
	wantRightMaxLen = 5
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
		{"", "[가-힣]{1,6}[가-힣]{1,5}"},
		{"-", "[가-힣]{1,6}-[가-힣]{1,5}"},
		{"_", "[가-힣]{1,6}_[가-힣]{1,5}"},
		{"hello", "[가-힣]{1,6}hello[가-힣]{1,5}"},
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
