package main

import "fmt"

func bubbleSort(arr []int) {
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr) - 1 - i; j++ {
           if arr[j] > arr[j + 1] {
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
           } 
        }
    }

    for _, val := range arr {
        fmt.Println(val)
    }
}

func main() {
    array := []int{1,2,3,4,2,1,3,5,6,4,7}
    bubbleSort(array)
    
}
