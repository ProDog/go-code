package demochain
import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index int64
	Timestamp int64
	PrevBlockHash string
	Hash string
	Data string
}

func calculateHash(b Block) string {
	blockData := string(b.Index)+string(b.Timestamp)+b.PrevBlockHash+b.Data
	hashBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashBytes[:])
}

func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index=preBlock.Index+1
	newBlock.PrevBlockHash=preBlock.Hash
	newBlock.Timestamp=time.Now().Unix()
	newBlock.Data=data
	newBlock.Hash=calculateHash(newBlock)
	return newBlock 
}

func GenerateGenesisiBlock() Block {
	preBlock := Block{}
	preBlock.Index=-1
	preBlock.Hash=""
	return GenerateNewBlock(preBlock, "Genesis Block")
}