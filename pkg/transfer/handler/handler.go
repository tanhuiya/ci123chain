package handler

import (
	"github.com/tanhuiya/ci123chain/pkg/abci/types"
	"github.com/tanhuiya/ci123chain/pkg/account/keeper"
	"github.com/tanhuiya/ci123chain/pkg/db"
	"github.com/tanhuiya/ci123chain/pkg/transaction"
	"github.com/tanhuiya/ci123chain/pkg/transfer"
	"reflect"
)

func NewHandler(
	txm transaction.TxIndexMapper,
	am keeper.AccountKeeper,
	sm *db.StateManager) types.Handler {
	return func(ctx types.Context, tx types.Tx) types.Result{
		ctx = ctx.WithTxIndex(txm.Get(ctx))
		defer func() {
			txm.Incr(ctx)
		}()
		switch tx := tx.(type) {
		case *transfer.TransferTx:
			return handlerTransferTx(ctx, am, tx)
		// todo

		default:
			errMsg := "Unrecognized Tx type: " + reflect.TypeOf(tx).Name()
			return types.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handlerTransferTx(ctx types.Context, am keeper.AccountKeeper, tx *transfer.TransferTx) types.Result {
	if err := am.Transfer(ctx, tx.Common.From, tx.To, tx.Amount); err != nil {
		return err.Result()
	}

	//交易成功，nonce+1
	account := am.GetAccount(ctx, tx.Common.From)
	saveErr := account.SetSequence(account.GetSequence() + 1)
	if saveErr != nil {
		return types.ErrInvalidSequence("Unexpected nonce of transaction").Result()
	}
	am.SetAccount(ctx, account)
	//
	return types.Result{}
}