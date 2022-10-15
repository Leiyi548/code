package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd01(head *ListNode, n int) *ListNode {
	length := getListNodeLength(head)
	dummyHead := &ListNode{}
	dummyHead.Next = head
	if length == 1 {
		return nil
		// 如果删除的是头节点
	} else if n == length {
		dummyHead.Next = head.Next
		// 如果删除的是末尾节点
	} else if n == 1 {
		for i := 0; i < length-2; i++ {
			head = head.Next
		}
		head.Next = nil
		// 如果删除的节点在中间
		// 用两个节点来进行遍历
	} else {
		prev := &ListNode{}
		for i := 0; i < length-n; i++ {
			prev = head
			head = head.Next
		}
		prev.Next = head.Next
	}

	return dummyHead.Next
}

func getListNodeLength(node *ListNode) (length int) {
	for node != nil {
		length++
		node = node.Next
	}
	return
}

func createListNode(n int) *ListNode {
	prev := &ListNode{}
	next := prev
	fmt.Println("please input val")
	for i := 0; i < n; i++ {
		var tmpVal int
		fmt.Scanln(&tmpVal)
		tmpNode := &ListNode{
			Val:  tmpVal,
			Next: nil,
		}
		next.Next = tmpNode
		next = tmpNode
	}
	return prev.Next
}

func printListNodeInformation(node *ListNode) {
	fmt.Println("===printListNodeInformation===")
	for node != nil {
		fmt.Printf("%v ", node.Val)
		node = node.Next
	}
	fmt.Println()
}

func removeNthFromEnd02(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	// 设置快慢指针 slow fast
	slow := dummyHead
	fast := dummyHead
	// 让快指针先走n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 删除那个节点
	delNode := slow.Next
	slow.Next = delNode.Next
	delNode.Next = nil
	return dummyHead.Next
}

func main() {
	var n int
	fmt.Println("plase input n")
	fmt.Scanln(&n)
	node := createListNode(n)
	printListNodeInformation(removeNthFromEnd02(node, 1))
}
