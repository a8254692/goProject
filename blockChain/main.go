package main

import "fmt"

func main() {
	bc := newBlockChain()

	bc.addBlock("转入100个比特币")
	bc.addBlock("转入500个比特币")

	for i, block := range bc.blocks {

		fmt.Printf("当前区块高度 %d \n", i)
		fmt.Printf("前区块哈希 %x \n", block.PrevHash)
		fmt.Printf("当前块哈希 %x \n", block.Hash)
		fmt.Printf("区块数据 %s \n", block.Data)
	}
}

