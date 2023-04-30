package tests

//Tests Lista Enlazada

import (
	"fmt"
	"guia3/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	list := linkedlist.NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	if list.Size() != 5 {
		t.Error("Error en Append")
	}

	for i := 0; i < 5; i++ {
		v, _ := list.Get(i)
		if v != i+1 {
			t.Error("Error en Append")
		}
	}
}

func TestDelete(t *testing.T) {
	list := linkedlist.NewLinkedList[int]()
	list.Append(0)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	list.Remove(0)
	list.Remove(2)
	list.Remove(3)

	if list.Size() != 3 {
		t.Error("Error en Remove")
	}

	v, err := list.Get(0)
	fmt.Println(v)
	if err != nil || v != 1 {
		t.Error("Error en Remove")
	}
	v, err = list.Get(1)
	fmt.Println(v)
	if err != nil || v != 4 {
		t.Error("Error en Remove")
	}
}

func TestSize(t *testing.T) {
	l := linkedlist.NewLinkedList[int]()

	assert.Equal(t, 0, l.Size())

	l.Append(3)

	assert.Equal(t, 1, l.Size())

	l.Append(234)
	l.Append(0)

	assert.Equal(t, 3, l.Size())

	l.Remove(0)

	assert.Equal(t, 2, l.Size())

	l.Remove(234)

	assert.Equal(t, 1, l.Size())

	l.Remove(3)

	assert.Equal(t, 0, l.Size())

	//______________________________________________

	l2 := linkedlist.NewLinkedList[int]()
	l2.Prepend(23)

	assert.Equal(t, 1, l2.Size())

	l2.Prepend(233)
	l2.Prepend(512)

	assert.Equal(t, 3, l2.Size())

	l2.Remove(000)

	assert.Equal(t, 3, l2.Size())

	l2.Remove(233)

	assert.Equal(t, 2, l2.Size())

	//______________________________________________

	l3 := linkedlist.NewLinkedList[int]()

	l3.InsertAt(-22, 0)

	assert.Equal(t, 1, l3.Size())

	l3.InsertAt(-22, 0)
	l3.InsertAt(-22, 0)

	assert.Equal(t, 3, l3.Size())

	l3.Pop()
	l3.Pop()

	assert.Equal(t, 1, l3.Size())

}
