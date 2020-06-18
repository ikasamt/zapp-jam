package clefs

import (
	"reflect"
	"sort"

	"github.com/ikasamt/goslicer/typeslicer"
)

type Anything struct { //generic.Type
} //generic.Type
type Anythings []Anything

// First is ..
func (anythings Anythings) First() (Anything, bool) {
	if len(anythings) == 0 {
		return Anything{}, false
	}
	return anythings[0], true
}

// Where is ..
func (anythings Anythings) Where(f func(anything Anything) bool) (result Anythings) {
	for _, a := range anythings {
		if f(a) {
			result = append(result, a)
		}
	}
	return
}

// Count is ..
func (anythings Anythings) Count() (counter int) {
	return len(anythings)
}

// CountIf is ..
func (anythings Anythings) CountIf(f func(anything Anything) bool) (counter int) {
	for _, a := range anythings {
		if f(a) {
			counter++
		}
	}
	return
}

// Select is ..
func (anythings Anythings) Select(fieldName string) (result typeslicer.InterfaceSlice) {
	for _, a := range anythings {
		i := reflect.ValueOf(a).FieldByName(fieldName).Interface()
		result = append(result, i)
	}
	return
}

// SortBy is ..
func (anythings Anythings) SortBy(sortFunc func(Anything, Anything) bool) (result Anythings) {
	f := func(i, j int) bool {
		a := anythings[i]
		b := anythings[j]
		return sortFunc(a, b)
	}

	tmp := make(Anythings, len(anythings))
	copy(tmp, anythings)
	sort.Slice(tmp, f)
	return tmp
}

// DistinctBy is ..
func (anythings Anythings) DistinctBy(f func(anything Anything) interface{}) (result Anythings) {
	tmp := map[interface{}]Anything{}
	for _, a := range anythings {
		tmp[f(a)] = a
	}
	for _, t := range tmp {
		result = append(result, t)
	}
	return
}

func (anythings Anythings) Mapper(f func(anything Anything) Anything) (result Anythings) {
	for _, a := range anythings {
		result = append(result, f(a))
	}
	return
}
