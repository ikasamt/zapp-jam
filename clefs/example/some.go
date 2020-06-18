package example

type Something struct { //generic.Type
	ID int //generic.Type
} //generic.Type

func (any Anything) WhySomething(ID int, some Something) bool {
	return any.ID == some.ID
}
