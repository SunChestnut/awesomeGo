package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var res ListNode
	p := &res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			p.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		p = p.Next
	}
	p.Next = list1
	if p.Next == nil {
		p.Next = list2
	}
	return res.Next
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	pre = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// packageListNode 组装链表
func packageListNode(nums []int) *ListNode {
	head := ListNode{}
	p := &head
	for _, v := range nums {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return head.Next
}

// printListNode 打印链表
func printListNode(head *ListNode) {
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

// detectCycle 检测环形链表
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// longestPalindrome 返回通过字母 s 可构造的最大回文串
func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, v := range []rune(s) {
		if m[v] > 0 {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}

	res := 0
	for _, v := range m {
		res += v / 2 * 2
		if res%2 == 0 && v%2 == 1 {
			res += 1
		}
	}
	return res
}

func main() {
	//node := packageListNode([]int{3, 2, 0, -4})
	//res := middleNode(node)
	//printListNode(res)

	fmt.Println(
		longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"),
		//longestPalindrome("abccccdd"),
		//longestPalindrome("abc"),
		//longestPalindrome("a"),
		//longestPalindrome("aaaaaa"),
	)

}
