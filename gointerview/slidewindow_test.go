package interview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlideWindow(t *testing.T) {
	str := "abcabcbb"
	assert.Equal(t, 3, lengthOfLongestSubstring(str))

	str = "pwwkew"
	assert.Equal(t, 3, lengthOfLongestSubstring(str))

	str = "abba"
	assert.Equal(t, 2, lengthOfLongestSubstring(str))

	str = " "
	assert.Equal(t, 1, lengthOfLongestSubstring(str))
}
