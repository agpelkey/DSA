package main

import "fmt"

//import "fmt"
/*
func binarySearch(arr []int, needle int) (int, bool) {
    lo := 0
    hi := len(arr) - 1
    for lo <= hi {
        mid := (lo + hi)/2

        if arr[mid] == needle {
            return arr[mid], true
        } else if arr[mid] > needle {
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }

    return -1, false
}


func main() {
    arr := []int{1,2,3,4,5,6,7,8,9,10}
    fmt.Println(binarySearch(arr, 3))

}
*/


func binarySearch(haystack []int, needle int) bool {
    lo := 0
    hi := len(haystack) - 1
    for lo <= hi {
        mid := (lo + hi)/2
        if haystack[mid] == needle {
            return true
        } else if haystack[mid] > needle {
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }

    return false
}

func main() {
    arr := []int{1,2,3,4,5,6,7,8,9,10}
    fmt.Println(binarySearch(arr, 7))
}









