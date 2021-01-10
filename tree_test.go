package renoder

import (
	"testing"
)

func TestChildNode(t *testing.T) {
	n1 := NewNode()
	n2 := NewNode()
	n3 := NewNode()

	t.Logf("some uuid: %v", n1.GetID())

	// Put n2 below n1
	n1.InsertChild(&n2)

	if n1.GetChild().GetID() != n2.GetID() {
		t.Errorf("n1 child id is not n2 id: %v, %v", n1.GetChild().GetID(), n2.GetID())
	}
	if n1.GetChild().GetParent().GetID() != n1.GetID() {
		t.Errorf("n2 parent is not n1! %v, %v", n1.GetChild().GetParent().GetID(), n1.GetID())
	}

	// insert n3 between n1 and n2
	n1.InsertChild(&n3)

	if n1.child.GetID() != n3.GetID() {
		t.Errorf("n1 doesn't have n3 as child")
	}
	if n3.child.GetID() != n2.GetID() {
		t.Errorf("n3 doesn't have n2 as child")
	}
	if n2.parent.GetID() != n3.GetID() {
		t.Errorf("n2 parent not udpated correctly")
	}
	if n2.child != nil {
		t.Errorf("n2 appears to have a child!")
	}

	// Find a child by walking from another node.
	n2p, ok := n1.Find(n2.GetID())
	if !ok {
		t.Errorf("unable to find n2 from n1")
	}
	if n2p.GetID() != n2.GetID() {
		t.Errorf("expecting uuid %v but got %v", n2.GetID(), n2p.GetID())
	}
}

func TestSiblingNode(t *testing.T) {
	n1 := NewNode()
	n2 := NewNode()
	n3 := NewNode()

	n1.InsertSibling(&n2)
	n2.InsertChild(&n3)

	// Find a child by walking from another node, through sibling.
	n3p, ok := n1.Find(n3.GetID())
	if !ok {
		t.Errorf("unable to find n3 from n1")
	}
	if n3p.GetID() != n3.GetID() {
		t.Errorf("expecting uuid %v but got %v", n3.GetID(), n3p.GetID())
	}
}
