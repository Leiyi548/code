package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// TODO: write a function to create a ListNode
func createListNode() *ListNode {
	pre := &ListNode{Val: 0}
	cur := pre
	var length int
	var newNodeVal int
	fmt.Println("Please input ListNode length")
	_, _ = fmt.Scanf("%d\n", &length)
	for i := 0; i < length; i++ {
		fmt.Sprintf("Please input node Val")
		_, _ = fmt.Scanf("%d\n", &newNodeVal)
		cur.Next = &ListNode{Val: newNodeVal}
		cur = cur.Next
	}
	return pre.Next
}

func printListNodeInformation(node *ListNode) {
	fmt.Println("printListNodeInformation")
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
	fmt.Println()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{Val: 0}
	cur := pre
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next

		if carry > 0 {
			cur.Next = &ListNode{Val: 1}
		}
	}
	return pre.Next
}

func main() {
	node1 := createListNode()
	node2 := createListNode()
	printListNodeInformation(node1)
	printListNodeInformation(node2)
	nodeSum := addTwoNumbers(node1, node2)
	printListNodeInformation(nodeSum)
}
