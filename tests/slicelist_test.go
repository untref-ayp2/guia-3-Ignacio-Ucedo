package tests

// Tests Lista Enlazada

import (
	"fmt"
	"guia3/slicelist"
	"testing"
)

func TestInsertSliceList(t *testing.T) {
	list := slicelist.NewSliceList[int]()
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

func TestDeleteSliceList(t *testing.T) {
	list := slicelist.NewSliceList[int]()
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

func TestConcatenarSliceList(t *testing.T) {
	list := slicelist.NewSliceList[int]()
	list.Append(0)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	list2 := slicelist.NewSliceList[int]()
	list2.Append(0)
	list2.Append(1)
	list2.Append(2)
	list2.Append(3)
	list2.Append(4)
	list2.Append(5)

	list.ConcatenarLista(*list2)

	if list.Size() != 12 {
		t.Error("Error en Concatenear lista")
	}

}

func TestIntercalarSliceList(t *testing.T) {
	list := slicelist.NewSliceList[int]()
	list.Append(0)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	list2 := slicelist.NewSliceList[int]()
	list2.Append(0)
	list2.Append(1)
	list2.Append(2)
	list2.Append(3)
	list2.Append(4)
	list2.Append(5)

	list.IntercalarLista(*list2)

	if list.Size() != 12 {
		t.Error("Error en Intercalar lista")
	}

}

func TestIntercalarSliceList2(t *testing.T) {
	list := slicelist.NewSliceList[int]()
	list.Append(0)
	list.Append(1)
	list.Append(2)

	list2 := slicelist.NewSliceList[int]()
	list2.Append(0)
	list2.Append(1)
	list2.Append(2)
	list2.Append(3)
	list2.Append(4)
	list2.Append(5)

	list.IntercalarLista(*list2)

	if list.Size() != 9 {
		t.Error("Error en Intercalar lista")
	}

}
