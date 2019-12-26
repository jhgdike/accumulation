package list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list := &LinkList{}

	// 判断是否为空
	fmt.Println("是否为空：", list.IsEmpty())
	// 添加头部
	list.Add("头部")
	// 添加尾部
	list.Append("尾部")
	// 获取长度
	fmt.Println(list.Length())
	// 获取头节点
	fmt.Println(list.GetHeadNode())
	// 判断是否为空
	fmt.Println("是否为空：", list.IsEmpty())
	// 从第二个位置插入
	list.Insert(2, "第二个位置")
	// 获取第二个位置
	data, _ := list.GetNodeAtIndex(2)
	fmt.Println(data.Data)
	// 删除第100个元素
	fmt.Println(list.RemoveAtIndex(100))
}
