package list

import "errors"

var (
	outRangeError = errors.New("out of range")
)

type LinkList struct {
	headNode *Node
	len int
}

func (l *LinkList) IsEmpty() bool {
	return l.len == 0
}

func (l *LinkList) Length() int {
	return l.len
}

func (l *LinkList) GetHeadNode() *Node {
	return l.headNode
}

// 从头部增加
func (l *LinkList) Add(data interface{}) {
	node := &Node{Data:data}
	node.Next = l.headNode
	l.headNode = node
	l.len += 1
}

// 从尾部增加
func (l *LinkList) Append(data interface{}) {
	node := &Node{Data:data}
	if l.IsEmpty() {
		l.headNode = node
	} else {
		cur := l.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
	l.len += 1
}

func (l *LinkList) Insert(index int, data interface{}) {
	if index <= 0{
		l.Add(data)
	} else if index >= l.Length() {
		l.Append(data)
	} else {
		node := &Node{Data:data}
		cur := l.headNode
		for i := 0; i < index - 1; i ++ {
			cur = cur.Next
		}
		node.Next = cur.Next
		cur.Next = node
		l.len += 1
	}
}

func (l *LinkList)Find(data interface{}) *Node {
	cur := l.headNode
	for cur != nil {
		if cur.Data == data {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

func (l *LinkList)GetNodeAtIndex(index int) (node *Node, error error) {
	if index < 0 || index >= l.Length() {
		return nil, outRangeError
	} else {
		cur := l.headNode
		for i := 0; i < index; i ++ {
			cur = cur.Next
		}
		return cur, nil
	}
}

func (l *LinkList)Remove(data interface{}) {
	if l.IsEmpty() {
		return
	}
	if l.headNode.Data == data {
		l.headNode = l.headNode.Next
		l.len -= 1
	}
	for cur := l.headNode; cur.Next != nil; cur = cur.Next {
		if cur.Next.Data == data {
			cur.Next = cur.Next.Next
			l.len -= 1
			return
		}
	}
}

func (l *LinkList) RemoveAtIndex(index int) (node *Node, error error) {
	if index < 0 || index >= l.Length() {
		return nil, outRangeError
	} else if index == 0 {
		node = l.headNode
		l.headNode = l.headNode.Next
	} else {
		cur := l.headNode
		for i := 1; i < index; i ++ {
			cur = cur.Next
		}
		node = cur.Next
		cur.Next = cur.Next.Next
	}
	l.len -= 1
	return
}

func (l *LinkList) Contain(data interface{}) bool {
	return l.Find(data) != nil
}
