package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	mergeNode := dummyHead
	for list1 != nil && list2 != nil {
		v1 := list1.Val
		v2 := list2.Val
		if v1 > v2 {
			tmpNode := &ListNode{
				Val:  v2,
				Next: nil,
			}
			// 将这个节点的值赋值给mergeNode
			mergeNode.Next = tmpNode
			mergeNode = tmpNode
			// 节点往下走
			list2 = list2.Next
		} else {
			tmpNode := &ListNode{
				Val:  v1,
				Next: nil,
			}
			// 将这个节点的值赋值给mergeNode
			mergeNode.Next = tmpNode
			mergeNode = tmpNode
			// 节点往下走
			list1 = list1.Next
		}
	}
	// 接下来是有剩余的，那么就直接加在mergeNode的后面
	if list1 != nil {
		mergeNode.Next = list1
	}
	if list2 != nil {
		mergeNode.Next = list2
	}
	return dummyHead.Next
}

func createListNode(n int) *ListNode {
	dummyHead := &ListNode{}
	head := dummyHead
	fmt.Println("please input val")
	for i := 0; i < n; i++ {
		var tmpVal int
		fmt.Scanln(&tmpVal)
		tmpNode := &ListNode{
			Val:  tmpVal,
			Next: nil,
		}
		head.Next = tmpNode
		head = tmpNode
	}
	return dummyHead.Next
}

func printNodeInformation(list *ListNode) {
	for list != nil {
		fmt.Printf("%v ", list.Val)
		list = list.Next
	}
	fmt.Println()
}

func main() {
	// list1 := createListNode(3)
	// list2 := createListNode(3)
	printNodeInformation(mergeTwoLists(nil, &ListNode{}))
}
