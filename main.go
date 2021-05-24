package main 
import (
	"fmt"
    "golang-blockchain/blockchain"
)


func main() {
    chain := blockchain.InitBlockChain()
    chain.AddBlock("after genesis")
    chain.AddBlock("after genesis")
    chain.AddBlock("d Block after Genesis")

    for _, block := range chain.Blocks {
        fmt.Printf("Previous Hash: &x\n", block.PrevHash)
        fmt.Printf("Data in Block: %s\n", block.Data)
        fmt.Printf("Hash: %x\n", block.Hash)
    }
}