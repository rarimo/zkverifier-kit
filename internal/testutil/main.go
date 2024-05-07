package testutil

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type MockCaller struct {
	root [32]byte
}

func NewMockCaller(root string) *MockCaller {
	return &MockCaller{root: hexToBytes(root)}
}

func (m *MockCaller) SetRoot(root string) {
	m.root = hexToBytes(root)
}

func (m *MockCaller) IcaoMasterTreeMerkleRoot(_ *bind.CallOpts) ([32]byte, error) {
	return m.root, nil
}

func hexToBytes(h string) [32]byte {
	var b [32]byte
	bs, err := hex.DecodeString(h)
	if err != nil {
		panic(fmt.Errorf("failed to decode hex: %w", err))
	}

	copy(b[:], bs)
	return b
}
