package lib

import "time"

type Block struct {
	id string 
	timestamp time.Time
	next_block []*Block
	prev_block *Block
	
	//transactions []Transaction
}

func NewBlock(id string, timestamp time.Time) *Block {
	block := &Block{
		id,
		timestamp,
		nil,
		nil,
	}
	
	return block
}
