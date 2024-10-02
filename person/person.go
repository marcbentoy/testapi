package testperson

type Person struct {
	Name string
}

func NewPerson(name string) *Person {
	return &Person{
		Name: name,
	}
}
