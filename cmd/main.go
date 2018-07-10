package main

import "github.com/Tinywan/blockchain/core"


// func clusHash(toBeHashed string) string {
// 	// []byte(toBeHashed) []byte和string相互转换
// 	hashInBytes := sha256.Sum256([]byte(toBeHashed))
// 	// 把[]byte表示成16进制（用String的形式）
// 	hashInString := hex.EncodeToString(hashInBytes[:])
// 	// 打印log日志
// 	log.Printf("%s,%s",toBeHashed,hashInString)
// 	return hashInString
// }

func main(){
	// clusHash("test1")
	// clusHash("test1")
	// clusHash("testl")
	bc := core.NewBlockchain();
	bc.SendData("Send 1 BTC to Jacky")
	bc.SendData("Send 1 Eox to Jack")
	bc.Print()
}