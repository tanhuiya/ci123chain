package sdk

import (
	"github.com/tanhuiya/ci123chain/pkg/abci/types"
	"github.com/tanhuiya/ci123chain/pkg/app"
	"github.com/tanhuiya/ci123chain/pkg/client"
	"github.com/tanhuiya/ci123chain/pkg/client/helper"
	"github.com/tanhuiya/ci123chain/pkg/cryptosuit"
	"github.com/tanhuiya/ci123chain/pkg/mortgage"
	types2 "github.com/tanhuiya/ci123chain/pkg/mortgage/types"
)
var cdc = app.MakeCodec()
// 生成 Mortgage 消息，抵押coin
func SignMortgage(from, to string, amount, gas uint64, uniqueID string, priv []byte) ([]byte, error) {
	tx, err := buildMortgageTx(from, to, amount, gas, uniqueID)
	if err != nil {
		return nil, err
	}
	sid := cryptosuit.NewFabSignIdentity()
	pub, err  := sid.GetPubKey(priv)

	tx.SetPubKey(pub)
	signbyte := tx.GetSignBytes()
	signature, err := sid.Sign(signbyte, priv)
	tx.SetSignature(signature)
	return tx.Bytes(), nil
}

func buildMortgageTx (from, to string, amount, gas uint64, uniqueID string) (*types2.MsgMortgage, error) {
	fromAddr, err := helper.StrToAddress(from)
	if err != nil {
		return nil, err
	}
	toAddr, err := helper.StrToAddress(to)
	if err != nil {
		return nil, err
	}
	ctx, err := client.NewClientContextFromViper(cdc)
	if err != nil {
		return nil,err
	}
	nonce, err := ctx.GetNonceByAddress(fromAddr)
	mort := mortgage.NewMortgageMsg(fromAddr, toAddr, gas, nonce, types.NewUInt64Coin(amount), []byte(uniqueID))
	return mort, nil
}

// 生成 MortgageDone 完成交易
func SignMortgageDone(from string, gas uint64, uniqueID string, priv []byte) ([]byte, error) {
	tx, err := buildMortgageDoneTx(from, gas, uniqueID)
	if err != nil {
		return nil, err
	}
	sid := cryptosuit.NewFabSignIdentity()
	pub, err  := sid.GetPubKey(priv)

	tx.SetPubKey(pub)
	signbyte := tx.GetSignBytes()
	signature, err := sid.Sign(signbyte, priv)
	tx.SetSignature(signature)
	return tx.Bytes(), nil
}



func buildMortgageDoneTx (from string, gas uint64, uniqueID string) (*types2.MsgMortgageDone, error) {
	fromAddr, err := helper.StrToAddress(from)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	ctx, err := client.NewClientContextFromViper(cdc)
	if err != nil {
		return nil,err
	}
	nonce, err := ctx.GetNonceByAddress(fromAddr)
	mort := mortgage.NewMsgMortgageDone(fromAddr, gas, nonce, []byte(uniqueID))
	return mort, nil
}


// MortgageCancel
// 取消 交易，将 coin 返还
func SignMortgageCancel(from string, gas uint64, uniqueID string, priv []byte) ([]byte, error) {
	tx, err := buildMortgageCancelTx(from, gas, uniqueID)
	if err != nil {
		return nil, err
	}
	sid := cryptosuit.NewFabSignIdentity()
	pub, err  := sid.GetPubKey(priv)

	tx.SetPubKey(pub)
	signbyte := tx.GetSignBytes()
	signature, err := sid.Sign(signbyte, priv)
	tx.SetSignature(signature)
	return tx.Bytes(), nil
}

func buildMortgageCancelTx (from string, gas uint64, uniqueID string) (*types2.MsgMortgageCancel, error) {
	fromAddr, err := helper.StrToAddress(from)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	ctx, err := client.NewClientContextFromViper(cdc)
	if err != nil {
		return nil,err
	}

	nonce, err := ctx.GetNonceByAddress(fromAddr)
	mort := mortgage.NewMsgMortgageCancel(fromAddr, gas, nonce, []byte(uniqueID))
	return mort, nil
}