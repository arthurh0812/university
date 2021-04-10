package university

type Professor struct {
	*Person
	Salary float32
}

func NewProfessor(p *Person, salary float32) *Professor {
	return &Professor{
		Person: p,
		Salary: salary,
	}
}

func (p *Professor) SetFirstName() string {
	return ""
}