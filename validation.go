package zkverifier_kit

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/cosmos/btcutil/bech32"
	val "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	eventDataRule struct {
		wantAddr string
		wantRaw  any
	}

	timeRule struct {
		point    time.Time
		isBefore bool
	}
)

func (r eventDataRule) Validate(data interface{}) error {
	if r.wantAddr == "" {
		return val.Validate(data, val.In(r.wantRaw))
	}

	str, ok := data.(string)
	if !ok {
		return fmt.Errorf("invalid type: %T, expected string", data)
	}

	addr, err := bech32.EncodeFromBase256("rarimo", []byte(decodeInt(str)))
	if err != nil {
		return fmt.Errorf("invalid bech32 string: %w", err)
	}

	return val.Validate(addr, val.In(r.wantAddr))
}

func (r timeRule) Validate(date interface{}) error {
	raw, ok := date.(string)
	if !ok {
		return fmt.Errorf("invalid type: %T, expected string", date)
	}

	parsed, err := time.Parse("060102", raw)
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

func matchesData(raw any) eventDataRule {
	return eventDataRule{
		wantRaw: raw,
	}
}

func matchesAddress(addr string) eventDataRule {
	return eventDataRule{
		wantAddr: addr,
	}
}

// decode big int from the proof to string
func decodeInt(s string) string {
	b, ok := new(big.Int).SetString(s, 10)
	if !ok {
		b = new(big.Int)
	}
	return string(b.Bytes())
}

func validateOnOptSet(value, option any, rule val.Rule) error {
	return val.Validate(value, val.When(
		!val.IsEmpty(option),
		val.Required,
		rule,
	))
}
