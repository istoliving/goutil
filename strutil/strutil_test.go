package strutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestTrim(t *testing.T) {
	is := assert.New(t)

	// Trim
	tests := map[string]string{
		"abc ":  "",
		" abc":  "",
		" abc ": "",
		"abc,,": ",",
		"abc,.": ",.",
	}
	for sample, cutSet := range tests {
		is.Equal("abc", strutil.Trim(sample, cutSet))
	}

	is.Equal("abc", strutil.Trim("abc,.", ".,"))
	// is.Equal("", Trim(nil))

	// TrimLeft
	is.Equal("abc ", strutil.TrimLeft(" abc "))
	is.Equal("abc ,", strutil.TrimLeft(", abc ,", " ,"))
	is.Equal("abc ,", strutil.TrimLeft(", abc ,", ", "))

	// TrimRight
	is.Equal(" abc", strutil.TrimRight(" abc "))
	is.Equal(", abc", strutil.TrimRight(", abc ,", ", "))
}

func TestURLEnDecode(t *testing.T) {
	is := assert.New(t)

	is.Equal("a.com/?name%3D%E4%BD%A0%E5%A5%BD", strutil.URLEncode("a.com/?name=你好"))
	is.Equal("a.com/?name=你好", strutil.URLDecode("a.com/?name%3D%E4%BD%A0%E5%A5%BD"))
	is.Equal("a.com", strutil.URLEncode("a.com"))
	is.Equal("a.com", strutil.URLDecode("a.com"))
}

func TestFilterEmail(t *testing.T) {
	is := assert.New(t)
	is.Equal("THE@inhere.com", strutil.FilterEmail("   THE@INHere.com  "))
	is.Equal("inhere.xyz", strutil.FilterEmail("   inhere.xyz  "))
}

func TestSimilarity(t *testing.T) {
	is := assert.New(t)
	_, ok := strutil.Similarity("hello", "he", 0.3)
	is.True(ok)
}

func TestSplit(t *testing.T) {
	ss := strutil.Split("a, , b,c", ",")
	assert.Equal(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.Split(" ", ",")
	assert.Nil(t, ss)
}

func TestSubstr(t *testing.T) {
	assert.Equal(t, "abc", strutil.Substr("abcDef", 0, 3))
	assert.Equal(t, "cD", strutil.Substr("abcDef", 2, 2))
}

func TestRepeat(t *testing.T) {
	assert.Equal(t, "aaa", strutil.Repeat("a", 3))
	assert.Equal(t, "D", strutil.Repeat("D", 1))
	assert.Equal(t, "D", strutil.Repeat("D", 0))
	assert.Equal(t, "D", strutil.Repeat("D", -3))
}

func TestPadding(t *testing.T) {
	tests := []struct {
		want, give, pad string
		len             int
		pos             uint8
	}{
		{"ab000", "ab", "0", 5, strutil.PosRight},
		{"000ab", "ab", "0", 5, strutil.PosLeft},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, strutil.Padding(tt.give, tt.pad, tt.len, tt.pos))

		if tt.pos == strutil.PosRight {
			assert.Equal(t, tt.want, strutil.PadRight(tt.give, tt.pad, tt.len))
		} else {
			assert.Equal(t, tt.want, strutil.PadLeft(tt.give, tt.pad, tt.len))
		}
	}
}
