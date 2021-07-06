package tstree

type Entry interface {
	Equals(Entry) bool
}

type Node struct {
	Left    *Node
	Right   *Node
	Entries []Entry
	limU    float32
	limL    float32
}

func (n *Node) Find(v Entry) Entry {
	for _, e := range n.Entries {
		if e.Equals(v) {
			return e
		}
	}

	return nil
}
