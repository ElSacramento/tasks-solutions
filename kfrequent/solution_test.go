package kfrequent

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTopKFrequentNums(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 1, 2, 2, 2, 2, 4, 4, 4, 4, 5, 5, 5, 6, 6}
	result := topKFrequentNums(numbers, 3)
	require.Equal(t, []int{2, 4, 5}, result)

	numbers = []int{1, 1, 1, 2, 2, 3}
	result = topKFrequentNums(numbers, 2)
	require.Equal(t, []int{1, 2}, result)

	numbers = []int{1}
	result = topKFrequentNums(numbers, 1)
	require.Equal(t, []int{1}, result)
}

func TestTopKFrequentWords(t *testing.T) {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	result := topKFrequentWords(words, 2)
	require.Equal(t, []string{"i", "love"}, result)

	words = []string{"i"}
	result = topKFrequentWords(words, 1)
	require.Equal(t, []string{"i"}, result)

	words = []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	result = topKFrequentWords(words, 4)
	require.Equal(t, []string{"the", "is", "sunny", "day"}, result)
}

func TestMostCommonWord(t *testing.T) {
	text := "Bob hit a ball, the hit BALL flew far after it was hit."
	result := mostCommonWord(text, []string{"hit"})
	require.Equal(t, "ball", result)

	text = "a a, b a; b c"
	result = mostCommonWord(text, []string{"a", "b"})
	require.Equal(t, "c", result)

	text = "a? a, b a; b c."
	result = mostCommonWord(text, []string{})
	require.Equal(t, "a", result)

	text = "a, a, a, a, b,b,b,c, c"
	result = mostCommonWord(text, []string{"a"})
	require.Equal(t, "b", result)
}

func TestReorganizeString(t *testing.T) {
	require.Equal(t, "aba", reorganizeString("aab"))
	require.Equal(t, "", reorganizeString("aaab"))
	require.Equal(t, "", reorganizeString("aa"))
	require.Equal(t, "a", reorganizeString("a"))

	{
		result := reorganizeString("aabbc")
		require.True(t, result == "ababc" || result == "babac")
	}
}
