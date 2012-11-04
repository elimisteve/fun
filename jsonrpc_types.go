// Steve Phillips / elimisteve
// 2012.04.29
// Originally part of Decentra prototype

package fun

type JsonRpc1Request struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
	Id     interface{}   `json:"id"`
}

type JsonRpc1Response struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
	Id     interface{} `json:"id"`
}
