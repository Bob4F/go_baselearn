package main

import "fmt"

func main() {
	l1 :=ListNode{2,&ListNode{4,&ListNode{3,nil}}}
	l2 :=ListNode{5,&ListNode{6,&ListNode{4,nil}}}
	result :=twonumcount(&l1,&l2)
	for result!=nil  {
		fmt.Print(result.Val,"->")
		result=result.Next
	}
}


func twonumcount(l1 *ListNode,l2 *ListNode)*ListNode{
	if l1 ==nil {
		return l2
	}else if l2== nil{
		return l1
	}
	var result = &ListNode{0,nil}
	dump := result
	var carry = 0
	for l1 != nil || l2!=nil {
			var sum int
			if l1==nil {
				sum =  l2.Val + carry
			}else if l2==nil{
				sum = l1.Val +  carry
			}else {
				sum = l1.Val + l2.Val + carry
			}

			fmt.Println("init sum:",sum)
			carry = sum / 10
			fmt.Println("carry:",carry)
			sum %= 10
			fmt.Println("sum&:",sum)

			dump.Next = &ListNode{sum,nil}

			dump =dump.Next
			if l1 != nil{
				l1 = l1.Next
			}
			if(l2 != nil){
				l2 = l2.Next
			}
	}
	if carry==1 {
		dump.Next = &ListNode{1,nil}
	}
	return result.Next;
}



type ListNode struct {
	Val int
	Next *ListNode
}