package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// special case
	// 当节点个数小于2，不交换直接返回原节点
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{}
	cur := head
	newNode := dummyHead
	// 使用切片来模拟堆栈
	var stack []*ListNode
	// 还有两个节点
	for cur != nil && cur.Next != nil {
		// 将两个节点给放进堆栈中
		stack = append(stack, cur, cur.Next)
		// 让当前节点往前走两步
		cur = cur.Next.Next
		// 1. 取出堆栈中的数据
		// 2. 去除最后一个元素
		newNode.Next = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		newNode = newNode.Next
		newNode.Next = stack[len(stack)-1]
		newNode = newNode.Next
	}

	// 如果这个节链表的长度是奇数
	// 那么就把最后一个节点直接加在最后面
	if cur != nil {
		newNode.Next = cur
	} else {
		newNode.Next = nil
	}
	return dummyHead.Next
}

func createListNode(n int) *ListNode {
	dummyHead := &ListNode{}
	head := dummyHead
	fmt.Println("Create ListNode")
	for i := 0; i < n; i++ {
		var tmpVal int
		fmt.Scan(&tmpVal)
		tmpNode := &ListNode{
			Val:  tmpVal,
			Next: nil,
		}
		head.Next = tmpNode
		head = head.Next
	}
	return dummyHead.Next
}

func printListNodeInformation(node *ListNode) {
	for node != nil {
		fmt.Printf("%v ", node.Val)
		node = node.Next
	}
	fmt.Println()
}

func main() {
	// printListNodeInformation(createListNode(4))
	// printListNodeInformation(swapPairs(nil))
	printListNodeInformation(swapPairs(createListNode(4)))
}
