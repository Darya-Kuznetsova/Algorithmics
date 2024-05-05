package main

import (
	"container/list"
	"fmt"
)

func main() {

	// 1. DLL:
	ll_1 := list.New()
	ll_1.PushBack(8)
	ll_1.PushBack(11)
	ll_1.PushFront(7)
	ll_1.PushFront(13)

	fmt.Println("Print DLL: ")

	print(ll_1)

	// RETURN Front and Back elements:

	var_1 := ll_1.Back() // last node
	fmt.Println("Last DLL element: ", var_1.Value)

	var_2 := ll_1.Front()
	fmt.Println("First DLL element: ", var_2.Value)

	// LENGTH:

	fmt.Println("DLL length: ", ll_1.Len())

	// INIT. It clears (doesn't delete) a DLL:
	ll_2 := list.New()
	ll_2.PushBack(7)
	ll_2.PushBack(18)

	fmt.Println("Print DLL 2")
	print(ll_2)

	fmt.Println("Init: ")
	ll_2.Init()
	print(ll_2)
	fmt.Println(ll_2.Len())

	fmt.Println("Print reverse DLL: ")
	print_reverse(ll_1)

	// INSERT AFTER / BEFORE (use Front, Back, Prev, Next or search Node)
	fmt.Println("INSERT AFTER: ")
	ll_1.InsertAfter(19, ll_1.Front())
	ll_1.InsertBefore(6, ll_1.Back().Prev())
	print(ll_1)

	// PUSH BACK / FRONT LIST (add copy of other list):

	ll_2.PushBack(9)
	ll_2.PushBack(4)

	ll_1.PushFrontList(ll_2)
	fmt.Println("INSERT AFTER LIST: ")
	print(ll_1)

	var_3 := ll_1.Remove(ll_1.Front())
	fmt.Println("Removed element (value): ", var_3)
	print(ll_1)

}

// PRINT DDL:
func print(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

func print_reverse(l *list.List) {
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
