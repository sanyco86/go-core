package stringer

import (
	"io"
)

// Write - prints only strings
func Write(w io.Writer, args ...interface{}) {
	for _, arg := range args {
		if val, ok := arg.(string); ok {
			w.Write([]byte(val))
		}
	}
}
