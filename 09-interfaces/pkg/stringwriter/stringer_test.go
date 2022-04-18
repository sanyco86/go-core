package stringer

import (
	"strings"
	"testing"
)

func TestWrite(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
		want string
	}{
		{
			name: "#1 строка",
			args: []interface{}{"воробей"},
			want: "воробей",
		},
		{
			name: "#2 строки",
			args: []interface{}{"воробей ", "и", " ", "голубь"},
			want: "воробей и голубь",
		},
		{
			name: "#3 строки и числа",
			args: []interface{}{"ласточка", nil, 42, 3.14},
			want: "ласточка",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := new(strings.Builder)

			Write(w, tt.args...)

			if got := w.String(); got != tt.want {
				t.Errorf("Write() = %v, want %v", got, tt.want)
			}
		})
	}
}
