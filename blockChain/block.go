package main

import (
	"crypto/sha256"
)

//定义结构
type block struct {
	//前区块哈希
	PrevHash []byte
	//当前区块哈希
	Hash []byte
	//数据
	Data []byte
}

//创建区块
func newBlock(data string, prevBloackHash []byte) *block {
	block := block{
		PrevHash: prevBloackHash,
		Hash:     []byte{},
		Data:     []byte(data),
	}

	block.setHash()

	return &block
}

//生成哈希
func (block *block) setHash() {
	//拼装数据
	blockInfo := append(block.PrevHash, block.Data...)
	//sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
