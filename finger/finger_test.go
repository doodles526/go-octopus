package finger

import (
	"crypto/sha256"
	"github.com/doodles526/go-octopus/valhash"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetNonTrivialFingerTable() FingerTable {
	fingerTable := NewFingerTable()
	addr := []byte("hash_val_3")
	hash := []byte("3")
	entry := valhash.NewValHashWithHash(addr, hash)
	fingerTable.insert(entry)

	addr = []byte("hash_val_5")
	hash = []byte("5")
	entry = valhash.NewValHashWithHash(addr, hash)
	fingerTable.insert(entry)

	addr = []byte("hash_val_1")
	hash = []byte("1")
	entry = valhash.NewValHashWithHash(addr, hash)
	fingerTable.insert(entry)

	addr = []byte("hash_val_2")
	hash = []byte("2")
	entry = valhash.NewValHashWithHash(addr, hash)
	fingerTable.insert(entry)

	return fingerTable
}

func TestInsertHelperSingleValue(t *testing.T) {
	fingerTable := NewFingerTable()
	addr := []byte("127.0.0.1")
	hash := sha256.Sum256([]byte(addr))
	entry := valhash.NewValHashWithHash(addr, hash[:sha256.Size])
	fingerTable.insert(entry)
	assert.Equal(t, entry, fingerTable.value)
}

// Confirms the tree is constructed correctly
func TestInsertHelperMultValues(t *testing.T) {
	fingerTable := GetNonTrivialFingerTable()

	assert.Equal(t, []byte("hash_val_3"), fingerTable.value.Value)
	assert.Equal(t, []byte("hash_val_5"), fingerTable.right.value.Value)
	assert.Equal(t, []byte("hash_val_1"), fingerTable.left.value.Value)
	assert.Equal(t, []byte("hash_val_2"), fingerTable.left.right.value.Value)
}

func TestClosestPredecessorSingleValue(t *testing.T) {
	fingerTable := NewFingerTable()
	fingerTable.Insert("127.0.0.1")
	succ := fingerTable.ClosestPredecessor([]byte("123"))
	assert.Equal(t, "127.0.0.1", succ)

}
