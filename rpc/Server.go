package main

import (
	"io"
	"encoding/json"
	"net/http"
	"github.com/Tinywan/blockchain/core"
)

var blockchain *core.Blockchain 

func run(){
	http.HandleFunc("/blockchain/get",blockchainGetHanle)
	http.HandleFunc("/blockchain/write",blockchainWriteHanle)
	// 开始监http听端口
	http.ListenAndServe("localhost:8888",nil)
}

// 获取区块链
func blockchainGetHanle(w http.ResponseWriter,r *http.Request){
	bytes, error := json.Marshal(blockchain)
	if error != nil {
		http.Error(w,error.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(w,string(bytes))
}

// 写入区块链
func blockchainWriteHanle(w http.ResponseWriter,r *http.Request){
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGetHanle(w,r) 
}

func main(){
	blockchain = core.NewBlockchain()
	run()
}