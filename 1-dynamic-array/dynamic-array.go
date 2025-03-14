package main

import "fmt"

type DynamicArray struct {
	darray []int
}

func NewDynamicArray(capacity int) *DynamicArray {
	var arr = make([]int, 1)
	if capacity > 0 {
		arr = make([]int, capacity)
	}

	return &DynamicArray{darray: arr}
}

// Get value at index i, assuming i is valid.
func (da *DynamicArray) Get(i int) int {
	return da.darray[i]
}

// Set index i to val, assuming i is valid.
func (da *DynamicArray) Set(i, val int) {
	da.darray[i] = val
}

func (da *DynamicArray) GetSize() int {
	return len(da.darray)
}

func (da *DynamicArray) GetCapacity() int {
	return cap(da.darray)
}

func (da *DynamicArray) PopBack() int {
	lastIndex := da.GetSize() - 1
	lastIndexVal := da.darray[lastIndex]

	// da.darray = append(da.darray[:lastIndex], da.darray[lastIndex+1:]...)
	da.darray = da.darray[:lastIndex]
	return lastIndexVal
}

func (da *DynamicArray) PushBack(n int) {
	currSize := da.GetSize()
	currCap := da.GetCapacity()

	// when full
	if currSize == currCap {
		da.resize()
		da.darray[currSize] = n
	} else if currSize < currCap {
		da.darray = da.darray[:currCap]
		da.darray[currCap-1] = n
	}
}

func (da *DynamicArray) resize() {
	// appending to a full slice doubles its capacity automatically
	da.darray = append(da.darray, 0)
}

func main() {
	arr := NewDynamicArray(3)
	fmt.Println("after create", arr.darray, arr.GetCapacity(), arr.GetSize())
	arr.Set(2, 100)
	fmt.Println("after set", arr.darray, arr.GetCapacity(), arr.GetSize())
	arr.PushBack(999)
	fmt.Println("after pushback", arr.darray, arr.GetCapacity(), arr.GetSize())
	arr.PushBack(556)
	fmt.Println("after pushback again", arr.darray, arr.GetCapacity(), arr.GetSize())
	pop := arr.PopBack()
	fmt.Println("after popback", arr.darray, arr.GetCapacity(), arr.GetSize(), pop)
}
