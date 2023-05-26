package main

import "fmt"
/*
// define node stuct
type Node struct {
    data int
    next *Node
}

// define list struct
type SingleList struct {
    head *Node
    size int
}

// head
func (s *SingleList) GetHead() (int, error) {
    if s.head == nil {
        return -1, fmt.Errorf("List is empty")
    }

    return s.head.data, nil 
}

// tail
func (s *SingleList) GetTail() (*Node, error) {
    if s.head == nil {
        return nil , fmt.Errorf("List is empty")
    }

    tmp := s.head
    for tmp != nil {
        tmp = tmp.next
    }

    return tmp, nil 
}

// prepend
func (s *SingleList) Prepend(n *Node) {
    
    if s.head == nil {
        s.head = n 
    } else {
        n.next = s.head
        s.head = n 
    }
    s.size++
    return 
    
}

// append
func (s *SingleList) Append(n *Node) {
    if s.head == nil {
        s.head = n 
    } else {
        current := s.head
        for current.next != nil {
            current = current.next
        }
        current.next = n 
    }
    s.size++
    return 

}

// delete
func (s *SingleList) Delete(elem int) {
    if s.head == nil {
        fmt.Println("List is empty")
    }

    current := s.head
    for current != nil {
        if current.data == elem {
            current = nil
        }
        current = current.next
    }

    s.size--
    return 
}


func main() {
    l := SingleList{}

    node1 := &Node{data: 1} 
    node2 := &Node{data: 2} 
    node3 := &Node{data: 3} 
    node4 := &Node{data: 4} 

    l.Prepend(node1)
    l.Prepend(node2)
    l.Prepend(node3)
    l.Prepend(node4)
    fmt.Println(l)

    fmt.Println("Head node is:\n")
    fmt.Println(l.GetHead())
    fmt.Println("tail node is:\n")
    fmt.Println(l.GetTail())
}
     
*/

/*
================================================
*/
/*
type Node struct {
    data int
    next *Node
}

type SingleList struct {
    head *Node
    len int
}

// Initialize list
func initList() *SingleList {
    return &SingleList{}
}

// Prepend
func (s *SingleList) Prepend (data int) {
    node := &Node{
        data: data,
    } 

    if s.head == nil {
        s.head = node
    } else {
        secondNode := s.head
        s.head = node
        s.head.next = secondNode
    }

    s.len++
    return
}

// Append -> currently broken
func (s *SingleList) Append(data int) {
    node := &Node{
        data: data,
    }

    if s.head == nil {
        s.head = node
    } else {
        current := s.head
        for current.next != nil {
            current = current.next
        }
        current.next = node 
    }

    s.len++
    return 
}

// Delete Front
func (s *SingleList) DeleteFront() error {
    if s.head == nil {
        return fmt.Errorf("Linked List is empty")
    }

    s.head = s.head.next 
    s.len--
    return nil 
}

// Delete Back
func (s *SingleList) DeleteBack() error {
    if s.head == nil {
        return fmt.Errorf("Linked List is empty")
    }

    var prev *Node
    curr := s.head 
    for curr != nil {
        prev = curr
        curr = curr.next
    }
    if prev != nil {
        prev.next = nil
    } else {
        s.head = nil 
    }
    
    s.len--
    return nil 
}


// Get Head
func (s *SingleList) GetHead() int {
    if s.head == nil {
        return -1
    }

    return s.head.data 
}

// Get tail 
func (s *SingleList) GetTail() (int, error) {
    if s.head == nil {
        return -1, fmt.Errorf("Linked List is empty")
    }

    curr := s.head
    for curr.next != nil {
        curr = curr.next
    }

    return curr.data, nil 
}

// Print list
func (s *SingleList) PrintList() {
    if s.head == nil {
        fmt.Println("List is empty")
        return
    }

    curr := s.head
    for i := 0; i < s.len; i++ {
        fmt.Println(curr)
        curr = curr.next
    }
}

func main() {
    s := SingleList{}

    s.Prepend(1)
    s.Prepend(2)
    s.Prepend(3)
    s.PrintList()

    fmt.Println("Appending to end")
    s.Append(4)


    fmt.Println("Node is being deleted from front")
    s.DeleteFront()
    fmt.Println("New linked list is:")
    s.PrintList()

    fmt.Println("Node is being deleted from front")
    s.DeleteBack()
    fmt.Println("New linked list is")
    s.PrintList()

    fmt.Println("Head node is:")
    s.GetHead()

    fmt.Println("Tail node is: ")
    s.GetTail()

}

==============================================================
*/

type Node struct {
    data int
    next *Node
}

type SingleList struct {
    head *Node
    len int
}

// create list
func InitList() *SingleList {
    return &SingleList{}
}

// add front
func (s *SingleList) Prepend(data int) {
    node := &Node{
    data: data,
    }
    if s.head == nil {
        s.head = node
    } else {
        secondNode := s.head
        s.head = node
        s.head.next = secondNode
    }

    s.len++
    return
}

// add back
func (s *SingleList) Append(data int) {
    node := &Node{
        data: data,
    }

    if s.head == nil {
        s.head = nil
    } else {
        curr := s.head
        for curr != nil {
            curr = curr.next
        }
        curr.next = node
    }
    s.len++
    return
}

// remove front
func (s *SingleList) DeleteFront() error {
   if s.head == nil {
        return fmt.Errorf("List is empty")
   } 

   s.head = s.head.next
   s.len--
   return nil
}

// remove back
func (s *SingleList) DeleteBack() error {
    if s.head == nil {
        return fmt.Errorf("List is empty")
    }

    var prev *Node
    curr := s.head
    for curr != nil {
        prev = curr
        curr = curr.next
    } 
    if prev != nil {
        prev.next = nil
    } else {
        s.head = nil
    }

    s.len--
    return nil
}

// get head
func (s *SingleList) GetHead() (int, error) {
    if s.head == nil {
        return -1, fmt.Errorf("List is empty")
    }

    return s.head.data, nil
}

// get tail
func (s *SingleList) GetTail() (int, error) {
    if s.head == nil {
        return -1, fmt.Errorf("List is empty")
    }

    curr := s.head
    for curr.next != nil {
        curr = curr.next
    }

    return curr.data, nil 
}


func main() {
    s := InitList()

    s.Append(1)
    s.Append(2)
    s.Append(3)
    s.Append(4)
    s.Append(5)

    fmt.Println(s)

    s.Prepend(6)
    s.Prepend(7)
    s.Prepend(8)
    s.Prepend(9)

    fmt.Println(s)

    s.DeleteFront()

    s.DeleteBack()

    fmt.Println(s)

    s.GetHead()

    s.GetTail()

    fmt.Println(s)
}







