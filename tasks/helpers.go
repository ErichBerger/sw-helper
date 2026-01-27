package tasks

import (
	"io"
	"os"
	"strings"
	"unicode"
)

func ensureTrailingNewline(path string, w io.Writer) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if len(data) == 0 {
		return nil
	}

	last := data[len(data)-1]

	// Handle both LF and CRLF
	if last != '\n' {
		_, err = w.Write([]byte("\n"))
		return err
	}

	return nil
}
func pascalCaseToKebab(s string) string {
	runes := []rune(s)
	builder := strings.Builder{}

	for i, r := range runes {
		var prev, next rune
		if i > 0 {
			prev = runes[i-1]
		}
		if i+1 < len(runes) {
			next = runes[i+1]
		}

		if unicode.IsLower(r) {
			builder.WriteRune(r)
			continue
		}

		if unicode.IsNumber(r) && i+1 < len(runes) && unicode.IsUpper(next) {
			builder.WriteRune(r)
			builder.WriteString("-")
			continue
		}

		if unicode.IsUpper(r) && i > 0 && unicode.IsLower(prev) {
			builder.WriteString("-")
			builder.WriteRune(unicode.ToLower(r))
			continue
		}

		if unicode.IsUpper(r) && i+1 < len(runes) && unicode.IsLower(next) && i > 0 && unicode.IsUpper(prev) {
			builder.WriteString("-")
			builder.WriteRune(unicode.ToLower(r))
			continue
		}
		builder.WriteRune(unicode.ToLower(r))
	}
	return builder.String()
}
