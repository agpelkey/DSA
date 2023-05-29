package main

import "fmt"

func partition(arr []int, lo, hi int) int {
    pivot := arr[hi]
    for j := lo; j < hi; j++ {
        if arr[j] < pivot {
            arr[j], arr[lo] = arr[lo], arr[j]
            lo++
        }
    }

    arr[lo], arr[hi] = arr[hi], arr[lo]
    return lo
    
}


func quickSort(arr []int, lo, hi int) {
        if lo > hi {
            return
        }

        pivotIndex := partition(arr, lo, hi)
        quickSort(arr, lo, pivotIndex - 1)
        quickSort(arr, pivotIndex + 1, hi)
}



func main() {
    lst := []int{1,5,2,4,8,9,7,4,1,2}

    quickSort(lst, 0, len(lst) - 1)
    fmt.Println(lst)
}



