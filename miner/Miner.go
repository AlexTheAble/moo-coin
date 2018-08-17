package main

func main() {

}

type Transaction struct {
	toKey   string
	fromKey string
	amount  float64
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


func validationTransaction(trans Transaction*, chain Chain*) {
	user := trans.toKey
	
	for 
}
