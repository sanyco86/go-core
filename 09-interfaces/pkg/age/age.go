package age

// Person - интерфейс человека.
type Person interface {
	Age() int
}

// Employee - сотрудник.
type Employee struct {
	name string
	age  int
}

// Age - геттер поля age для сотрудника.
func (e Employee) Age() int {
	return e.age
}

// Customer - клиент.
type Customer struct {
	premium bool
	age     int
}

// Age - геттер поля age для клиента.
func (c Customer) Age() int {
	return c.age
}

// MaxAge - возвращает возраст самого старшего человека.
func MaxAge(people ...Person) int {
	var max int

	for _, person := range people {
		if max < person.Age() {
			max = person.Age()
		}
	}

	return max
}
