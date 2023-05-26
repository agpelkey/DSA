package main

import "fmt"

// Given an array nums of size n, return the majority element
// The Majority element is the element that appears more than (n/2) times.
// You may assume that the majority element always exists in the array

func majorityElement(nums []int) int {
    m := make(map[int]int)
    for _, n := range nums {m[n]++}

    res := 0
    max := 0
    for i, n := range m {
        if n > max {
            max = n
            res = i
    }
    } 
    return res
}


func main() {
    arr := []int{1,2, 2, 2,3,3, 3, 3, 3, 4,5}
    fmt.Println(majorityElement(arr))
}
