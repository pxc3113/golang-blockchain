package main

import (
	"fmt"
	"github.com/pxc3110/golang-blockchain/blockhain"
)

func main()  {
	chain := InitBlockChain()
	chain.AddBlock("1")
	chain.AddBlock("2")
	chain.AddBlock("3")

	for _, block := range chain.blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow:=blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate(0)))
		fmt.Println()
	}
}