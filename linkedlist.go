package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type SinglyLinkedList struct {
	Head   *Node
	Length int
}

func (singlyList *SinglyLinkedList) Insert(data interface{}) {
	newNode := Node{}
	newNode.Data = data
	if singlyList.Length == 0 {
		singlyList.Head = &newNode
		singlyList.Length++
		return
	}
	temp := singlyList.Head
	for i := 0; i < singlyList.Length; i++ {
		if temp.Next == nil {
			temp.Next = &newNode
			singlyList.Length++
			return
		}
		temp = temp.Next
	}
}

func (singlyList *SinglyLinkedList) Display() {
	temp := singlyList.Head
	for i := 0; i < singlyList.Length; i++ {
		fmt.Println(temp.Data)
		temp = temp.Next
	}
}

func (singlyList *SinglyLinkedList) Search(val interface{}) bool {
	temp := singlyList.Head
	for i := 0; i < singlyList.Length; i++ {
		if val == temp.Data {
			return true
		}
		temp = temp.Next
	}
	return false
}

func (singlyList *SinglyLinkedList) InsertPosition(val interface{}, position int) {
	if singlyList.Length == 0 {
		singlyList.Insert(val)
		return
	} else if position == 0 {
		newNode := &Node{}
		newNode.Data = val
		newNode.Next = singlyList.Head
		singlyList.Head = newNode
		singlyList.Length++
		return
	}

	temp := singlyList.Head
	for i := 0; i < position-1; i++ {
		temp = temp.Next
	}
	fmt.Println("temp is ", temp.Data)
	newNode := &Node{}
	newNode.Data = val
	newNode.Next = temp.Next
	temp.Next = newNode
	singlyList.Length++
}

func (singlyList *SinglyLinkedList) GetNodeAt(position int) *Node {

	temp := singlyList.Head
	for i := 0; i < position; i++ {
		temp = temp.Next
	}
	fmt.Println("temp is ", temp.Data)
	return temp
}

func (singlyList *SinglyLinkedList) DeleteNodeByPosition(position int) {
	prev := singlyList.GetNodeAt(position - 1)
	after := singlyList.GetNodeAt(position)
	prev.Next = after.Next
	after.Next = nil
	singlyList.Length--
}

func main() {
	n := &SinglyLinkedList{}
	n.Insert(5)
	n.Display()
	fmt.Println("==========")
	n.Insert(6)
	n.Insert(7)
	n.Display()
	fmt.Println("==========")
	fmt.Println(n.Search(4))
	fmt.Println(n.Search(5))
	fmt.Println("==========")
	n.InsertPosition(8, 0)
	n.Display()
	fmt.Println("==========")
	nn := n.GetNodeAt(3)
	fmt.Println("Node is ", nn.Data)
	n.DeleteNodeByPosition(1)
	n.Display()
}
