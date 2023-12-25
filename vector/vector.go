package vector

import "fmt"

// Vector which can save any elements
// the type of elements must be the same
// the type of elements must be comparable
type Vector[T any] struct {
	Items []T
}

// PushBack push an element to the back of the vector
func (v *Vector[T]) PushBack(item ...T) {
	v.Items = append(v.Items, item...)
}

// PushFront push an element to the front of the vector
func (v *Vector[T]) PushFront(item ...T) {
	v.Items = append(item, v.Items...)
}

// Front return the first element of the vector
func (v *Vector[T]) Front() (T, error) {
	if v.Empty() {
		var zero T
		return zero, fmt.Errorf("vector is empty")
	}
	return v.Items[0], nil
}

// Back return the last element of the vector
func (v *Vector[T]) Back() (T, error) {
	if v.Empty() {
		var zero T
		return zero, fmt.Errorf("vector is empty")
	}
	return v.Items[len(v.Items)-1], nil
}

// At return the element at the given index
func (v *Vector[T]) At(index int) (T, error) {
	if v.Empty() {
		var zero T
		return zero, fmt.Errorf("vector is empty")
	}

	if index < 0 || index >= len(v.Items) {
		var zero T
		return zero, fmt.Errorf("index out of range")
	}

	return v.Items[index], nil
}

// Clear the vector
func (v *Vector[T]) Clear() {
	v.Items = nil
}

// Empty return true if the vector is empty
func (v *Vector[T]) Empty() bool {
	return len(v.Items) == 0
}

// Size return the size of the vector
func (v *Vector[T]) Size() int {
	return len(v.Items)
}

// Length return the length of the vector
func (v *Vector[T]) Length() int {
	return len(v.Items)
}

// PopBack pop the last element of the vector
func (v *Vector[T]) PopBack() error {
	if v.Empty() {
		return fmt.Errorf("vector is empty")
	}
	v.Items = v.Items[:len(v.Items)-1]
	return nil
}

// Insert an element at the given position
func (v *Vector[T]) Insert(position int, item T) error {
	if v.Empty() {
		return fmt.Errorf("vector is empty")
	}
	if position < 0 || position > len(v.Items) {
		return fmt.Errorf("index out of range")
	}
	v.Items = append(v.Items[:position], append([]T{item}, v.Items[position:]...)...)
	return nil
}
