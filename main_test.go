package main

import (
  "crypto/sha256"
	"time"
  "testing"
)

func equals(b1, b2 *Block) bool {
  equal := true
  if b1.Data != b2.Data {
    equal = false
  }
  if string(b1.Hash) != string(b2.Hash) {
    equal = false
  }
  if string(b1.Parent) != string(b2.Parent) {
    equal = false
  }
  if b1.Timestamp != b2.Timestamp {
    equal = false
  }
  return equal
}

func TestNewBlock(t *testing.T) {
  parent1 := sha256.New()
  parent1.Write([]byte("parent number 1"))
  parent2 := sha256.New()
  parent2.Write([]byte("parent number 2"))
  parents := [][]byte{
    parent1.Sum(nil),
    parent2.Sum(nil),
  }
  hash1 := sha256.New()
  hash1.Write([]byte("rainbows"))
  hash2 := sha256.New()
  hash2.Write([]byte("jelly"))
  test_blocks := []*Block{
    &Block{hash1.Sum(nil), "rainbows", parents[0], time.Now().Unix()},
    &Block{hash2.Sum(nil), "jelly", parents[1], time.Now().Unix()},
  }
  for _, c := range test_blocks {
    b := NewBlock(c.Data, c.Parent)
    b.Timestamp = c.Timestamp
    if equals(b, c) != true {
      t.Error("New Block is not as expected")
      t.Fail()
    }
  }
}
