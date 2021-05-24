package bloccChain
import (
    "bytes"
    "crypto/sha256"
)

//i am creating a blockchain with local MEMORY
type BloccChain struct {
    Blocks []*Block
}
type Block struct {
    Hash    []byte
    Data    []byte
    PHash   []byte
}

func (b *Block) DeriveHash() {
    info := bytes.Join([][]byte{b.Data, b.PHash}, []byte{}) 
    hash := sha256.Sum256(info)
    b.Hash = hash[:]
}

func CreateBlock(data string, phash []byte) *Block {
    block := &Block{[]byte{}, []byte(data), phash}
    block.DeriveHash()
    return block
}

func (chain *BloccChain) AddBlock(data string) {
    pBlock := chain.Blocks[len(chain.Blocks)-1]
    new := CreateBlock(data, pBlock.Hash)
    chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
    return CreateBlock("Genesis", []byte{})
}
func InitBloccChain() *BloccChain {
    return &BloccChain{[]*Block{Genesis()}}
}