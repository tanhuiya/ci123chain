package types

import (
	"github.com/tanhuiya/ci123chain/pkg/abci/codec"
	"github.com/tanhuiya/ci123chain/pkg/transaction"
)

var MortgageCdc *codec.Codec

func RegisterCodec(cdc *codec.Codec) {

	cdc.RegisterConcrete(&MsgMortgage{}, "ci123chain/MsgMortgage", nil)
	cdc.RegisterConcrete(&MsgMortgageDone{}, "ci123chain/MsgMortgageDone", nil)
	cdc.RegisterConcrete(&MsgMortgageCancel{}, "ci123chain/MsgMortgageCancel", nil)
	cdc.RegisterConcrete(&Mortgage{}, "ci123chain/Mortgage", nil)
}

func init()  {
	MortgageCdc = codec.New()
	MortgageCdc.RegisterInterface((*transaction.Transaction)(nil), nil)
	MortgageCdc.RegisterConcrete(&transaction.CommonTx{}, "ci123chain/commontx", nil)
	RegisterCodec(MortgageCdc)
	MortgageCdc.Seal()
}