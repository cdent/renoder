package renoder

import (
    "testing"
)

func TestNode(t *testing.T) {
    n1 := NewNode()
    n2:= NewNode()
    n3:= NewNode()

    t.Logf("some uuid: %v", n1.id)

    // Put n2 below n1
    n1.InsertChild(&n2)

    if n1.GetChild().id != n2.id {
        t.Errorf("n1 child id is not n2 id: %v, %v", n1.GetChild().id, n2.id)
    }
    if n1.GetChild().GetParent().id != n1.id {
        t.Errorf("n2 parent is not n1! %v, %v", n1.GetChild().GetParent().id, n1.id)
    }

    // insert n3 between n1 and n2
    n1.InsertChild(&n3)

    if n1.child.id != n3.id {
        t.Errorf("n1 doesn't have n3 as child")
    }
    if n3.child.id != n2.id {
        t.Errorf("n3 doesn't have n2 as child")
    }
    if n2.parent.id != n3.id {
        t.Errorf("n2 parent not udpated correctly")
    }
    if n2.child != nil {
        t.Errorf("n2 appears to have a child!")
    }

}
