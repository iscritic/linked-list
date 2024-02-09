package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var noda ListNode
	ListPushBack(&noda, 4)
	ShowNodeList(&noda)
}

func ListPushBack(head *ListNode, val int) *ListNode {

	newNode := &ListNode{Val: val, Next: nil}

	if head == nil {
		return newNode
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode

	return head
}

func ShowNodeList(head *ListNode) {

	if head == nil {
		fmt.Println("Список пуст")
		return
	}

	current := head

	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}
