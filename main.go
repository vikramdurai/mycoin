package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

type Block struct {
	Hash      []byte
	Data      string
	Parent    []byte
	Timestamp int64
}

func NewBlock(data string, parent []byte) *Block {
	h := sha256.New()
	h.Write([]byte(data))
	return &Block{h.Sum(nil), data, parent, time.Now().Unix()}
}

type Blockchain struct {
	Blocks []*Block
	sync.Mutex
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	bc.Blocks = append(bc.Blocks, NewBlock("Genesis Block", nil))
	return bc
}

func (bc *Blockchain) AddBlock(data string) {
	b := NewBlock(data, bc.LastBlock().Hash)
	bc.Lock()
	bc.Blocks = append(bc.Blocks, b)
	bc.Unlock()
}

func (bc *Blockchain) LastBlock() *Block {
	bc.Lock()
	defer bc.Unlock()
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) String() string {
	s := "["
	for _, v := range bc.Blocks {
		if v == bc.LastBlock() {
			s += fmt.Sprintf("\"%s\"]", v.Data)
		} else {
			s += fmt.Sprintf("\"%s\", ", v.Data)
		}
	}
	return s
}

func (bc *Blockchain) Save() {
	b, err := json.MarshalIndent(bc, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("blockchain.dat", b, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	chain := NewBlockchain()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("Blockchain =>", chain.String())
	go (func(chain *Blockchain) {
		chain.AddBlock("First block since genesis")
		wg.Done()
	})(chain)
	chain.AddBlock("Latest and greatest block")
	fmt.Println("Blockchain =>", chain.String())
	wg.Wait()
	fmt.Println("Blockchain =>", chain.String())
	chain.Save()
}
