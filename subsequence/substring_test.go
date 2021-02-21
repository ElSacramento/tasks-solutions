package subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, LongestSubstring("aaabb", 3))
	assert.Equal(t, 5, LongestSubstring("ababbc", 2))
	assert.Equal(t, 1, LongestSubstring("a", 1))
	assert.Equal(t, 0, LongestSubstring("a", 10))
	assert.Equal(t, 10, LongestSubstring("abaabcdcdaef", 2))
	assert.Equal(t, 4, LongestSubstring("babacbaba", 2))
	assert.Equal(t, 4, LongestSubstring("aadbabacdaab", 2))
}
