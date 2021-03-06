package main

import (
	"fmt"
	"strconv"
	"github.com/fatih/color"
	"github.com/mit-dci/lit/lnutil"
	"github.com/mit-dci/lit/litrpc"
)

// Send sends coins somewhere
func (lc *litAfClient) Send(textArgs []string) error {
	if len(textArgs) > 0 && textArgs[0] == "-h" {
		fmt.Fprintf(color.Output,"%s%s\n", lnutil.White("send"), lnutil.ReqColor("addr", "amount"))
		fmt.Fprintf(color.Output,"Send the given amount of satoshis to the given address.\n")
		return nil
	}

	args := new(litrpc.SendArgs)
	reply := new(litrpc.TxidsReply)

	// need args, fail
	if len(textArgs) < 2 {
		return fmt.Errorf("need args: ssend address amount(satoshis) wit?")
	}
	/*
		adr, err := btcutil.DecodeAddress(args[0], lc.Param)
		if err != nil {
			fmt.Fprintf(color.Output,"error parsing %s as address\t", args[0])
			return err
		}
	*/
	amt, err := strconv.Atoi(textArgs[1])
	if err != nil {
		return err
	}

	fmt.Fprintf(color.Output,"send %d to address: %s \n", amt, textArgs[0])

	args.DestAddrs = []string{textArgs[0]}
	args.Amts = []int64{int64(amt)}

	err = lc.rpccon.Call("LitRPC.Send", args, reply)
	if err != nil {
		return err
	}
	fmt.Fprintf(color.Output,"sent txid(s):\n")
	for i, t := range reply.Txids {
		fmt.Fprintf(color.Output,"\t%d %s\n", i, t)
	}
	return nil
}

// Sweep moves utxos with many 1-in-1-out txs
func (lc *litAfClient) Sweep(textArgs []string) error {
	if len(textArgs) > 0 && textArgs[0] == "-h" {
		fmt.Fprintf(color.Output,"%s%s%s\n", lnutil.White("sweep"), lnutil.ReqColor("addr", "howmany"), lnutil.OptColor("drop"))
		fmt.Fprintf(color.Output,"Move UTXOs with many 1-in-1-out txs.\n")
		// TODO: Make this more clear.
		return nil
	}

	args := new(litrpc.SweepArgs)
	reply := new(litrpc.TxidsReply)

	var err error

	if len(textArgs) < 2 {
		return fmt.Errorf("sweep syntax: sweep adr howmany (drop)")
	}

	args.DestAdr = textArgs[0]
	args.NumTx, err = strconv.Atoi(textArgs[1])
	if err != nil {
		return err
	}

	if len(textArgs) > 2 {
		args.Drop = true
	}

	err = lc.rpccon.Call("LitRPC.Sweep", args, reply)
	if err != nil {
		return err
	}
	fmt.Fprintf(color.Output,"Swept\n")
	for i, t := range reply.Txids {
		fmt.Fprintf(color.Output,"%d %s\n", i, t)
	}

	return nil
}

//// ------------------------- fanout
//type FanArgs struct {
//	DestAdr      string
//	NumOutputs   uint32
//	AmtPerOutput int64
//}

func (lc *litAfClient) Fan(textArgs []string) error {
	if len(textArgs) > 0 && textArgs[0] == "-h" {
		fmt.Fprintf(color.Output,"%s%s\n", lnutil.White("fan"), lnutil.ReqColor("addr", "howmany", "howmuch"))
		// TODO: Add description.
		return nil
	}

	args := new(litrpc.FanArgs)
	reply := new(litrpc.TxidsReply)
	if len(textArgs) < 3 {
		return fmt.Errorf("fan syntax: fan adr howmany howmuch")
	}
	var err error
	args.DestAdr = textArgs[0]

	outputs, err := strconv.Atoi(textArgs[1])
	if err != nil {
		return err
	}
	args.NumOutputs = uint32(outputs)

	amt, err := strconv.Atoi(textArgs[2])
	if err != nil {
		return err
	}
	args.AmtPerOutput = int64(amt)

	err = lc.rpccon.Call("LitRPC.Fanout", args, reply)
	if err != nil {
		return err
	}

	fmt.Fprintf(color.Output,"Fanout:\n")
	for i, t := range reply.Txids {
		fmt.Fprintf(color.Output,"\t%d %s\n", i, t)
	}
	return nil
}

// Adr makes new addresses
func (lc *litAfClient) Adr(textArgs []string) error {
	if len(textArgs) > 0 && textArgs[0] == "-h" {
		fmt.Fprintf(color.Output,lnutil.White("adr\n"))
		fmt.Fprintf(color.Output,"Makes a new address.\n")
		// TODO: Make this more clear.
		return nil
	}

	args := new(litrpc.AdrArgs)
	args.NumToMake = 1
	reply := new(litrpc.AdrReply)
	err := lc.rpccon.Call("LitRPC.Address", args, reply)
	if err != nil {
		return err
	}
	fmt.Fprintf(color.Output,"new adr(s): %s\nold: %s\n", lnutil.Address(reply.WitAddresses), lnutil.Address(reply.LegacyAddresses))
	return nil
}
