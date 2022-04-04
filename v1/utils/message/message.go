package message

type IRequestMessage struct {
	Jsonrpc string `json:"jsonrpc"`
	Service string `json:"service"`
	Method  string `json:"method"`
	Params  string `json:"params"`
	ID      int    `json:"id"`
	Client  string `json:"client"`
}

type IResponseMessage struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	Error   string `json:"error"`
	Type    string `json:"type"`
	Params  string `json:"params"`
	ID      int64  `json:"id"`
}
