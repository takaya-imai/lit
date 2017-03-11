package portxo

import (
	"testing"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

var (
	inU1 = PorTxo{
		Op: wire.OutPoint{
			Hash:  chainhash.DoubleHashH([]byte("test")),
			Index: 3,
		},
		Value: 1234567890,
		Mode:  TxoP2PKHComp,
		Seq:   65535,
		KeyGen: KeyGen{
			Depth: 3,
			Step:  [5]uint32{0x8000002C, 1, 0x80000000},
		},
		PkScript: []byte("1234567890123456"),
	}

	inU2 = PorTxo{
		Op: wire.OutPoint{
			Hash:  chainhash.DoubleHashH([]byte("test2")),
			Index: 3,
		},
		Value: 5565989,
		Mode:  TxoP2WSHComp,
		Seq:   0,
		KeyGen: KeyGen{
			Depth: 1,
			Step:  [5]uint32{0x8000002C},
		},
		PkScript:    []byte("00112233"),
		PreSigStack: [][]byte{[]byte("SIGSTACK00000"), []byte(".....STACK001"), nil},
	}

	inUNil PorTxo
)

// TxoSliceByBip69
// this makes PorTxo slice sortable by Bip69
func TestLenInTxoSliceByBip69(t *testing.T) {
	// test a normal situation
	// input: TxoSliceByBip69 contains
	//   PorTxo: inU1
	//   PorTxo: inU2
	// want: 2
	var txoSlice TxoSliceByBip69 = []*PorTxo{&inU1, &inU2}
	if txoSlice.Len() != 2 {
		t.Fatalf("it needs to be 2")
	}

	// test a normal situation
	// input: TxoSliceByBip69 contains no elements
	// want: 0
	var txoSliceNil TxoSliceByBip69 = []*PorTxo{}
	if txoSliceNil.Len() != 0 {
		t.Fatalf("it needs to be 2")
	}

	// test an anomaly situation
	// input: TxoSliceByBip69 contains
	//   PorTxo: inU1
	//   PorTxo: inUNil
	// want: 2
	var txoSliceNilElm TxoSliceByBip69 = []*PorTxo{&inU1, &inUNil}
	if txoSliceNilElm.Len() != 2 {
		t.Fatalf("it needs to be 2")
	}
}

func TestSwapInTxoSliceByBip69(t *testing.T) {
	// test a normal situation
	// input: TxoSliceByBip69 contains the follows and swap them
	//   PorTxo: inU1
	//   PorTxo: inU2
	// want: first elm and second elm is swapped
	var txoSlice TxoSliceByBip69 = []*PorTxo{&inU1, &inU2}
	txoSlice.Swap(0, 1)
	if !txoSlice[0].Equal(&inU2) {
		t.Fatalf("it needs to be equal")
	}
	if !txoSlice[1].Equal(&inU1) {
		t.Fatalf("it needs to be equal")
	}

	// test an anomaly situation
	// input: TxoSliceByBip69 contains the follows and swap them
	//   PorTxo: inU1
	//   PorTxo: inUNil
	// want: first elm and second elm is swapped
	var txoSliceNilElm TxoSliceByBip69 = []*PorTxo{&inU1, &inUNil}
	txoSliceNilElm.Swap(0, 1)
	if !txoSliceNilElm[0].Equal(&inUNil) {
		t.Fatalf("it needs to be equal")
	}
	if !txoSliceNilElm[1].Equal(&inU1) {
		t.Fatalf("it needs to be equal")
	}

	// TODO: fix it
	// test an anomaly situation
	// input: TxoSliceByBip69 contains the follows and swap first with second
	//   PorTxo: inU1
	// want: no changes happen because there is no second element
	/*
		var txoSliceOneElm TxoSliceByBip69 = []*PorTxo{&inU1}
		txoSliceOneElm.Swap(0, 1)
		if !txoSliceNilElm[0].Equal(&inU1) {
			t.Fatalf("it needs to be equal")
		}
	*/
}

func TestLessInTxoSliceByBip69(t *testing.T) {
	// test a normal situation
	// input: TxoSliceByBip69 contains
	//   PorTxo: inU1
	//   PorTxo: inU2
	var txoSlice TxoSliceByBip69 = []*PorTxo{&inU1, &inU2}
	if txoSlice.Less(0, 1) != false {
		t.Fatalf("it needs to be equal")
	}
	if txoSlice.Less(1, 0) != true {
		t.Fatalf("it needs to be equal")
	}
	if txoSlice.Less(0, 0) != false {
		t.Fatalf("it needs to be equal")
	}

	// test an anomaly situation
	// input: TxoSliceByBip69 contains
	//   PorTxo: inU1
	//   PorTxo: inUNil
	var txoSliceNilElm TxoSliceByBip69 = []*PorTxo{&inU1, &inUNil}
	if txoSliceNilElm.Less(0, 1) != false {
		t.Fatalf("it needs to be equal")
	}
	if txoSliceNilElm.Less(1, 0) != true {
		t.Fatalf("it needs to be equal")
	}
	if txoSliceNilElm.Less(0, 0) != false {
		t.Fatalf("it needs to be equal")
	}

	// TODO: fix it
	// test an anomaly situation
	// input: TxoSliceByBip69 contains
	//   PorTxo: inU1
	// want: nil because it can not decide an order
	/*
		var txoSliceOneElm TxoSliceByBip69 = []*PorTxo{&inU1}
		if txoSliceOneElm.Less(0, 1) != nil {
			t.Fatalf("it needs to be equal")
		}
	*/
}

// TxoSliceByAmt
// this makes PorTxo slice sortable by amount
func TestLenInTxoSliceByAmt(t *testing.T) {
	// test a normal situation
	// input: TxoSliceByAmt contains
	//   PorTxo: inU1
	//   PorTxo: inU2
	// want: 2
	var txoSlice TxoSliceByAmt = []*PorTxo{&inU1, &inU2}
	if txoSlice.Len() != 2 {
		t.Fatalf("it needs to be 2")
	}

	// test a normal situation
	// input: TxoSliceByAmt contains no elements
	// want: 0
	var txoSliceNil TxoSliceByAmt = []*PorTxo{}
	if txoSliceNil.Len() != 0 {
		t.Fatalf("it needs to be 2")
	}

	// test an anomaly situation
	// input: TxoSliceByAmt contains
	//   PorTxo: inU1
	//   PorTxo: inUNil
	// want: 2
	var txoSliceNilElm TxoSliceByAmt = []*PorTxo{&inU1, &inUNil}
	if txoSliceNilElm.Len() != 2 {
		t.Fatalf("it needs to be 2")
	}
}

func TestSwapInTxoSliceByAmt(t *testing.T) {
	// test a normal situation
	// input: TxoSliceByAmt contains the follows and swap them
	//   PorTxo: inU1
	//   PorTxo: inU2
	// want: first elm and second elm is swapped
	var txoSlice TxoSliceByAmt = []*PorTxo{&inU1, &inU2}
	txoSlice.Swap(0, 1)
	if !txoSlice[0].Equal(&inU2) {
		t.Fatalf("it needs to be equal")
	}
	if !txoSlice[1].Equal(&inU1) {
		t.Fatalf("it needs to be equal")
	}

	// test an anomaly situation
	// input: TxoSliceByAmt contains the follows and swap them
	//   PorTxo: inU1
	//   PorTxo: inUNil
	// want: first elm and second elm is swapped
	var txoSliceNilElm TxoSliceByAmt = []*PorTxo{&inU1, &inUNil}
	txoSliceNilElm.Swap(0, 1)
	if !txoSliceNilElm[0].Equal(&inUNil) {
		t.Fatalf("it needs to be equal")
	}
	if !txoSliceNilElm[1].Equal(&inU1) {
		t.Fatalf("it needs to be equal")
	}

	// TODO: fix it
	// test an anomaly situation
	// input: TxoSliceByAmt contains the follows and swap first with second
	//   PorTxo: inU1
	// want: no changes happen because there is no second element
	/*
		var txoSliceOneElm TxoSliceByAmt = []*PorTxo{&inU1}
		txoSliceOneElm.Swap(0, 1)
		if !txoSliceNilElm[0].Equal(&inU1) {
			t.Fatalf("it needs to be equal")
		}
	*/
}

func TestLessInTxoSliceByAmt(t *testing.T) {
	// test a normal situation
	// input: TxoSliceByAmt contains
	//   PorTxo: inU1
	//   PorTxo: inU2
	var txoSlice TxoSliceByAmt = []*PorTxo{&inU1, &inU2}
	if txoSlice.Less(0, 1) != false {
		t.Fatalf("it needs to be equal")
	}
	if txoSlice.Less(1, 0) != true {
		t.Fatalf("it needs to be equal")
	}
	if txoSlice.Less(0, 0) != false {
		t.Fatalf("it needs to be equal")
	}

	// test an anomaly situation
	// input: TxoSliceByAmt contains
	//   PorTxo: inU1
	//   PorTxo: inUNil
	var txoSliceNilElm TxoSliceByAmt = []*PorTxo{&inU1, &inUNil}
	if txoSliceNilElm.Less(0, 1) != false {
		t.Fatalf("it needs to be equal")
	}
	if txoSliceNilElm.Less(1, 0) != true {
		t.Fatalf("it needs to be equal")
	}
	if txoSliceNilElm.Less(0, 0) != false {
		t.Fatalf("it needs to be equal")
	}

	// TODO: fix it
	// test an anomaly situation
	// input: TxoSliceByAmt contains
	//   PorTxo: inU1
	// want: nil because it can not decide an order
	/*
		var txoSliceOneElm TxoSliceByAmt = []*PorTxo{&inU1}
		if txoSliceOneElm.Less(0, 1) != nil {
			t.Fatalf("it needs to be equal")
		}
	*/
}

// KeyGenSortableSlice
// this makes KeyGen slice sortable
func TestLenInKeyGenSortableSlice(t *testing.T) {
}

func TestSwapInKeyGenSortableSlice(t *testing.T) {
}

func TestLessInKeyGenSortableSlice(t *testing.T) {
}