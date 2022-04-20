package oldest

import (
	"reflect"
	"testing"
)

func TestOldest(t *testing.T) {
	tests := []struct {
		name   string
		people []interface{}
		want   interface{}
	}{
		{
			name: "Сотрудники.",
			people: []interface{}{
				Employee{
					name: "Миша",
					age:  17,
				},
				Employee{
					name: "Маша",
					age:  32,
				},
			},
			want: Employee{
				name: "Маша",
				age:  32,
			},
		},
		{
			name: "Клиенты.",
			people: []interface{}{
				Customer{
					age:     42,
					premium: false,
				},
				Customer{
					age:     60,
					premium: true,
				},
			},
			want: Customer{
				age:     60,
				premium: true,
			},
		},
		{
			name: "Сотрудники и клиенты.",
			people: []interface{}{
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
			want: Customer{
				age:     60,
				premium: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Oldest(tt.people...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
