package transfer

import (
	"github.com/tanhuiya/ci123chain/pkg/abci/codec"
	"github.com/tanhuiya/ci123chain/pkg/transaction"
	"github.com/tendermint/go-amino"
)

func RegisterCodec(cdc *amino.Codec)  {
	cdc.RegisterConcrete(&TransferTx{}, "transfer/transfer", nil)
}

var transferCdc *codec.Codec

func init()  {
	transferCdc = codec.New()
	transaction.RegisterCodec(transferCdc)
	RegisterCodec(transferCdc)
}