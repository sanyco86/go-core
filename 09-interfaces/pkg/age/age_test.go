package age

import (
	"testing"
)

func TestMaxAge(t *testing.T) {
	tests := []struct {
		name    string
		people  []Person
		wantAge int
	}{
		{
			name: "Сотрудники.",
			people: []Person{
				Employee{
					name: "Миша",
					age:  17,
				},
				Employee{
					name: "Маша",
					age:  32,
				},
			},
			wantAge: 32,
		},
		{
			name: "Клиенты.",
			people: []Person{
				Customer{
					age:     42,
					premium: false,
				},
				Customer{
					age:     60,
					premium: true,
				},
			},
			wantAge: 60,
		},
		{
			name: "Сотрудники и клиенты.",
			people: []Person{
				Employee{
					name: "Женя",
					age:  17,
				},
				Employee{
					name: "Вася",
					age:  32,
				},
				Customer{
					age:     42,
					premium: false,
				},
				Customer{
					age:     60,
					premium: true,
				},
			},
			wantAge: 60,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAge(tt.people...); got != tt.wantAge {
				t.Errorf("got %v, want %v", got, tt.wantAge)
			}
		})
	}
}
