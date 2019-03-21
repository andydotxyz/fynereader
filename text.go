package fynereader

import "strings"

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
