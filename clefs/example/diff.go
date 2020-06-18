package clefs

func (any Anything) Diff(ID int) bool{
	return  any.ID != ID
}
