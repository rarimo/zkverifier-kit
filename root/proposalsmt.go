package root

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rarimo/zkverifier-kit/internal/proposalsmt"
)

// ProposalSMTVerifier performs validation with filtering RootUpdated events of
// ProposalSMT contract by root value. Currently used for PollParticipation proof
// type.
type ProposalSMTVerifier struct {
	rpc     string
	addr    string
	timeout time.Duration
}

// NewProposalSMTVerifier creates basic ProposalSMTVerifier with RPC only. You
// must call WithContract to use it.
func NewProposalSMTVerifier(rpcURL string, timeout time.Duration) *ProposalSMTVerifier {
	return &ProposalSMTVerifier{
		rpc:     rpcURL,
		timeout: timeout,
	}
}

// WithContract returns new instance of ProposalSMTVerifier which will call the
// provided contract. Provided address must be a valid 20-byte hex.
func (v *ProposalSMTVerifier) WithContract(addr string) Verifier {
	return &ProposalSMTVerifier{
		rpc:     v.rpc,
		addr:    addr,
		timeout: v.timeout,
	}
}

func (v *ProposalSMTVerifier) VerifyRoot(root string) error {
	cli, addr, err := prepareBindingData(v.rpc, v.addr)
	if err != nil {
		return fmt.Errorf("failed to prepare binding data: %w", err)
	}

	bytes := decimalTo32Bytes(root)
	if bytes == nil {
		return ErrInvalidRoot
	}

	ctx, cancel := context.WithTimeout(context.Background(), v.timeout)
	defer cancel()

	filter, err := proposalsmt.NewProposalSMTFilterer(addr, cli)
	if err != nil {
		return fmt.Errorf("failed to bind ProposalSMT filter: %w", err)
	}

	it, err := filter.FilterRootUpdated(&bind.FilterOpts{Context: ctx}, [][32]byte{*bytes})
	if err != nil {
		return fmt.Errorf("filtering RootUpdated events: %w", err)
	}

	if ok := it.Next(); !ok {
		return ErrInvalidRoot
	}

	return nil
}
