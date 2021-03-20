package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestIsAlphabet(t *testing.T) {
	assert.False(t, strutil.IsAlphabet('9'))
	assert.False(t, strutil.IsAlphabet('+'))

	assert.True(t, strutil.IsAlphabet('A'))
	assert.True(t, strutil.IsAlphabet('a'))
	assert.True(t, strutil.IsAlphabet('Z'))
	assert.True(t, strutil.IsAlphabet('z'))
}

func TestIsAlphaNum(t *testing.T) {
	assert.False(t, strutil.IsAlphaNum('+'))

	assert.True(t, strutil.IsAlphaNum('9'))
	assert.True(t, strutil.IsAlphaNum('A'))
	assert.True(t, strutil.IsAlphaNum('a'))
	assert.True(t, strutil.IsAlphaNum('Z'))
	assert.True(t, strutil.IsAlphaNum('z'))
}

func TestEquals(t *testing.T) {
	assert.True(t, strutil.Equal("a", "a"))
	assert.False(t, strutil.Equal("a", "b"))
}

func TestLen(t *testing.T) {
	str := "Hello, 世界"

	assert.Equal(t, 7, strutil.Len("Hello, "))
	assert.Equal(t, 13, strutil.Len(str))
	assert.Equal(t, 9, strutil.Utf8len(str))
}

func TestStrPos(t *testing.T) {
	// StrPos
	assert.Equal(t, -1, strutil.StrPos("xyz", "a"))
	assert.Equal(t, 0, strutil.StrPos("xyz", "x"))
	assert.Equal(t, 2, strutil.StrPos("xyz", "z"))

	// RunePos
	assert.Equal(t, -1, strutil.RunePos("xyz", 'a'))
	assert.Equal(t, 0, strutil.RunePos("xyz", 'x'))
	assert.Equal(t, 2, strutil.RunePos("xyz", 'z'))
	assert.Equal(t, 5, strutil.RunePos("hi时间", '间'))

	// BytePos
	assert.Equal(t, -1, strutil.BytePos("xyz", 'a'))
	assert.Equal(t, 0, strutil.BytePos("xyz", 'x'))
	assert.Equal(t, 2, strutil.BytePos("xyz", 'z'))
	// assert.Equal(t, 2, strutil.BytePos("hi时间", '间')) // will build error
}

func TestIsStartOf(t *testing.T) {
	tests := []struct {
		give string
		sub  string
		want bool
	}{
		{"abc", "a", true},
		{"abc", "d", false},
	}

	for _, item := range tests {
		assert.Equal(t, item.want, strutil.HasPrefix(item.give, item.sub))
		assert.Equal(t, item.want, strutil.IsStartOf(item.give, item.sub))
	}
}

func TestIsEndOf(t *testing.T) {
	tests := []struct {
		give string
		sub  string
		want bool
	}{
		{"abc", "c", true},
		{"abc", "d", false},
		{"some.json", ".json", true},
	}

	for _, item := range tests {
		assert.Equal(t, item.want, strutil.HasSuffix(item.give, item.sub))
		assert.Equal(t, item.want, strutil.IsEndOf(item.give, item.sub))
	}
}

func TestIsSpace(t *testing.T) {
	assert.True(t, strutil.IsSpace(' '))
	assert.True(t, strutil.IsSpace('\n'))
	assert.True(t, strutil.IsSpaceRune('\n'))
	assert.True(t, strutil.IsSpaceRune('\t'))
	assert.True(t, strutil.IsBlank([]byte(" ")))
	assert.True(t, strutil.IsBlank([]byte("   ")))
}

func TestIsSymbol(t *testing.T) {
	assert.False(t, strutil.IsSymbol('a'))
	assert.True(t, strutil.IsSymbol('●'))
}