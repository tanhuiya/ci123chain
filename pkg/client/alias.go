package client

import "github.com/tanhuiya/ci123chain/pkg/client/types"

var (
	ErrNewClientCtx		= types.ErrNewClientCtx
	ErrGetInputAddr		= types.ErrGetInputAddrCtx
	ErrParseAddr		= types.ErrParseAddr
	ErrNoAddr       	= types.ErrNoAddr
	ErrGetPassPhrase	= types.ErrGetPassPhrase
	ErrGetSignData		= types.ErrGetSignData
	ErrBroadcast		= types.ErrBroadcast
	ErrGetCheckPassword	= types.ErrGetCheckPassword
	ErrGetPassword		= types.ErrGetPassword
	ErrPhrasesNotMatch	= types.ErrPhrasesNotMatch
	ErrNode				= types.ErrNode
)

type TMResponse struct {
	Jsonrpc  string `json:"jsonrpc"`
	ID      string  `json:"id"`
	Result   interface{} `json:"result"`
}

type Response struct {
	Ret 	uint32 	`json:"ret"`
	Data 	interface{}	`json:"data"`
	Message	string	`json:"message"`
}