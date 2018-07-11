package core

import (
	"fmt"
	"log"
)

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := GenerateGenesisBlock() // 创建初始区块
	blockchain := Blockchain{}
	blockchain.AppendBlock(&genesisBlock)
	return &blockchain
} 

func (bc *Blockchain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks) - 1]
	nextBlock := GenerateNewBlock(*preBlock ,data)
	bc.AppendBlock(&nextBlock)
}

// 添加新的区块
func (bc *Blockchain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return 
	}

	if isValid(*newBlock,*bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks,newBlock)
	}else{
		log.Fatal("Invalid Block")
	}
}

// 打印区块链数据
func (bc *Blockchain) Print() {
	for _,block := range bc.Blocks {
		fmt.Printf("Index: %d \n",block.Index)
		fmt.Printf("PrevBlockHash: %s \n",block.PrevBlockHash)
		fmt.Printf("CurrHash: %s \n",block.Hash)
		fmt.Printf("Data: %s \n",block.Data)
		fmt.Printf("Timestamp: %d \n",block.Timestamp)
		fmt.Println()
	}
}

func isValid(newBlock Block,oldBlock Block) bool {
	if newBlock.Index - 1 != oldBlock.Index {
		return false
	}

	// 检测Hash值
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}

	// if calculateHash(newBlock) != newBlock.Hash {
	// 	return false
	// }

	return true
}