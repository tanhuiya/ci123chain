package types

import (
	sdk "github.com/tanhuiya/ci123chain/pkg/abci/types"
)


type CodeType = sdk.CodeType
const (
	DefaultCodespace sdk.CodespaceType = "wasm"
)

const (
	CodeCheckParamsError	CodeType = 50
	CodeInvalidMsgError     CodeType = 51
	CodeHandleMsgFailedError  CodeType = 52
	CodeSetSequenceFailedError CodeType = 53
)

func ErrCheckParams(codespace sdk.CodespaceType, keyname string) sdk.Error {
	return sdk.NewError(codespace, CodeCheckParamsError, "param invalid: %s", keyname)
}

func ErrInvalidMsg(codespce sdk.CodespaceType, err error) sdk.Error {
	return sdk.NewError(codespce, CodeInvalidMsgError, "msg invalid: %s", err.Error())
}

func ErrCreateFailed(codespce sdk.CodespaceType, err error) sdk.Error {
	return sdk.NewError(codespce, CodeHandleMsgFailedError, "create failed: %s", err.Error())
}

/*func ErrCheckWasmCode(codespce sdk.CodespaceType, err error) sdk.Error {
	return sdk.NewError(codespce, CodeHandleMsgFailedError, "uncompress code failed", err)
}*/

func ErrInstantiateFailed(codespce sdk.CodespaceType, err error) sdk.Error {
	return sdk.NewError(codespce, CodeHandleMsgFailedError, "instantiate failed: %s", err.Error())
}


func ErrExecuteFailed(codespce sdk.CodespaceType, err error) sdk.Error {
	return sdk.NewError(codespce, CodeHandleMsgFailedError, "execute failed: %s", err.Error())
}

func ErrSetNewAccountSequence(codespce sdk.CodespaceType, err error) sdk.Error {
	return sdk.NewError(codespce, CodeSetSequenceFailedError, "set sequence of account failed: %s", err.Error())
}
