package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index int
	Data  string
	Time  string
	Hash  []byte
	Phash []byte
}

func (b *Block) createH() {
	s := strconv.Itoa(b.Index) + b.Data + b.Time + hex.EncodeToString(b.Phash)
	h := sha256.New()
	h.Write([]byte(s))
	b.Hash = h.Sum(nil)
}

func newBlock(pre *Block, d string) Block {

	newb := Block{pre.Index + 1, d, time.Now().String(), []byte{}, pre.Hash}
	newb.createH()
	return newb
}

func main() {
	genesis := Block{0, "First Block", time.Now().String(), []byte{}, []byte{}}

	genesis.createH()

	fmt.Println(genesis)

	second := newBlock(&genesis, "Second Block")

	fmt.Println(second)

}
