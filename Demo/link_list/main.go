package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//l1 := newList([]int{8, 9, 9}) // 5    // 9 8
	//l2 := newList([]int{2})       // 5    // 1
	//cur := addTwoNumbers(l1, l2)
	//for cur != nil {
	//	fmt.Print(cur.Val, " ")
	//	cur = cur.Next
	//}
	//fmt.Println(removeNthFromEnd(newList([]int{1, 3, 4, 10, 15}), 2))
	//cur := mergeTwoLists2(newList([]int{1, 3, 4, 10}), newList([]int{2, 5, 6, 7, 8}))
	node := iterationReverse(newList2([]int{1, 2, 3, 4, 5, 6}))
	//cur := swapPairs(newList([]int{1, 2, 3, 4, 5, 6}))
	//cur := rotateRight(node, 2)
	fmt.Println(node)
}

func iterationReverse(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	} else {
		var beg *ListNode
		mid := node
		end := node.Next
		for {
			mid.Next = beg
			if end == nil {
				break
			}
			beg = mid
			mid = end
			end = end.Next
		}
		node = mid
		return node
	}
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := count(head)
	step := k % n
	for step > 0 {
		pre, tail := tailNodes(head) // 获取到倒数二个节点：pre->tail->nil
		pre.Next = nil               // 删除 tail
		tail.Next = head             // 将 tail 作为新的 head
		head = tail
		fmt.Println("head = ", head)
		step--
	}
	return head
}

func tailNodes(head *ListNode) (*ListNode, *ListNode) {
	pre, tail := head, head.Next
	for tail.Next != nil {
		pre = tail
		tail = tail.Next
	}
	return pre, tail
}

func count(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}

// 处理好 2 种异常 case
// 将两两节点视为内部短链表，交换两个节点后返回新的头节点（原第二节点），接到尾巴上即可
// 一般 case：pre->cur->next 经 swap() 调换后 next->pre，手动接上 pre->next->cur
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // 异常case: [] 或 [1]
		return head
	}
	newHead := swap(head, head.Next)
	pre := head
	if pre.Next == nil || pre.Next.Next == nil { // 异常case: [1 ,2] 或 [1 ,2, 3]
		return newHead
	}
	cur := pre.Next
	next := cur.Next
	for cur != nil && next != nil { // 一般case: 两两结对向后遍历
		pre.Next = swap(cur, next) // 内部交换
		if cur.Next == nil {
			break
		}
		pre = cur
		cur = cur.Next
		next = cur.Next
	}
	return newHead
}

func swap(cur, next *ListNode) *ListNode {
	cur.Next = next.Next
	next.Next = cur
	return next
}

func mergeTwoLists2(l1, l2 *ListNode) *ListNode {
	cur := new(ListNode)
	dummy := cur
	cur1, cur2 := l1, l2
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			cur.Next = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur2 = cur2.Next
		}
		cur = cur.Next
	}
	traverse := func(c, l *ListNode) {
		if l != nil {
			c.Next = l
			cur = cur.Next
			l = l.Next
		}
	}
	traverse(cur, cur1)
	traverse(cur, cur2)
	return dummy.Next
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	var nums []int
	cur1, cur2 := l1, l2
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			nums = append(nums, cur1.Val)
			cur1 = cur1.Next
		} else {
			nums = append(nums, cur2.Val)
			cur2 = cur2.Next
		}
	}
	traverse := func(nums []int, l *ListNode) []int {
		if l != nil {
			nums = append(nums, l.Val)
			l = l.Next
		}
		return nums
	}
	nums = traverse(nums, cur1)
	nums = traverse(nums, cur2)
	return newList(nums)
}

// 和倒数相关的问题考虑天生有间距的"双指针"，一次遍历解决
// 借助哑节点是关键，用于防止删除倒数第 len(nums) 个节点会导致链表丢失
// 若要兼顾正常节点的移动，又要处理头结点的特殊情况，请考虑"哑节点"
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Val: 0, Next: head}
	front, rear := dummyNode, dummyNode
	for counts := 0; counts <= n; counts++ {
		rear = rear.Next
	}
	for rear != nil {
		front = front.Next
		rear = rear.Next
	}
	front.Next = front.Next.Next // 删除节点
	return dummyNode.Next
}

func addTwoNumbers(node1, node2 *ListNode) *ListNode {
	var nums []int
	cur1, cur2 := node1, node2
	var carryBit bool // 是否要进位
	for cur1 != nil && cur2 != nil {
		sum := cur1.Val + cur2.Val
		if carryBit {
			sum += 1
		}
		carryBit = false
		if sum > 0 {
			carryBit = true
		}
		nums = append(nums, sum%10)
		cur1, cur2 = node1.Next, node2.Next
	}
	nums = append(nums, traverse(cur1, carryBit)...)
	nums = append(nums, traverse(cur2, carryBit)...)
	if cur1 == nil && cur2 == nil && carryBit {
		nums = append(nums, 1)
	}
	return newList(nums)
}

func traverse(cur *ListNode, carryBit bool) (remainNums []int) {
	if cur == nil {
		return
	}
	for cur != nil {
		if carryBit {
			res := cur.Val + 1
			if res >= 10 {
				remainNums = append(remainNums, 0)
				carryBit = true
			} else {
				remainNums = append(remainNums, res)
				carryBit = false
			}
			cur = cur.Next
			continue
		}
		remainNums = append(remainNums, cur.Val)
		cur = cur.Next
	}
	if carryBit {
		remainNums = append(remainNums, 1)
	}
	return
}

//正序构建
func newList(nums []int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0], Next: nil}
	cur := head
	for i := 1; i < n; i++ {
		newNode := &ListNode{Val: nums[i], Next: nil}
		cur.Next = newNode
		cur = newNode
	}
	return head
}

//逆序构建
func newList2(nums []int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	var tail *ListNode
	for i := n - 1; i >= 0; i-- {
		node := &ListNode{Val: nums[i], Next: tail}
		tail = node
	}
	return tail
}

func (cur *ListNode) String() string {
	counts := 0
	var nums []int
	for cur != nil {
		nums = append(nums, cur.Val)
		counts++
		cur = cur.Next
	}
	return fmt.Sprintf("%d nodes: %v", counts, nums)
}
