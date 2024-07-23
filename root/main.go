package root

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ErrInvalidRoot shows that identity root verification flow completed
// without internal errors, but the root itself is invalid
var ErrInvalidRoot = errors.New("invalid identity root")

// Verifier is an abstraction to verify the root value against some state. Root
// must be provided as decimal string, acquired from ZK-proof public signals,
// which is then converted tot 32-byte array.
type Verifier interface {
	VerifyRoot(root string) error
}

type VerifierType string

const (
	PoseidonSMT VerifierType = "poseidonsmt_root_verifier"
	ProposalSMT              = "proposalsmt_root_verifier"
)

func decimalTo32Bytes(root string) *[32]byte {
	b, ok := new(big.Int).SetString(root, 10)
	if !ok {
		return nil
	}

	var bytes [32]byte
	b.FillBytes(bytes[:])

	return &bytes
}

func prepareBindingData(rpc, contract string) (cli *ethclient.Client, addr common.Address, err error) {
	if !common.IsHexAddress(contract) {
		err = errors.New("invalid hex address")
		return
	}

	cli, err = ethclient.Dial(rpc)
	if err != nil {
		err = fmt.Errorf("dial RPC: %w", err)
		return
	}

	addr = common.HexToAddress(contract)
	return
}
