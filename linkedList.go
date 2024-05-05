package main

import "fmt"

/*
Linked list contains nodes. Node: value + next element link
Double linked list: value + next element link + previous element link.
Подходит, когда данные должны читаться последовательно (и требуется много вставок и удалений).
*/
func main() {
	list := Double_linked_list{}
	list.insert(5)
	list.insert(7)
	list.insert(4)
	fmt.Println("DLL:")
	list.print()

	list.reverse()
	fmt.Println("Reversed DLL:")
	list.print()

	fmt.Println("SEARCH:")
	var_1 := list.search(7)
	fmt.Println(var_1)

	fmt.Println("DELETE:")
	list.delete(var_1)
	list.print()

	fmt.Println("Insert after:")
	var_2 := list.search(4)
	list.insert_after(11, var_2)
	list.print()

}

// LL NODE:
type Node_1 struct {
	value int
	next  *Node_1
}

// DLL NODE:
type Node struct {
	value int // элемент списка
	// next element adress
	next *Node // next is a pointer
	// previous element adress (есть в двусвязном списке):
	previous *Node // previous is a pointer
}

// LINKED LIST:
type Linked_list struct {
	head *Node_1
}

// DOUBLE LINKED LIST:
type Double_linked_list struct {
	// начальный элемент списка:
	head *Node
	// последний элемент списка
	tail *Node
}

// 1. PRINT:

func (l *Double_linked_list) print() {
	current_node := l.head
	for current_node != nil {
		// fmt.Printf("%+v ->", current_node.value)
		fmt.Println(current_node.value)
		current_node = current_node.next
	}
}

// 2. SEARCHING:

func (l *Double_linked_list) search(value int) *Node {
	current_node := l.head
	for current_node != nil {
		if current_node.value == value {
			// fmt.Println(current_node)
			return current_node
		}
		current_node = current_node.next
	}
	return nil
}

// 3. INSERT (END):

func (l *Double_linked_list) insert(value int) {
	// Create a new node:
	new_node := &Node{}
	new_node.value = value
	if l.head == nil {
		l.head = new_node
		l.tail = new_node
	} else {
		l.tail.next = new_node
		new_node.previous = l.tail
		l.tail = new_node
	}
}

// 4. INSERT AFTER NODE:

func (l *Double_linked_list) insert_after(value int, node *Node) {
	next := node.next
	new_node := &Node{}
	new_node.value = value
	node.next = new_node
	new_node.previous = node
	if next == nil {
		l.tail = new_node
	} else {
		next.previous = new_node
		new_node.next = next
	}

}

// 5. DELETING:

func (l *Double_linked_list) delete(node *Node) {
	previous := node.previous
	next := node.next
	if previous == nil {
		next.previous = nil
		l.head = next
	}
	if next == nil {
		previous.next = nil
		l.tail = previous
	} else {
		previous.next = next
		next.previous = previous
	}

}

// 6. REVERSE (LL):

func (l *Linked_list) reverse_1(current, previous Node_1) {
	if current.next == nil {
		l.head = &current
	} else {
		l.reverse_1(*current.next, current)
	}
	current.next = &previous
	previous.next = nil
}

// 7. REVERSE (DLL):

func (l *Double_linked_list) reverse() {
	current_node := l.head
	for current_node != nil {
		next := current_node.next
		previous := current_node.previous
		current_node.next = previous
		current_node.previous = next
		if previous == nil {
			l.tail = current_node
		}
		if next == nil {
			l.head = current_node
		}
		current_node = next
	}
}

// STACK (LAST IN - FIRST OUT) and LINKED LIST:

func (l *Linked_list) push(value int) {
	new_node := &Node_1{}
	new_node.value = value
	new_node.next = l.head
	l.head = new_node

}

func (l *Linked_list) pop() int {
	result := -1
	if l.head != nil {
		result = l.head.value
		l.head = l.head.next
	}
	return result
}

// LINE (FIRST IN - FIRST OUT) and DOUBLE LINKED LIST:

func (l *Double_linked_list) push_line(value int) {
	new_node := &Node{}
	new_node.value = value
	new_node.next = l.head
	l.head.previous = new_node
	l.head = new_node
}
func (l *Double_linked_list) peek() int {
	result := -1
	if l.tail != nil {
		result = l.tail.value
		l.tail.previous.next = nil
		l.tail = l.tail.previous
	}
	return result

}
