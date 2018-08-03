// MooChain
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var blockchain = generateBlockchain()
	blockchain.blocks = append(blockchain.blocks, generateBlock(blockchain.blocks[len(blockchain.blocks)-1], 5, nil))
	fmt.Printf("Final hash %s \n", blockchain.blocks[len(blockchain.blocks)-1].hash)
}

type Chain struct {
	blocks []*Block
}

type Block struct {
	index        int
	time         string
	hash         string
	prevHash     string
	difficulty   int
	nonce        int32
	reward       float64
	minerKey     string
	transactions []Transaction
}

type Transaction struct {
	toKey   string
	fromKey string
	amount  float64
}

func generateBlockchain() *Chain {
	genBlock := new(Block)
	genBlock.index = 0
	genBlock.time = time.Now().String()
	genBlock.prevHash = ""
	genBlock.difficulty = 1
	genBlock.minerKey = ""
	genBlock.reward = 1
	genBlock.transactions = nil
	genBlock.hash = calculateHash(genBlock)

	chain := new(Chain)
	chain.blocks = []*Block{genBlock}
	return chain
}

func generateBlock(prevBlock *Block, difficulty int, transactions []Transaction) *Block {
	block := new(Block)
	block.index = prevBlock.index + 1
	block.prevHash = prevBlock.hash
	block.time = time.Now().String()
	block.transactions = transactions
	block.difficulty = difficulty

	block.nonce = rand.Int31()
	block.hash = calculateHash(block)
	var index = 1
	for !isDifficultyMet(block.hash, difficulty) {
		fmt.Printf("try: %d hash: %s \n", index, block.hash)
		block.nonce = rand.Int31()
		block.hash = calculateHash(block)
		index = index + 1
	}

	block.hash = calculateHash(block)
	return block
}

func calculateHash(block *Block) string {
	transJson, err := json.Marshal(block.transactions)
	if err != nil {
		return ""
	}
	record := strconv.Itoa(block.index) + block.time + block.prevHash + strconv.Itoa(int(block.nonce)) + string(transJson)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func isDifficultyMet(hash string, difficulty int) bool {
	for i, char := range hash {
		if char != 48 {
			return false
		} else if i == difficulty-1 {
			return true
		}
	}
	return false
}
