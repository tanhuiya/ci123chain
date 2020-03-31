package types

import (
	"encoding/json"
	sdk "github.com/tanhuiya/ci123chain/pkg/abci/types"
	"github.com/tanhuiya/ci123chain/pkg/transaction"
)


type StoreCodeTx struct {
	transaction.CommonTx
	Sender      sdk.AccAddress    `json:"sender"`
	WASMByteCode []byte           `json:"wasm_byte_code"`
	Source      string            `json:"source"`
	Builder     string            `json:"builder"`
}

func NewStoreCodeTx(from sdk.AccAddress, gas, nonce uint64, sender sdk.AccAddress, wasmCode []byte, source, builder string) StoreCodeTx{

	return StoreCodeTx{
		CommonTx:     transaction.CommonTx{
			From:  from,
			Gas:   gas,
			Nonce: nonce,
		},
		Sender:       sender,
		WASMByteCode: wasmCode,
		Source:       source,
		Builder:      builder,
	}
}

//TODO
func (msg *StoreCodeTx) ValidateBasic() sdk.Error {
	return nil
}

func (msg *StoreCodeTx) GetSignBytes() []byte {
	tmsg := *msg
	tmsg.Signature = nil
	signBytes := tmsg.Bytes()
	return signBytes
}

func (msg *StoreCodeTx) SetSignature(sig []byte) {}

func (msg *StoreCodeTx) Bytes() []byte {
	bytes, err := WasmCodec.MarshalBinaryLengthPrefixed(msg)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (msg *StoreCodeTx) SetPubKey(pub []byte) {
	msg.PubKey = pub
}

func (msg *StoreCodeTx) Route() string {
	return RouteKey
}

func (msg *StoreCodeTx)GetGas() uint64 {
	return msg.Gas
}

func (msg *StoreCodeTx)GetNonce() uint64 {
	return msg.Nonce
}

func (msg *StoreCodeTx) GetFromAddress() sdk.AccAddress {
	return msg.From
}


type InstantiateContractTx struct {
	transaction.CommonTx
	Sender       sdk.AccAddress       `json:"sender"`
	CodeID       uint64               `json:"code_id"`
	Label        string               `json:"label"`
	InitMsg      json.RawMessage      `json:"init_msg"`
	InitFunds    sdk.Coin             `json:"init_funds"`
}

func NewInstantiateContractTx(from sdk.AccAddress, gas, nonce, codeID uint64, sender sdk.AccAddress, label string,
	initMsg json.RawMessage, initFunds sdk.Coin) InstantiateContractTx{

		return InstantiateContractTx{
			CommonTx: transaction.CommonTx{
				From:  from,
				Gas:   gas,
				Nonce: nonce,
			},
			Sender:    sender,
			CodeID:    codeID,
			Label:     label,
			InitMsg:   initMsg,
			InitFunds: initFunds,
		}
}


//TODO
func (msg *InstantiateContractTx) ValidateBasic() sdk.Error {
	return nil
}

func (msg *InstantiateContractTx) GetSignBytes() []byte {
	tmsg := *msg
	tmsg.Signature = nil
	signBytes := tmsg.Bytes()
	return signBytes
}

func (msg *InstantiateContractTx) SetSignature(sig []byte) {
	msg.Signature = sig
}

func (msg *InstantiateContractTx) Bytes() []byte {
	bytes, err := WasmCodec.MarshalBinaryLengthPrefixed(msg)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (msg *InstantiateContractTx) SetPubKey(pub []byte) {
	msg.PubKey = pub
}

func (msg *InstantiateContractTx) Route() string {
	return RouteKey
}

func (msg *InstantiateContractTx) GetGas() uint64 {
	return msg.Gas
}

func (msg *InstantiateContractTx) GetNonce() uint64 {
	return msg.Nonce
}

func (msg *InstantiateContractTx) GetFromAddress() sdk.AccAddress {
	return msg.From
}


type ExecuteContractTx struct {
	transaction.CommonTx
	Sender           sdk.AccAddress      `json:"sender"`
	Contract         sdk.AccAddress      `json:"contract"`
	Msg              json.RawMessage     `json:"msg"`
	SendFunds        sdk.Coin           `json:"send_funds"`
}

func NewExecuteContractTx(from sdk.AccAddress, gas, nonce uint64, sender sdk.AccAddress,
	contractAddress sdk.AccAddress, msg json.RawMessage, sendFunds sdk.Coin) ExecuteContractTx {

	return ExecuteContractTx{
		CommonTx:  transaction.CommonTx{
			From:      from,
			Nonce:     nonce,
			Gas:       gas,
		},
		Sender:    sender,
		Contract:  contractAddress,
		Msg:       msg,
		SendFunds: sendFunds,
	}
}

//TODO
func (msg *ExecuteContractTx) ValidateBasic() sdk.Error {
	return nil
}

func (msg *ExecuteContractTx) GetSignBytes() []byte {
	tmsg := *msg
	tmsg.Signature = nil
	signBytes := tmsg.Bytes()
	return signBytes
}

func (msg *ExecuteContractTx) SetSignature(sig []byte) {
	msg.Signature = sig
}

func (msg *ExecuteContractTx) Bytes() []byte {
	bytes, err := WasmCodec.MarshalBinaryLengthPrefixed(msg)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (msg *ExecuteContractTx) SetPubKey(pub []byte) {
	msg.PubKey = pub
}

func (msg *ExecuteContractTx) Route() string {
	return RouteKey
}

func (msg *ExecuteContractTx) GetGas() uint64 {
	return msg.Gas
}

func (msg *ExecuteContractTx) GetNonce() uint64 {
	return msg.Nonce
}

func (msg *ExecuteContractTx) GetFromAddress() sdk.AccAddress {
	return msg.From
}
