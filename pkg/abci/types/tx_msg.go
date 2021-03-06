package types

//__________________________________________________________

// Transactions objects must fulfill the Tx
type Tx interface {
	// ValidateBasic does a simple and lightweight validation check that doesn't
	// require access to any other information.
	ValidateBasic() Error

	Route() string
}

//__________________________________________________________

// TxDecoder unmarshals transfer bytes
type TxDecoder func(txBytes []byte) (Tx, Error)
