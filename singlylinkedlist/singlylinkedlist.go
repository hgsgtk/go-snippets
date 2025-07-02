package singlylinkedlist

type MyLinkedList struct {
	head *Node
	tail *Node
	length int
}

type Node struct {
	Data int
	Next *Node
}


func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
		tail: nil,
		length: 0,
	}
}


func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.length {
		return -1
	}

	current := this.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Data
}


func (this *MyLinkedList) AddAtHead(val int)  {
	newNode := &Node{
		Data: val,
		Next: this.head,
	}
	this.head = newNode
	this.length++
	if this.tail == nil {
		this.tail = newNode
	}
}

func (this *MyLinkedList) AddAtTail(val int)  {
	newNode := &Node{
		Data: val,
		Next: nil,
	}
	if this.tail == nil {
		// List is empty
		this.head = newNode
		this.tail = newNode
	} else {
		this.tail.Next = newNode
		this.tail = newNode
	}
	this.length++
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index < 0 || index > this.length {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	current := this.head
	for i := 0; i < index - 1; i++ {
		current = current.Next
	}
	newNode := &Node{
		Data: val,
		Next: current.Next,
	}
	current.Next = newNode
	this.length++
	if index == this.length - 1 {
		this.tail = newNode
	}
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index >= this.length {
		return
	}

	if index == 0 {
		this.head = this.head.Next
		this.length--
		if this.length == 0 {
			this.tail = nil
		}
		return
	}

	current := this.head
	for i := 0; i < index - 1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	this.length--
	if current.Next == nil {
		this.tail = current
	} else if this.length == 0 {
		this.tail = nil
	}
}