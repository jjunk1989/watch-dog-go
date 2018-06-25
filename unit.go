package watch

import (
	"strings"
	"unicode/utf16"
)

func uint16BufToString(buf []uint16) string {
	return runeToString(utf16.Decode(buf))
}

func runeToString(buf []rune) string {
	b := &strings.Builder{}
	for _, r := range buf {
		if r != 0 {
			b.WriteRune(r)
		}
	}
	return b.String()
}
