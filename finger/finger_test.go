package finger

import (
	"crypto/sha256"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetNonTrivialFingerTable() FingerTable {
	fingerTable := NewFingerTable()
	addr := "hash_val_3"
	hash := "3"
	entry := fingerEntry{Ip: addr, hashID: hash}
	fingerTable.insert(entry)

	addr = "hash_val_5"
	hash = "5"
	entry = fingerEntry{Ip: addr, hashID: hash}
	fingerTable.insert(entry)

	addr = "hash_val_1"
	hash = "1"
	entry = fingerEntry{Ip: addr, hashID: hash}
	fingerTable.insert(entry)

	addr = "hash_val_2"
	hash = "2"
	entry = fingerEntry{Ip: addr, hashID: hash}
	fingerTable.insert(entry)

	return fingerTable
}

func TestInsertHelperSingleValue(t *testing.T) {
	fingerTable := NewFingerTable()
	addr := "127.0.0.1"
	hash := sha256.Sum256([]byte(addr))
	hashStr := string(hash[:sha256.Size])
	entry := fingerEntry{Ip: "127.0.0.1", hashID: hashStr}
	fingerTable.insert(entry)
	assert.Equal(t, entry, *fingerTable.value)
}

func TestInsertHelperMultValues(t *testing.T) {
	fingerTable := GetNonTrivialFingerTable()

	assert.Equal(t, "3", fingerTable.value.hashID)
	assert.Equal(t, "5", fingerTable.right.value.hashID)
	assert.Equal(t, "1", fingerTable.left.value.hashID)
	assert.Equal(t, "2", fingerTable.left.right.value.hashID)
}

func TestClosestPredecessorSingleValue(t *testing.T) {
	fingerTable := NewFingerTable()
	fingerTable.Insert("127.0.0.1")
	succ := fingerTable.ClosestPredecessor("123")
	assert.Equal(t, "127.0.0.1", succ)

}

func TestClosestPredecessorMultValue(t *testing.T) {
	fingerTable := GetNonTrivialFingerTable()

	hash := "7"
	assert.Equal(t, "hash_val_5", fingerTable.ClosestPredecessor(hash))

	hash = "0"
	assert.Equal(t, "hash_val_5", fingerTable.ClosestPredecessor(hash))

	hash = "4"
	assert.Equal(t, "hash_val_3", fingerTable.ClosestPredecessor(hash))

	hash = "3"
	assert.Equal(t, "hash_val_3", fingerTable.ClosestPredecessor(hash))
}
