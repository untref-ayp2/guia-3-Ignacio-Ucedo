package slicelist

import "guia3/linkedlist"

//implementación de cola sobre lista enlazada simple.
type Queue[T comparable] struct {
	items linkedlist.LinkedList[T]
}

//Agrega un elemento a la cola. O(1)
func (q Queue[T]) Enqueue(e T) {
	q.items.Prepend(e)
	//Se encola al principio de la lista y se desencola al final.
}

//Elimina el primer elemento de la cola. O(n)
func (q Queue[T]) Dequeue() {
	q.items.Pop()
}

//Devuelve si la cola está vacía. O(1)
func (q Queue[T]) IsEmpty() bool {
	return q.items.Size() == 0
}

//Devuelve el primero en la cola. O(n)
func (q Queue[T]) Front() T {

	front, _ := q.items.Get(0)

	return front
}
