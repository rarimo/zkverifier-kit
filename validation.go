package zkverifier_kit

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/cosmos/btcutil/bech32"
)

type (
	eventDataRule struct {
		want string
	}

	timeRule struct {
		point    time.Time
		isBefore bool
	}
)

func (r eventDataRule) Validate(data interface{}) error {
	str, ok := data.(string)
	if !ok {
		return fmt.Errorf("invalid type: %T, expected string", data)
	}

	addr, err := bech32.EncodeFromBase256("rarimo", []byte(decodeInt(str)))
	if err != nil {
		return fmt.Errorf("invalid bech32 string: %w", err)
	}

	if addr != r.want {
		return errors.New("event data does not match the address")
	}

	return nil
}

func (r timeRule) Validate(date interface{}) error {
	raw, ok := date.(string)
	if !ok {
		return fmt.Errorf("invalid type: %T, expected string", date)
	}

	parsed, err := time.Parse("060102", decodeInt(raw))
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

func beforeDate(point time.Time) timeRule {
	return timeRule{
		point:    point,
		isBefore: true,
	}
}

func afterDate(point time.Time) timeRule {
	return timeRule{
		point:    point,
		isBefore: false,
	}
}

func matchesAddress(addr string) eventDataRule {
	return eventDataRule{want: addr}
}

// decode big int from the proof to string
func decodeInt(s string) string {
	b, ok := new(big.Int).SetString(s, 10)
	if !ok {
		b = new(big.Int)
	}
	return string(b.Bytes())
}
