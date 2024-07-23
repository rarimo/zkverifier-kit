package root

// DisabledVerifier returns nil error on verification
type DisabledVerifier struct{}

func (v DisabledVerifier) VerifyRoot(_ string) error {
	return nil
}
