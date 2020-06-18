package clefs

// Anything is defined as base Struct for genny
type Anything struct { //generic.Type
	ID         int       //generic.Type
} //generic.Type


func (any Anything) Is(ID int) bool{
	return  any.ID == ID
}