package renoder

import (
	"fmt"

	"github.com/google/uuid"
)

type Renoder interface {
	InsertChild(Renoder)
	InsertSibling(Renoder)
    GetID() uuid.UUID
	GetChild() Renoder
	GetSibling() Renoder
	//Delete()
	GetParent() Renoder
    SetChild(Renoder)
    SetParent(Renoder)
    SetSibling(Renoder)
	Find(id uuid.UUID) (Renoder, bool)
}

type Node struct {
	child   Renoder
	parent  Renoder
	sibling Renoder
	id      uuid.UUID
}

func NewNode() Node {
	return Node{id: uuid.New()}
}

func (n *Node) GetID() uuid.UUID {
    return n.id
}

func (n *Node) GetChild() Renoder {
	child := n.child
    if child != nil {
	    fmt.Printf("Child:   %v\n", child.GetID())
    }
	return child
}

func (n *Node) GetSibling() Renoder {
	sibling := n.sibling
    if sibling != nil {
	    fmt.Printf("Sibling: %v\n", sibling.GetID())
    }
	return sibling
}

func (n *Node) GetParent() Renoder {
	return n.parent
}

func (n *Node) SetParent(parent Renoder) {
    n.parent = parent
}

func (n *Node) SetChild(child Renoder) {
    n.child = child
}

func (n *Node) SetSibling(sibling Renoder) {
    n.sibling = sibling
}

func (n *Node) InsertChild(child Renoder) {
	if n.child == nil {
		n.SetChild(child)
		child.SetParent(n)
	} else {
		heldChild := n.child
		n.SetChild(child)
		child.SetParent(n)
		child.SetChild(heldChild)
		heldChild.SetParent(child)
	}
}

func (n *Node) InsertSibling(sibling Renoder) {
	if n.sibling == nil {
		n.SetSibling(sibling)
		sibling.SetParent(n)
	} else {
		heldSibling := n.sibling
		n.SetSibling(sibling)
		sibling.SetParent(n)
		sibling.SetSibling(heldSibling)
		heldSibling.SetParent(sibling)
	}
}

func (n *Node) Find(id uuid.UUID) (Renoder, bool) {
	if n.id == id {
		return n, true
	}
	child := n.GetChild()
	if child != nil {
		if found, ok := child.Find(id); ok {
			return found, ok
		}
	}
	sibling := n.GetSibling()
	if sibling != nil {
		if found, ok := sibling.Find(id); ok {
			return found, ok
		}
	}
	return nil, false
}
