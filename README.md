simple Blockchain model

## Block generation process of this repo

### step1/ creates a Block instance

```
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
```

newBlock() generates a new block, and calls `calcRoot()`, `createH()` methods to complete step2, 3.


### step2/ calculates the merkle root of transactions

`calcRoot()` calculate the merkle root of transactions

if odd number of leaves, duplicates the last node

Ref: https://bitcoin.stackexchange.com/questions/79364/are-number-of-transactions-in-merkle-tree-always-even


### step3/ creates a hash of the block

```
func (b *Block) createH() {
	s := strconv.Itoa(b.Index) + b.Root + b.Time + b.Phash
	b.Hash = calcHash(s)
}
```
Merkle root is included in the block hush calculation.
