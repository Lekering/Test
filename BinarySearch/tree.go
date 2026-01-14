package main

type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

type BST struct {
    Root *TreeNode
}

// Вставка
func (t *TreeNode) Insert(value int) *TreeNode {
    if t == nil {
        return &TreeNode{Value: value}
    }
    
    if value < t.Value {
        t.Left = t.Left.Insert(value)
    } else if value > t.Value {
        t.Right = t.Right.Insert(value)
    }
    return t
}

// Поиск
func (t *TreeNode) Search(value int) bool {
    if t == nil {
        return false
    }
    
    if value == t.Value {
        return true
    } else if value < t.Value {
        return t.Left.Search(value)
    }
    return t.Right.Search(value)
}

// Обход In-Order (левый-корень-правый)
func (t *TreeNode) InOrder(visit func(int)) {
    if t == nil {
        return
    }
    t.Left.InOrder(visit)
    visit(t.Value)
    t.Right.InOrder(visit)
}

// Обход Pre-Order (корень-левый-правый)
func (t *TreeNode) PreOrder(visit func(int)) {
    if t == nil {
        return
    }
    visit(t.Value)
    t.Left.PreOrder(visit)
    t.Right.PreOrder(visit)
}

// Обход Post-Order (левый-правый-корень)
func (t *TreeNode) PostOrder(visit func(int)) {
    if t == nil {
        return
    }
    t.Left.PostOrder(visit)
    t.Right.PostOrder(visit)
    visit(t.Value)
}

// Обход Level-Order (BFS)
func (t *TreeNode) LevelOrder(visit func(int)) {
    if t == nil {
        return
    }
    
    queue := []*TreeNode{t}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        visit(node.Value)
        
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
}