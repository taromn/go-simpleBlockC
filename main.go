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
	Tdata []string
	Root  string
	Time  string
	Hash  string
	Phash string
}

func calcHash(raw string) string {

	h := sha256.New()
	h.Write([]byte(raw))
	return hex.EncodeToString(h.Sum(nil))

}

func (b *Block) createH() {
	s := strconv.Itoa(b.Index) + b.Root + b.Time + b.Phash
	b.Hash = calcHash(s)
}

func newBlock(pre *Block, tdata []string) Block {
	var newb Block
	if pre == nil {
		newb = Block{0, tdata, "", time.Now().String(), "", ""} // genesis
	} else {
		newb = Block{pre.Index + 1, tdata, "", time.Now().String(), "", pre.Hash} // not genesis
	}
	newb.calcRoot()
	newb.createH()
	return newb
}

// calcRoot calculates the merkle root of transactions
func (b *Block) calcRoot() {
	var hashes []string
	for _, i := range b.Tdata {
		hashes = append(hashes, calcHash(i))
	}

	for len(hashes) > 1 {
		var temp []string
		for i := 0; i < len(hashes); i += 2 {
			if i+1 == len(hashes) {
				// if odd number of leaves, duplicate the last node
				// https://bitcoin.stackexchange.com/questions/79364/are-number-of-transactions-in-merkle-tree-always-even
				temp = append(temp, calcHash(hashes[i]+hashes[i]))

			} else {
				temp = append(temp, calcHash(hashes[i]+hashes[i+1]))

			}
		}
		hashes = temp
	}
	b.Root = hashes[0]

}

var tran1 = []string{"transaction1", "transaction2", "transaction3"}
var tran2 = []string{"transaction4", "transaction5"}

func main() {
	// if genesis block, the previous block should be nil
	genesis := newBlock(nil, tran1)
	fmt.Println(genesis)

	second := newBlock(&genesis, tran2)
	fmt.Println(second)

}
