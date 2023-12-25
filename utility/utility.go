package utility

import "reflect"

// Pair which can save any couple of elements
type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

// Tuple which can save any elements
type Tuple struct {
	Items []any
}

// InitTuple Initial the tuple
func (t *Tuple) InitTuple(args ...any) {
	t.Items = make([]any, 0)
	for _, item := range args {
		t.Items = append(t.Items, item)
	}
}

// Compare with another tuple
func (t *Tuple) Compare(other Tuple) (bool, string) {
	if len(t.Items) != len(other.Items) {
		return false, "Length of two tuples are not equal."
	}
	for index, items := range t.Items {
		if reflect.TypeOf(items) != reflect.TypeOf(other.Items[index]) {
			return false, "Type of two tuples are not equal."
		}
	}
	for index, item := range t.Items {
		if item != other.Items[index] {
			return false, "Value of two tuples are not equal."
		}
	}
	return true, "Two tuples are equal."
}
