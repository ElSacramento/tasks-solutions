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

func TestLengthOfLongestSubstring(t *testing.T) {
	require.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	require.Equal(t, 6, lengthOfLongestSubstring("#x Ab9"))
	require.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	require.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	require.Equal(t, 4, lengthOfLongestSubstring("abcda"))
	require.Equal(t, 3, lengthOfLongestSubstring("abcba"))
	require.Equal(t, 0, lengthOfLongestSubstring(""))
}

func TestLetterCombinations(t *testing.T) {
	{
		st := "23"
		expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
		require.Equal(t, expected, letterCombinations(st))
	}
	{
		st := "9"
		expected := []string{"w", "x", "y", "z"}
		require.Equal(t, expected, letterCombinations(st))
	}
	{
		require.Equal(t, []string{}, letterCombinations(""))
	}
	{
		st := "37"
		expected := []string{
			"dp", "dq", "dr", "ds",
			"ep", "eq", "er", "es",
			"fp", "fq", "fr", "fs",
		}
		require.Equal(t, expected, letterCombinations(st))
	}
	{
		st := "372"
		expected := []string{
			"dpa", "dpb", "dpc",
			"dqa", "dqb", "dqc",
			"dra", "drb", "drc",
			"dsa", "dsb", "dsc",

			"epa", "epb", "epc",
			"eqa", "eqb", "eqc",
			"era", "erb", "erc",
			"esa", "esb", "esc",

			"fpa", "fpb", "fpc",
			"fqa", "fqb", "fqc",
			"fra", "frb", "frc",
			"fsa", "fsb", "fsc",
		}
		require.Equal(t, expected, letterCombinations(st))
	}
}
