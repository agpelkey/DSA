package main

import "fmt"

func bubblesort(arr []int) []int {
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr) - i - 1; j++ {
            if arr[j] > arr[j + 1] {
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
            }
        }
    }
    return arr

}


func main() {
    arr := []int{1,2,3,4,2,1,2,7,5,3}
    fmt.Println(bubblesort(arr))
}
