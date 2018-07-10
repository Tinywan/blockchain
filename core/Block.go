package core

import (
	"time"
	"crypto/sha256"
	"encoding/hex"
)
type Block struct {
	Index int64 // 区块编号
	Timestamp int64 // 区块时间戳
	PrevBlockHash string // 上一个区块的哈希值
	Hash string // 当前区块hash值

	Data string  // 区块数据
}

// 获取区块链hash 值
func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Hash
	blockInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(blockInBytes[:])
}

// 创建区块链
func GenerateNewBlock(prevBlock Block,data string) Block {
	newBlock := Block{}
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.PrevBlockHash = prevBlock.Hash
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

// 创始区块，不依赖下一个区块,上级区块为空
func GenerateGenesisBlock() Block{
	prevBlock := Block{}
	prevBlock.Index = -1
	prevBlock.Hash = ""
	return GenerateNewBlock(prevBlock,"Genesis Block")
}