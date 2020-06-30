package main

type Tree struct {
    Val int
    Left *Tree
    Right *Tree
    Repeat int
    Parent *Tree
}

func BFS(tree *Tree) []int {
    que := []*Tree{}
    que = append(que, tree)
    res := []int{}
    return Helper(que, res)
}


func Helper(que []*Tree, res []int) []int{
    if len(que) == 0 {
        return res
    }
    res = append(res, que[0].Val)
    if que[0].Right != nil {
        que = append(que, que[0].Right)
    }
    if que[0].Left != nil {
        que = append(que, que[0].Left)
    }
    return Helper(que[1:], res)
}
