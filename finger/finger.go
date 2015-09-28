package finger

import (
	"github.com/doodles526/go-octopus/valhash"
)

// Data is stored at successor ID
type fingerEntry struct {
	Ip     string
	hashID string
}

type FingerTable struct {
	value *valhash.ValHash
	left  *FingerTable
	right *FingerTable
}

func NewFingerTable() FingerTable {
	return FingerTable{}
}

func (f *FingerTable) Insert(ipAddr string) {
	entry := valhash.NewValHash([]byte(ipAddr))
	f.insert(entry)
}

// TODO: Get rid of nested conditional
func (f *FingerTable) insert(newNode *valhash.ValHash) {
	if f.value == nil {
		f.value = newNode
		return
	}

	if newNode.Compare(f.value) < 0 {
		if f.left == nil {
			f.left = &FingerTable{}
		}
		f.left.insert(newNode)
	} else {
		if f.right == nil {
			f.right = &FingerTable{}
		}
		f.right.insert(newNode)
	}
}

// ClosestSuccessor gets the closest node identified by the given hash
func (f *FingerTable) ClosestPredecessor(data []byte) string {
	valHash := valhash.NewValHash(data)
	if pred := f.closestPredecessor(valHash, nil); pred != nil {
		return string(pred.Value)
	}

	return string(f.largestNode().value.Value)
}

func (f *FingerTable) largestNode() *FingerTable {
	if f.right == nil {
		return f
	}
	return f.right.largestNode()
}

// closestGreaterThanEq returns the closest node which
// is greater than the given key
// if no node is greater then return nil
func (f *FingerTable) closestPredecessor(valHash *valhash.ValHash, retEntry *valhash.ValHash) *valhash.ValHash {
	if f == nil {
		return retEntry
	}

	if valHash.Compare(f.value) < 0 {
		return f.left.closestPredecessor(valHash, retEntry)
	} else if valHash.Compare(f.value) == 0 {
		return f.value
	} else {
		return f.right.closestPredecessor(valHash, f.value)
	}

}
