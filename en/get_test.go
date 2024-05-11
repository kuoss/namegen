package en

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	testCases := []struct {
		sep         string
		wantPattern string
	}{
		{"", "[a-z]{3,}[a-z]{2,}"},
		{"-", "[a-z]{3,}-[a-z]{2,}"},
		{"_", "[a-z]{3,}_[a-z]{2,}"},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			rx := regexp.MustCompile(tc.wantPattern)
			for i := 0; i < 10000; i++ {
				got := Get(tc.sep)
				require.Regexp(t, rx,  got)
			}
		})
	}
}
