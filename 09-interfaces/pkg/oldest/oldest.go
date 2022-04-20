package oldest

// Employee - сотрудник.
type Employee struct {
	name string
	age  int
}

// Customer - клиент.
type Customer struct {
	premium bool
	age     int
}

// Oldest - возвращает самого старшего человека.
func Oldest(people ...interface{}) interface{} {
	var maxAge int
	var oldest interface{}

	for _, person := range people {
		if p, ok := person.(Employee); ok {
			if p.age > maxAge {
				oldest = p
				maxAge = p.age
			}
		}

		if p, ok := person.(Customer); ok {
			if p.age > maxAge {
				oldest = p
				maxAge = p.age
			}
		}
	}

	return oldest
}
