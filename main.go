package main

import (
	"fmt"
	"strconv"
)

func main(){
	blockchain := NewBlockchain()
	blockchain.AddBlock("New first block")
	blockchain.AddBlock("New second block")
	
	for _, block := range blockchain.blocklist{
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		proof := ProofOfWorkfunc(block)
		fmt.Printf("Proof of work: %s\n", strconv.FormatBool(proof.Validatepow()))
		fmt.Println()
	}
}