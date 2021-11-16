package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct{
	Hash []byte
	Data []byte
	PrevHash []byte
	Nonce int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{},[]byte(data),prevHash, 0}
	pow :=NewProof9block
	nonce,hash:=pow.Run()
	block.Hash  = hash[:]
	block.Nonce = nonce
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
