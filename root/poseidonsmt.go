package root

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rarimo/zkverifier-kit/internal/poseidonsmt"
)

// PoseidonSMTVerifier is a wrapper around PoseidonSMT binding which calls
// IsRootValid on the contract. Currently used for GlobalPassport and
// GeorgianPassport proof types.
type PoseidonSMTVerifier struct {
	caller  *poseidonsmt.PoseidonSMTCaller
	timeout time.Duration
}

func NewPoseidonSMTVerifier(rpcURL, contract string, timeout time.Duration) (*PoseidonSMTVerifier, error) {
	cli, addr, err := prepareBindingData(rpcURL, contract)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare binding data: %w", err)
	}

	caller, err := poseidonsmt.NewPoseidonSMTCaller(addr, cli)
	if err != nil {
		return nil, fmt.Errorf("failed to bind PoseidonSMT caller: %w", err)
	}

	return &PoseidonSMTVerifier{caller: caller, timeout: timeout}, nil
}

func (v *PoseidonSMTVerifier) VerifyRoot(root string) error {
	bytes := decimalTo32Bytes(root)
	if bytes == nil {
		return ErrInvalidRoot
	}

	ctx, cancel := context.WithTimeout(context.Background(), v.timeout)
	defer cancel()

	valid, err := v.caller.IsRootValid(&bind.CallOpts{Context: ctx}, *bytes)
	if err != nil {
		return fmt.Errorf("call IsRootValid on PoseidonSMT: %w", err)
	}

	if !valid {
		return ErrInvalidRoot
	}

	return nil
}
