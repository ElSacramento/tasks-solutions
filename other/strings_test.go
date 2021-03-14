package other

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBreakPalindrome(t *testing.T) {
	require.Equal(t, "aab", breakPalindrome("aaa"))
	require.Equal(t, "acb", breakPalindrome("aca"))
	require.Equal(t, "aacba", breakPalindrome("abcba"))
	require.Equal(t, "aaddca", breakPalindrome("acddca"))
	require.Equal(t, "acce", breakPalindrome("ecce"))
}
