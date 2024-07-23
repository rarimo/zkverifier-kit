package voting

import (
	"context"
)

// RootValidator allows validate root value: it is valid when nil is returned
type RootValidator interface {
	ValidateRoot(ctx context.Context, root [32]byte) error
}
