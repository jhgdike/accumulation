package b_tree

type Node struct {
    val int
    left *Node
    right *Node
}

type Tree struct {
    root *Node
}

func NewTree() *Tree {
    return &Tree{&Node{}}
}

func (n *Node) Comp(val int) int {
    if n.val > val {
        return 1
    } else if n.val < val {
        return -1
    }
    return 0
}

func (t *Tree) Insert(val int) {
    if t.root == nil {
        t.root = &Node{val:val}
    } else {
        t.root.Insert(val)
    }
}

func (n *Node) Insert(val int) {
    comp := n.Comp(val)
    if comp == 1 {
        if n.left == nil {
            n.left = &Node{val:val}
        } else {
            n.left.Insert(val)
        }
    } else if comp == -1 {
        if n.right == nil {
            n.right = &Node{val:val}
        } else {
            n.right.Insert(val)
        }
    }
}

func (t *Tree) PreOrder() {
    if t.root == nil {
        return
    }
    t.root.PreOrder()
    print("\n")
}

func (n *Node) PreOrder() {
    print(n.val, " ")
    if n.left != nil {
        n.left.PreOrder()
    }
    if n.right != nil {
        n.right.PreOrder()
    }
}

func (t *Tree) InOrder() {
    if t.root == nil {
        return
    }
    t.root.InOrder()
    print("\n")
}

func (n *Node) InOrder() {
    if n.left != nil {
        n.left.InOrder()
    }
    print(n.val, " ")
    if n.right != nil {
        n.right.InOrder()
    }
}

func (t *Tree) PostOrder() {
    if t.root == nil {
        return
    }
    t.root.PostOrder()
    print("\n")
}

func (n *Node) PostOrder() {
    if n.left != nil {
        n.left.PostOrder()
    }
    if n.right != nil {
        n.right.PostOrder()
    }
    print(n.val, " ")
}
