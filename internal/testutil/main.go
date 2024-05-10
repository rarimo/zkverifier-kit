package testutil

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type MockCaller struct {
	root []byte
}

func (m *MockCaller) WithRoot(root string) *MockCaller {
	return &MockCaller{
		root: hexToBytes(root),
	}
}

func (m *MockCaller) IsRootValid(_ *bind.CallOpts, root [32]byte) (bool, error) {
	return bytes.Equal(root[:], m.root), nil
}

func hexToBytes(h string) []byte {
	bs, err := hex.DecodeString(h)
	if err != nil {
		panic(fmt.Errorf("failed to decode hex: %w", err))
	}

	return bs
}
