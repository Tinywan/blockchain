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
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
	blockInBytes := sha256.Sum256([]byte(blockData)) // 字符串转换为字节切片
	return hex.EncodeToString(blockInBytes[:]) // 数组转换为字节切片
}

// 创建区块链
func GenerateNewBlock(preBlock Block,data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Hash = calculateHash(newBlock)
	newBlock.Data = data
	return newBlock
}

// 创始区块，不依赖下一个区块,上级区块为空
func GenerateGenesisBlock() Block{
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock,"Genesis Block")
}