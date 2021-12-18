package mini_miner

type problem struct {
	Difficulty int   `json:"difficulty"`
	Block      block `json:"block"`
}

type block struct {
	Data  [][]interface{} `json:"data"`
	Nonce interface{}     `json:"nonce"`
}
