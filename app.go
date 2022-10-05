package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
)

type Block struct {
	data         string
	nonce        int
	previousHash string
	previous_ptr *Block
	next_ptr     *Block
}

var root Block
var blocks_counter int = 0

// Function to calculate hash of a given string
func CalculateHash(mystring string) string {
	if mystring != "" {
		return fmt.Sprintf("%x", sha256.Sum256([]byte(mystring)))
	}

	return ""
}

// Function to display blocks
func DisplayBlocks() {
	var temp Block
	temp = root

	if temp.next_ptr == nil {
		fmt.Println("There is no block in the block chain!")
		return
	}

	var counter int = 1
	for {

		fmt.Printf("----------Block %d----------", counter)
		fmt.Printf("\nTransaction : %s", temp.data)
		fmt.Printf("\nNonce : %d", temp.nonce)
		fmt.Printf("\nPrevious Hash : %s\n", temp.previousHash)
		fmt.Print("\n------------------------------\n\n")
		if temp.next_ptr == nil {
			break
		}

		temp = *temp.next_ptr
		counter++
	}
}

func ChangeBlock(block_id int, newTransaction string) bool {
	var flag bool = true
	if block_id > 0 && block_id < blocks_counter {
		iBlock := &root

		for i := 1; i < block_id; i++ {
			iBlock = iBlock.next_ptr
		}

		iBlock.data = newTransaction
		return flag
	} else {
		fmt.Println("\nInvalid Block ID!")
	}

	return !flag
}

func VerifyChain() bool {
	temp := &root
	var flag bool = true
	var hash_1 string
	for i := 1; i < blocks_counter; i++ {
		hash_1 = CalculateHash(temp.data)
		temp = temp.next_ptr

		if temp.previousHash != hash_1 {
			fmt.Printf("\n\nBlock %d data has been CHANGED \n\n", i)
			return !flag
		}

	}

	fmt.Println("\n\nData is secure")
	return flag
}

func NewBlock(transaction string, nonce int, previousHash string) bool {
	var flag bool = true
	if root.next_ptr == nil && root.previous_ptr == nil {
		var temp Block
		temp.data = ""
		temp.nonce = -1
		temp.previousHash = CalculateHash(transaction)
		temp.next_ptr = nil

		root.data = transaction
		root.nonce = rand.Intn(10000)
		root.previousHash = ""
		root.previous_ptr = nil

		root.next_ptr = &temp
		temp.previous_ptr = &root

		blocks_counter = 2
		return flag
	} else if transaction != "" {
		iteratorBlock := &root
		for {
			if iteratorBlock.next_ptr == nil {
				break
			}

			iteratorBlock = iteratorBlock.next_ptr
		}

		var temp Block
		temp.data = ""
		temp.nonce = -1
		temp.previousHash = CalculateHash(transaction)
		temp.next_ptr = nil

		iteratorBlock.data = transaction
		iteratorBlock.nonce = rand.Intn(10000)
		temp.previous_ptr = iteratorBlock
		iteratorBlock.next_ptr = &temp

		blocks_counter += 1
		return flag
	} else {
		fmt.Println("New Block Addition failed\nParameters Invalid!")
	}

	return !flag
}
