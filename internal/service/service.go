package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data []byte) []byte {
	str := string(data)
	isMorse := strings.ContainsRune(str, '.') || strings.ContainsRune(str, '-')
	var out string
	if isMorse {
		out = morse.ToText(str)
	} else {
		out = morse.ToMorse(str)
	}
	return []byte(out)
}
