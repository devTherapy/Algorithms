package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	curr := linkedList
	for curr.Next != nil {
		if curr.Value == curr.Next.Value {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}

	}
	return linkedList
}
