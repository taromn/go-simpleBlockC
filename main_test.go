package main

import "testing"

func TestNewBlock(t *testing.T) {
	gen := Block{0, []string{"transaction1", "transaction2"}, "TestMerkleRoot", "TestTime", "TestHash", ""}
	b := newBlock(&gen, []string{"transaction1", "transaction2"})

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
