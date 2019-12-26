package list

type Node struct {
	Data interface{}
	Next *Node
}

type List interface {
	IsEmpty() bool
	Length() int
	GetHeadNNode() *Node
	Add(data interface{})  // 从头部增加
	Append(data interface{})  // 从尾部增加
	Insert(index int, data interface{})
	Find(data interface{}) *Node
	GetNodeAtIndex(index int) (node *Node, error error)
	Remove(data interface{})
	RemoveAtIndex(index int) (node *Node, error error)
	Contain(data interface{}) bool
}
