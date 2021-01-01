package renoder

import (
    "github.com/google/uuid"
)

type Renoder interface {
    InsertChild(Renoder)
    InsertSibling(Renoder)
    GetChild() Renoder
    GetSibling() Renoder
    Delete()
    GetParent() Renoder
}

type Node struct {
    child *Node
    parent *Node
    sibling *Node
    id uuid.UUID
}

func NewNode() Node {
    return Node{id: uuid.New()}
}

func (n *Node) GetChild() *Node {
    return n.child
}

func (n *Node) GetParent() *Node {
    return n.parent
}

func (n *Node) InsertChild(child *Node) {
    if n.child == nil {
        n.child = child
        child.parent = n
    } else {
        heldChild := n.child
        n.child = child
        child.parent = n
        child.child = heldChild
        heldChild.parent = child
    }
}

