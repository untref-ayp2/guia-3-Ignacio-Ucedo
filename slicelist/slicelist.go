package slicelist

//import (
//	"errors"
//	"fmt"
//)

// Implementar la lista sobre slices

type SliceList[T comparable] struct {
	list []T // slice que formará la lista
}

//Crea una nueva lista. O(1)
func NewSliceList[T comparable]() *SliceList[T] {
	return &SliceList[T]{list: nil}
}

//Agrega un elemento al final de la lista. O(1)
func (l *SliceList[T]) Append(value T) {

	l.list = append(l.list, value)
}

//Agrega un elemento al principio de la lista. O(1)
func (l *SliceList[T]) Prepend(value T) {

	l.list = append([]T{value}, l.list...)
}

//Agrega un elemento en posición específica. O(1)
func (l *SliceList[T]) InsertAt(value T, position int) {

	temp := l.list[position:]  // segunda parte del slice
	l.list = l.list[:position] // primera parte del slice

	l.list = append(l.list, value)
	l.list = append(l.list, temp...)

}

//Remueve el primer elemento con el valor solicitado de la lista. O(n)
func (l *SliceList[T]) Remove(value T) {

	for i := 0; i < len(l.list); i++ {
		if l.list[i] == value {
			l.list = append(l.list[:i], l.list[i+1:]...)
			return
		}
	}
}

//Devuelve el elemento en la posición especificada. O(1)
func (l *SliceList[T]) Get(position int) (T, error) {

	return l.list[position], nil
}

//Devuelve el tamaño de la lista. O(1)
func (l *SliceList[T]) Size() int {

	return len(l.list)
}

//concatena una lista a la lista. O(1)
func (l *SliceList[T]) ConcatenarLista(e SliceList[T]) {
	l.list = append(l.list, e.list...)
}

//Intercala los valores de una lista a la lista. O(n)
func (l *SliceList[T]) IntercalarLista(e SliceList[T]) {
	var lista1 []T
	var lista2 []T

	if len(l.list) > len(e.list) {
		lista1 = l.list
		lista2 = e.list
	} else {
		lista2 = l.list
		lista1 = e.list
	}
	//lista1 siempre es la mayor
	l.list = nil
	for i := 0; i < len(lista2); i++ {

		l.list = append(l.list, lista1[0], lista2[i])
		lista1 = lista1[1:]
	}

	l.list = append(l.list, lista1...)

}
