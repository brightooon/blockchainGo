package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const targetBit = 20
var ( maxNonce = math.MaxInt64)

type Proofofwork struct{
	block *Block
	tbit *big.Int
}

func ProofOfWorkfunc(block *Block) *Proofofwork {
	tbit := big.NewInt(1)
	tbit.Lsh(tbit, uint(256-targetBit)) // tbit<<(256-targetbit)
	pow := &Proofofwork{block, tbit}
	return pow 
}

func (proof *Proofofwork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			proof.block.PrevBlockHash,
			proof.block.Data,
			IntToHex(proof.block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (proof *Proofofwork) Runpow() (int, []byte) {
	var hashint big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n", proof.block.Data)
	for nonce < maxNonce {
		data := proof.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashint.SetBytes(hash[:])
		if hashint.Cmp(proof.tbit) == -1 {
			break
		}else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

func (proof *Proofofwork) Validatepow() bool {
	var hashint big.Int
	data := proof.prepareData((proof.block.Nonce))
	hash := sha256.Sum256(data)
	hashint.SetBytes(hash[:])
	valid := hashint.Cmp(proof.tbit) == -1
	return valid
}