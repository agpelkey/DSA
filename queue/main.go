package main

import "fmt"
/*
type Queue struct {
    Elements []int
}


// Enqueue
func (q *Queue) Enqueue(elem int) {
    // check for overflow condition
    q.Elements = append(q.Elements, elem)
}

// Dequeue
func (q *Queue) Dequeue() (int, bool) {

    element := q.Elements[0]
    if len(q.Elements) == 1 {
        q.Elements = nil
        return element, false
    }

    q.Elements = q.Elements[1:]
    return element, true
        
}

// Peek
func (q *Queue) Peek() int {
    return q.Elements[0]
}


func main() {
        
    q := Queue{}
    fmt.Println(q)
    q.Enqueue(69)
    q.Enqueue(2)
    q.Enqueue(49)
    q.Enqueue(11)
    q.Enqueue(30)

   fmt.Println(q) 

   q.Dequeue()

   fmt.Println(q)

   q.Dequeue()

   fmt.Println(q)

}
*/

/*
type Queue struct {
    Elements []int
}


// Enqueue
func (q *Queue) Enqueue(elem int) {
    q.Elements = append(q.Elements, elem)
}

// Dequeue
func (q *Queue) Dequeue() (int, bool) {
   // grab item that is going to be dequeued
   element := q.Elements[0]

   // check if length is null
   if q.Elements == nil {
       fmt.Println("Error: queue empty")
       return -1, false
   }

   // create new queue
   q.Elements = q.Elements[1:]
   return element, true
}

// Peek
func (q *Queue) Peek() int {
    return q.Elements[0]
}

func main() {

    q := Queue{}
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    q.Enqueue(4)
    q.Enqueue(5)

    fmt.Println(q)

    q.Dequeue()
    q.Dequeue()
    q.Dequeue()
    q.Dequeue()

    fmt.Println(q)


}
*/

/*
========================================================
type Queue struct {
    Elements []int
}

// Enqueue
func (q *Queue) Enqueue(elem int) {
    q.Elements = append(q.Elements, elem)
}

// Dequeue
func (q *Queue) Dequeue() (int, bool) {
    if q.Elements == nil {
        return -1, false
    }

    element := q.Elements[0]

    q.Elements = q.Elements[1:]
    return element, true
}

// Peek

func main() {

    q := Queue{}

    fmt.Println(q)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    q.Enqueue(4)
    q.Enqueue(5)

    fmt.Println(q)
   
    for _, v := range q.Elements {
        q.Dequeue()
        fmt.Println(v)
    }

}
=======================================================
*/

type Queue struct {
    Elements []int
}

// Enqueue 
func (q *Queue) Enqueue(elem int) {
    q.Elements = append(q.Elements, elem)
}

// Dequeue
func (q *Queue) Dequeue() (int, bool) {
    if q.Elements == nil {
        return -1, false
    }

    // get element thats "first in line"
    element := q.Elements[0]

    // modify list to no longer include element 
    q.Elements = q.Elements[1:]
    return element, true

}

// Peek
func (q *Queue) Peek() int {
    return q.Elements[0]
}

func main() {
    q := Queue{}
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    q.Enqueue(4)

    fmt.Println(q)

    for _, v := range q.Elements {
       q.Dequeue()
       fmt.Println(v)
    }

}



// ============================================
// the following is a Queue implemented as a linked list
/*
type Node struct {
    value int
    next *Node
}

type Queue struct {
    head *Node
    tail *Node
    size int
}

// Enqueue
func (q *Queue) Enqueue(val int) {
    newNode := Node{value: val}
    if q.size == 0 {
        q.head = &newNode
        q.tail = &newNode
    } else {
        q.tail.next = &newNode
        q.tail = &newNode
    }
    q.size++
}

// Dequeue
func (q *Queue) Dequeue() (int, bool) {
    if q.size == 0 {
        return -1, false
    }

    value := q.head.value
    q.head = q.head.next
    q.size--

    return value, true
}

// Peek
func (q *Queue) Peek() int {
    return q.head.value
}

func (q *Queue) Size() int {
    return q.size
}

func main() {
    q := &Queue{}
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    q.Enqueue(4)
    q.Enqueue(5)
    fmt.Println(q)

    for q.Size() > 0 {
        val, success := q.Dequeue()
        if success {
            fmt.Println(val)
        }
    }

}
*/













