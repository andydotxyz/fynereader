package fynereader

import (
	"strings"
)

func textWrap(in string) string {
	wrapChars := 120
	if len(in) < wrapChars {
		return in
	}

	off := 0
	out := ""
	for off+wrapChars < len(in) {
		newline := strings.LastIndex(in[off:off+wrapChars], "\n")
		if newline != -1 { // if there is already a newline in this chunk then use it
			out += in[off : off+newline+1]
			off += newline + 1
			continue
		}

		// check for the nearest word boundary, if there is one
		space := strings.LastIndex(in[off:off+wrapChars], " ")
		chunk := wrapChars
		found := false
		if space != -1 {
			chunk = space
			found = true
		}

		out += in[off:off+chunk] + "\n"
		off += chunk
		if found {
			off++
		}
	}

	return out + in[off:]
}

func removeWhitespace(in string) string {
	in = strings.ReplaceAll(in, "\n", " ")
	return in
}

func stripTags(in string) string {
	if len(in) < 3 || in[0] != '<' {
		return in
	}

	ret := ""
	tagStart := strings.Index(in, "<")
	for tagStart > -1 {
		if tagStart > 0 {
			ret += removeWhitespace(in[:tagStart])
		}
		tagEnd := tagStart + 1 + strings.Index(in[tagStart+1:], ">")
		tag := in[tagStart+1 : tagEnd]
		if tag == "/p" || tag == "/h1" || tag == "/h2" || tag == "/h3" || tag == "/h4" || tag == "/h5" || tag == "/h6" {
			ret += "\n\n"
		}
		if tag == "br" || tag == "br/" || tag == "/ul" || tag == "/ol" || tag == "/li" {
			ret += "\n"
		}

		in = in[tagEnd+1:]
		tagStart = strings.Index(in, "<")
	}
	return ret + removeWhitespace(in)
}
