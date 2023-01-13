package solution

type Result interface {
	Value() any
}

type IntResult int

func (r IntResult) Value() any {
	return int(r)
}

func (r IntResult) Int() int {
	return int(r)
}
