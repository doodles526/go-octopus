package finger

import (
	"crypto/sha256"
)

// Data is stored at successor ID
type fingerEntry struct {
	Ip     string
	hashID string
}

type FingerTable struct {
	value *fingerEntry
	left  *FingerTable
	right *FingerTable
}

func NewFingerTable() FingerTable {
	return FingerTable{}
}

func (f *FingerTable) Insert(ipAddr string) {
	hash := sha256.Sum256([]byte(ipAddr))
	hashStr := string(hash[:sha256.Size])
	entry := fingerEntry{Ip: ipAddr, hashID: hashStr}
	f.insert(entry)
}

// TODO: Get rid of nested conditional
func (f *FingerTable) insert(newNode fingerEntry) {
	if f.value == nil {
		f.value = &newNode
		return
	}

	if newNode.hashID < f.value.hashID {
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
func (f *FingerTable) ClosestPredecessor(hashID string) string {
	if pred := f.closestPredecessor(hashID, nil); pred != nil {
		return pred.Ip
	}

	return f.largestNode().value.Ip
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
func (f *FingerTable) closestPredecessor(hashID string, retEntry *fingerEntry) *fingerEntry {
	if f == nil {
		return retEntry
	}

	if hashID < f.value.hashID {
		return f.left.closestPredecessor(hashID, retEntry)
	} else if hashID == f.value.hashID {
		return f.value
	} else {
		return f.right.closestPredecessor(hashID, f.value)
	}

}
