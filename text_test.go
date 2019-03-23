package fynereader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextWrap(t *testing.T) {
	long := "0123456789012345678901234567890123456789012345678901234567890123456789" +
		"012345678901234567890123456789012345678901234567890123456789"

	wrapped := textWrap(long)
	assert.Equal(t, "0123456789012345678901234567890123456789012345678901234567890123456789"+
		"01234567890123456789012345678901234567890123456789\n0123456789", wrapped)
}

func TestTextWrap_Space(t *testing.T) {
	long := "0123456789 0123456789 0123456789 0123456789 0123456789 0123456789 0123456789 " +
		"0123456789 0123456789 0123456789 0123456789 0123456789 0123456789"

	wrapped := textWrap(long)
	assert.Equal(t, "0123456789 0123456789 0123456789 0123456789 0123456789 0123456789 0123456789 "+
		"0123456789 0123456789 0123456789\n0123456789 0123456789 0123456789", wrapped)
}

func TestTextWrap_WithNewline(t *testing.T) {
	long := "0123456789 0123456789 0123456789 0123456789 0123456789\n0123456789 0123456789 " +
		"0123456789 0123456789 0123456789 0123456789 0123456789 0123456789"

	wrapped := textWrap(long)
	assert.Equal(t, "0123456789 0123456789 0123456789 0123456789 0123456789\n0123456789 0123456789 "+
		"0123456789 0123456789 0123456789 0123456789 0123456789 0123456789", wrapped)
}

func TestStripTags_P(t *testing.T) {
	in := "<p>Hello</p>"
	stripped := stripTags(in)

	assert.Equal(t, "Hello\n\n", stripped)
}

func TestStripTags_P_single(t *testing.T) {
	in := "<p>Hello<p>"
	stripped := stripTags(in)

	assert.Equal(t, "Hello", stripped)
}

func TestStripTags_A(t *testing.T) {
	in := "<a href=\"blah\">Hello</a>"
	stripped := stripTags(in)

	assert.Equal(t, "Hello", stripped)
}

func TestStripTags_Inline(t *testing.T) {
	in := "<p>No<span>-</span>Line</p>"
	stripped := stripTags(in)

	assert.Equal(t, "No-Line\n\n", stripped)
}

func TestStripTags_Newlines(t *testing.T) {
	in := "<p>No\n<span>-</span>\nLine</p>"
	stripped := stripTags(in)

	assert.Equal(t, "No - Line\n\n", stripped)
}

func TestStripTags_NewPara(t *testing.T) {
	in := "<p>One\n<span>New</span></p><p>Line</p>"
	stripped := stripTags(in)

	assert.Equal(t, "One New\n\nLine\n\n", stripped)
}

func TestStripTags_NewBreak(t *testing.T) {
	in := "<p>One\n<span>New</span><br/>Line</p>"
	stripped := stripTags(in)

	assert.Equal(t, "One New\nLine\n\n", stripped)
}
