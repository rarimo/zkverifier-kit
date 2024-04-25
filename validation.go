package zkverifier_kit

import (
	"errors"
	"fmt"
	"math/big"
	"time"
)

type (
	// timeRule is a helper structure that might be used for validating the time.
	timeRule struct {
		// point - timestamp by which date will be validated
		point time.Time
		// isBefore - flag that shows whether the point has to be before or after some date
		isBefore bool
	}
)

// Validate is an implementation for `github.com/go-ozzo/ozzo-validation/v4` package to validate natively.
func (r timeRule) Validate(date interface{}) error {
	raw, ok := date.(string)
	if !ok {
		return fmt.Errorf("invalid type: %T, expected string", date)
	}

	parsed, err := time.Parse("060102", mustDecodeInt(raw))
	if err != nil {
		return fmt.Errorf("invalid date string: %w", err)
	}

	if r.isBefore && parsed.After(r.point) {
		return errors.New("date is too late")
	}

	if !r.isBefore && parsed.Before(r.point) {
		return errors.New("date is too early")
	}

	return nil
}

// shorthand to init timeRule BEFORE the point
func beforeDate(point time.Time) timeRule {
	return timeRule{
		point:    point,
		isBefore: true,
	}
}

// shorthand to init timeRule AFTER the point
func afterDate(point time.Time) timeRule {
	return timeRule{
		point:    point,
		isBefore: false,
	}
}

// function for clear way to encode some integer in bytes representation to its string analogue.
// It uses big.Int to convert the value.
func encodeInt(b []byte) string {
	return new(big.Int).SetBytes(b).String()
}

// function to decode string in decimals representation to its binary format. If it is failed to
// set string in big.Int the new empty one will be returned.
func mustDecodeInt(s string) string {
	b, ok := new(big.Int).SetString(s, 10)
	if !ok {
		b = new(big.Int)
	}
	return string(b.Bytes())
}
