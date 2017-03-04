package portxo

import (
	"testing"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

// TxoSliceByBip69
// this makes PorTxo slice sortable by Bip69
func TestLenInTxoSliceByBip69(t *testing.T) {
	var u1 PorTxo
	u1.Op.Hash = chainhash.DoubleHashH([]byte("test"))
	u1.Op.Index = 3
	u1.Value = 1234567890
	u1.Mode = TxoP2PKHComp
	u1.Seq = 65535
	u1.KeyGen.Depth = 3
	u1.KeyGen.Step[0] = 0x8000002C
	u1.KeyGen.Step[1] = 1
	u1.KeyGen.Step[2] = 0x80000000
	u1.PkScript = []byte("1234567890123456")

	var u2 PorTxo
	u2.Op.Hash = chainhash.DoubleHashH([]byte("test"))
	u2.Op.Index = 3
	u2.Value = 1234567890
	u2.Mode = TxoP2PKHComp
	u2.Seq = 65535
	u2.KeyGen.Depth = 3
	u2.KeyGen.Step[0] = 0x8000002C
	u2.KeyGen.Step[1] = 1
	u2.KeyGen.Step[2] = 0x80000000
	u2.PkScript = []byte("1234567890123456")

	var txoSlice TxoSliceByBip69 = []*PorTxo{&u1, &u2}

	if txoSlice.Len() != 2 {
		t.Fatalf("it needs to be 2")
	}
}

func TestSwapInTxoSliceByBip69(t *testing.T) {
}

func TestLessInTxoSliceByBip69(t *testing.T) {
}

// TxoSliceByAmt
// this makes PorTxo slice sortable by amount
func TestLenInTxoSliceByAmt(t *testing.T) {
}

func TestSwapInTxoSliceByAmt(t *testing.T) {
}

func TestLessInTxoSliceByAmt(t *testing.T) {
}

// KeyGenSortableSlice
// this makes KeyGen slice sortable
func TestLenInKeyGenSortableSlice(t *testing.T) {
}

func TestSwapInKeyGenSortableSlice(t *testing.T) {
}

func TestLessInKeyGenSortableSlice(t *testing.T) {
}
