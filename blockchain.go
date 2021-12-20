package main

type Blockchain struct {
	// blocks are continual
	blocklist []*Block
}

// stores passed data(string) within a newBlock in the blockchain
func (blockchain *Blockchain) AddBlock(data string){
	prevBlock := blockchain.blocklist[len(blockchain.blocklist)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	blockchain.blocklist = append(blockchain.blocklist, newBlock)
}

//create new Blockchain with genesis block 
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenBlock()}}
}