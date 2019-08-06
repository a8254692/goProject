package main

type blockChain struct {
	//定义区块链数组
	blocks []*block
}

//引入区块链
func newBlockChain() *blockChain {

	//创建一个创世块，添加到区块链中
	genesisBlocks := genesisBlock()

	return &blockChain{
		blocks: []*block{genesisBlocks},
	}
}

//创世块
func genesisBlock() *block {
	return newBlock("创世块", []byte{})
}


//添加区块
func (bc *blockChain) addBlock(data string) {
	//获取前区块的hash
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash

	//生成区块
	block := newBlock(data, prevHash)

	//添加到区块组
	bc.blocks = append(bc.blocks, block)
}