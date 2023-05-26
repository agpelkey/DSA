package main

import "fmt"
/*
// This was an attempt at implementing a stack as a linked list.
// Ran into issues, but it might actually be functional
// create node object
type Node struct {
    value interface{}
    next *Node
}

// create stack
type Stack struct {
    head *Node
    size int
}

// push
func (s *Stack) Push(value interface{}) {
    s.head = &Node{value: s.head}
    s.size++
}

// pop
func (s *Stack) Pop()  (interface{}, bool) { 
    if s.head == nil {
        fmt.Println("stack is empty")
        return 0, false
    }

    value := s.head.value
    s.head = s.head.next
    s.size--
    return value, true
}    
// peek
func (s *Stack) Peek() interface{} {
    if s.head == nil {
        fmt.Println("stack is empty")
        return 0
    } 

    return s.head.value
}


func main() {

    var s Stack

    s.Push(3)
    s.Push(4)

   fmt.Println(s) 
    
    s.Pop()

    fmt.Println(s)
}
*/


// ====================================================
// The following is a stack implemented using a slice in Golang

/*
type Stack []int

// Pop
func (s Stack) Pop() (Stack, int) {
   if len(s) == 0 {
       return Stack{}, -1
   }
    
   l := len(s)
   return s[:l-1], s[l-1]
}


// Push
func (s Stack) Push(i int) Stack{
    return append(s, i)
}

// Peek


func main() {

    s := make(Stack, 0)
    s = s.Push(3)
    s = s.Push(132)
    s = s.Push(45)
    s = s.Push(2)

    fmt.Println(s)

    s, p := s.Pop()
    fmt.Println(p)

    fmt.Println(s)
    
}
*/

/*
type Stack []int

// Pop
func (s Stack) Pop() int {
   if len(s) == 0 {
       fmt.Println("Cannot pop stack is empty")
   }

   l := len(s)
   return s[l-1]
}
// Push
func (s Stack) Push(elem int) Stack {
    return append(s, elem)
}

// Peek
func (s Stack) Peek() int {
    return s[0]
}

func main() {

    s := make(Stack, 0)
    s = s.Push(1) 
    s = s.Push(2) 
    s = s.Push(3) 
    s = s.Push(4) 
    s = s.Push(5) 

    fmt.Println(s)

    s.Pop()
    s.Pop()
    s.Pop()

    for _, v := range s {
        s.Pop()
        fmt.Println(v)
    }
}
*/
/*
=========================================================
type Stack []int

// Pop
func (s *Stack) Pop() (int, bool) {
    if len(*s) == 0 {
        return -1, false
    } else {
        index := len(*s) - 1
        element := (*s)[index]
        *s = (*s)[:index]
        return element, true
    }
}

// Push
func (s Stack) Push(v int) Stack {
    return append(s, v)
}

// Peek
func (s Stack) Peek() int  {
    if len(s) == 0 {
        fmt.Println("Stack is empty")
    }

    return s[0]
}

func main() {

    s := make(Stack, 0)
    s = s.Push(1)
    s = s.Push(2)
    s = s.Push(3)

    fmt.Println(s)

    for len(s) > 0 {
        x, y := s.Pop()
        if y == true {
            fmt.Println(x)
        }
    }
}

=============================================
*/

type stack []int

// Push
func (s stack) Push(elem int) stack {
   return append(s, elem) 
}

// Pop
func (s *stack) Pop() (int, bool) {
    if len(*s) == 0 {
        return -1, false
    }

    index := len(*s) - 1
    element := (*s)[index]
    *s = (*s)[:index]

    return element, true

}

// Peek
func (s stack) Peek() (int, bool) {
    if len(s) == 0 {
        return -1, false
    }
    return s[0], true
}

func main() {
   var s stack
   s = s.Push(1)
   s = s.Push(2)
   s = s.Push(3)

   fmt.Println(s)

  for len (s) > 0 {
      x, y := s.Pop()
    if y == true {
        fmt.Println(x)
  } 
  }
}








