package server

import (
	"github.com/tanhuiya/ci123chain/pkg/gateway/types"
	"net/http"
)

type ErrProxy struct {

}

func NewErrProxy(pt types.ProxyType) *ErrProxy {
	return &ErrProxy{}
}

func (ep *ErrProxy) Handle(r *http.Request, backends []types.Instance, reqBody []byte) {
	//do nothing
	return
}

func (ep *ErrProxy) Response() *chan []byte {
	return nil
}