package main

import "fmt"

// heap struct
type MaxHeap struct {
    array []int
}

// === functions === 

// parent
func parent(i int) int {
    return (i - 1)/2
}

// left
func left(i int) int {
    return (2 * i) + 1
}

// right
func right(i int) int {
    return (2 * i) + 2
}

// === methods ===

// Swap
func (h *MaxHeap) swap (i1, i2 int) {
   h.array[i1], h.array[i2] =  h.array[i2], h.array[i1]
}

// Insert
func (h *MaxHeap) Insert(key int) {
    h.array = append(h.array, key)
    // after adding to the heap we need to heapify
    h.maxHeapifyUp(len(h.array) - 1)
}

// Extract 
func (h *MaxHeap) Extract() int {
    extracted := h.array[0]
    l := len(h.array) - 1

    if len(h.array) == 0 {
        fmt.Println("cannot extract")
        return -1
    }
    h.array[0] = h.array[l] // take out left index and put it in root
    h.array = h.array[:l]

    h.maxHeapifyDown(0)
    return extracted
}

// maxHeapifyUp
func (h *MaxHeap) maxHeapifyUp(index int) {
    for h.array[parent(index)] < h.array[index] {
        h.swap(parent(index), index)
        index = parent(index)
    }
}

// maxHeapifyDown
// will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
    lastIndex := len(h.array) - 1
    l, r := left(index), right(index)
    childToCompare := 0

    // loop while index has at least one child
    for l <= lastIndex {
        if l == lastIndex { // when left child is the only child
            childToCompare = l
        } else if h.array[l] > h.array[r] { // when left child is larger
            childToCompare = l
        } else { // when right child is larger
            childToCompare = r
        }
    }

    // compare array value of current index to larger child and swap if smaller
    if h.array[index] < h.array[childToCompare] {
        h.swap(index, childToCompare)
        index = childToCompare
        l, r = left(index), right(index)
    } else {
        return
    }
}



func main() {
    m := &MaxHeap{}
    fmt.Println(m)

    buildHeap := []int{10, 20, 40, 6, 2, 9}
    for _, v := range buildHeap {
        m.Insert(v)
        fmt.Println(m)
    }

    for i := 0; i < 5; i++ {
        m.Extract()
        fmt.Println(m)
    }
}













