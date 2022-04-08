package sorter

import (
	"reflect"
	"testing"
)

func TestSortString(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Тест №1",
			args: args{
				s: []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"},
			},
			want: []string{"Alpha", "Bravo", "Delta", "Go", "Gopher", "Grin"},
		},
		{
			name: "Тест №2",
			args: args{
				s: []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", ""},
			},
			want: []string{"", "Alpha", "Bravo", "Go", "Gopher", "Grin"},
		},
		{
			name: "Тест №3",
			args: args{
				s: []string{"Go", "Bravo", "1", "Alpha", "5.6", " "},
			},
			want: []string{" ", "1", "5.6", "Alpha", "Bravo", "Go"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortInts(t *testing.T) {
	i := []int{5, 2, 6, 3, 1, 4}       // тестовый пример
	got := SortInts(i)                 // вызов тестируемого кода
	want := []int{1, 2, 3, 4, 5, 6}    // заранее вычисленный результат
	if !reflect.DeepEqual(got, want) { // сравнение результата с правильным значением
		t.Errorf("SortInts() = %d, want %d", got, want)
	}
}

func BenchmarkSortInts(b *testing.B) {
	data := []int{5, 2, 6, 3, 1, 4}
	for i := 0; i < b.N; i++ {
		res := SortInts(data)
		_ = res
	}
}

func BenchmarkSortFloat64s(b *testing.B) {
	data := []float64{5., 2., 6., 3., 1., 4.}
	for i := 0; i < b.N; i++ {
		res := SortFloat64s(data)
		_ = res
	}
}
