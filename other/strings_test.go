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

func TestAlphabetBoardPath(t *testing.T) {
	{
		st := "leet"
		expected := "DDR!UURRR!!DDD!"
		require.Equal(t, expected, alphabetBoardPath(st))
	}
	{
		st := "code"
		expected := "RR!DDRR!UUL!R!"
		require.Equal(t, expected, alphabetBoardPath(st))
	}
	{
		st := "a"
		expected := "!"
		require.Equal(t, expected, alphabetBoardPath(st))
	}
	{
		st := "m"
		expected := "DDRR!"
		require.Equal(t, expected, alphabetBoardPath(st))
	}
	{
		st := "dzy"
		expected := "RRR!LLLDDDDD!URRRR!"
		require.Equal(t, expected, alphabetBoardPath(st))
	}
}

func TestReverseString(t *testing.T) {
	{
		s := []byte{'h', 'e', 'l', 'l', 'o'}
		reverseString(s)
		require.Equal(t, []byte{'o', 'l', 'l', 'e', 'h'}, s)
	}
	{
		s := []byte{'H', 'e', 'l', 'l'}
		reverseString(s)
		require.Equal(t, []byte{'l', 'l', 'e', 'H'}, s)
	}
	{
		s := []byte{'h'}
		reverseString(s)
		require.Equal(t, []byte{'h'}, s)
	}
}

func TestAnagram(t *testing.T) {
	{
		s := "anagram"
		tr := "nagrama"
		require.True(t, isAnagram(s, tr))
	}
	{
		s := "anagram"
		tr := "nangram"
		require.False(t, isAnagram(s, tr))
	}
	{
		s := "a"
		tr := "b"
		require.False(t, isAnagram(s, tr))
	}
	{
		s := "a"
		tr := "a"
		require.True(t, isAnagram(s, tr))
	}
	{
		s := "aaan"
		tr := "annn"
		require.False(t, isAnagram(s, tr))
	}
	{
		s := "aaan"
		tr := "aaan"
		require.True(t, isAnagram(s, tr))
	}
}