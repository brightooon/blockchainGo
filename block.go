package main

import(
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//init Block object
type Block struct {
	Data 			[]byte
	PrevBlockHash	[]byte
	Hash			[]byte
	Timestamp		int64
}

// set block hash 
func (block *Block) SetHash(){
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	block.Hash = hash[:]
}

// create newblock with its stirng and prev block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{[]byte(data), prevBlockHash, []byte{}, time.Now().Unix()}
	block.SetHash()
	return block
}

func NewGenBlock() *Block{
	return NewBlock("Genesis Block", []byte{})
}