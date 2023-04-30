package slicelist

import "guia3/linkedlist"

//implementación de pila sobre lista enlazada simple.
type Stack[T comparable] struct {
	items linkedlist.LinkedList[T]
}

//agrega un elemento a la pila. O(1)
func (pila Stack[T]) Push(e T) {
	pila.items.Append(e)
}

//elimina un elemento de la pila. O(n)
func (pila Stack[T]) Pop() {
	pila.items.Pop()
}

// Devuelve si la pila está vacía. O(1)
func (pila Stack[T]) IsEmpty() bool {
	return pila.items.Size() == 0
}

// Devuelve el elemento del tope de la pila. O(n)
func (pila Stack[T]) Top() T {

	top, _ := pila.items.Get(pila.items.Size() - 1)

	return top
}
