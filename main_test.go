package main

import "testing"

func TestGenesisBlock(t *testing.T) {
	got := newBlock(nil, []string{"transaction1", "transaction2"})
	// expected block ↓
	// Block{0, []string{"transaction1", "transaction2"}, "<calculated merkle root>", "<time>", "<calculated hash>", ""}
	if got.Index != 0 {
		t.Errorf("Incorrect genesis index, got %d", got.Index)
	}

	if got.Root == "" {
		t.Errorf("Merkle root in genesis block is not recorded")
	}

	if got.Time == "" {
		t.Errorf("Time in genesis block is not recorded")
	}

	if got.Hash == "" {
		t.Errorf("Hash in genesis block is not recorded")
	}

	if got.Phash != "" {
		t.Errorf("Genesis block shouldn't have the previous hash %s", got.Phash)
	}

}

func TestNewBlock(t *testing.T) {
	gen := newBlock(nil, []string{"transaction1", "transaction2"})
	b := newBlock(&gen, []string{"transaction3", "transaction4"})

	if b.Index != gen.Index+1 {
		t.Errorf("Expected index 1, got %d", b.Index)
	}

	if b.Root == "" {
		t.Errorf("Mercle root is not recorded")
	}

	if b.Time == "" {
		t.Errorf("Time is not recorded")
	}

	if b.Hash == "" {
		t.Errorf("Hash is not recorded")
	}

	if b.Phash != gen.Hash {
		t.Errorf("Hash doesn't match %s", b.Phash)
	}

}
