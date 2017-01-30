package lnutil

// Nothing to do unit-test for CommitScript and FundTxOut
// Because they depend on just other libraries and functions as follows.
// * github.com/btcsuite/btcd/txscript
// * github.com/btcsuite/btcd/wire
// * lnutil.P2WSHify
// It should test the other libraries and functions.

import (
	"testing"
)

// FundTxScript almost depends on github.com/btcsuite/btcd/txscript
// So it needs tests for just swapped value
//
// Define a test pubkey as testPub
// A testAPub is the same as testPub
// A testBPub is almost the same as testPub
// but increments just the last element of testBPub by one
// This tests three patterns, normal-order, inverse-order and the same
func TestSwappedInFundTxScript(t *testing.T) {
	testPub := [33]byte{
		0x37, 0x94, 0x38, 0x5f, 0x2a, 0x3e, 0xf7, 0xab,
		0x9d, 0x4a, 0xd3, 0xd1, 0xb3, 0xa3, 0x81, 0xb4,
		0x1f, 0x3f, 0xbf, 0xb5, 0x88, 0xa3, 0xef, 0xb9,
		0x0a, 0x59, 0x18, 0x83, 0x31, 0xb8, 0xb7, 0x53,
		0xd2,
	}

	// A testAPub is the same as testPub
	testAPub := testPub

	// A testBPub is almost the same as testPub
	// but increments just the last element of testBPub by one
	testBPub := testPub
	testBPub[len(testBPub)-1] = 0xd3

	_, swapped_true, _ := FundTxScript(testAPub, testBPub)
	if swapped_true != true {
		t.Fatalf("wrong swapped value")
	}

	_, swapped_false, _ := FundTxScript(testBPub, testAPub)
	if swapped_false != false {
		t.Fatalf("wrong swapped value")
	}

	_, swapped_same, _ := FundTxScript(testAPub, testAPub)
	if swapped_same != false {
		t.Fatalf("wrong swapped value")
	}
}
