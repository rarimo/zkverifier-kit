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
	filter  *proposalsmt.ProposalSMTFilterer
	timeout time.Duration
}

func NewProposalSMTVerifier(rpcURL, contract string, timeout time.Duration) (*ProposalSMTVerifier, error) {
	cli, addr, err := prepareBindingData(rpcURL, contract)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare binding data: %w", err)
	}

	filter, err := proposalsmt.NewProposalSMTFilterer(addr, cli)
	if err != nil {
		return nil, fmt.Errorf("failed to bind ProposalSMT caller: %w", err)
	}

	return &ProposalSMTVerifier{filter: filter, timeout: timeout}, nil
}

func (v *ProposalSMTVerifier) VerifyRoot(root string) error {
	bytes := decimalTo32Bytes(root)
	if bytes == nil {
		return ErrInvalidRoot
	}

	ctx, cancel := context.WithTimeout(context.Background(), v.timeout)
	defer cancel()

	it, err := v.filter.FilterRootUpdated(&bind.FilterOpts{Context: ctx}, [][32]byte{*bytes})
	if err != nil {
		return fmt.Errorf("filtering RootUpdated events: %w", err)
	}

	if ok := it.Next(); !ok {
		return ErrInvalidRoot
	}

	return nil
}
