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

func MiddleNode(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	fastPointer, slowPointer := linkedList, linkedList
	switcher := false
	for fastPointer.Next != nil {

		if switcher {
			slowPointer = slowPointer.Next
		}
		switcher = !switcher
		fastPointer = fastPointer.Next
	}
	return slowPointer
}
