package zkverifier_kit

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"time"

	val "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	eventData []byte

	timeRule struct {
		point       time.Time
		isBefore    bool
		isEqualDate bool
	}
)

func (val eventData) Validate(data interface{}) error {
	str, ok := data.(string)
	if !ok {
		return fmt.Errorf("invalid type: %T, expected string", data)
	}

	if !bytes.Equal([]byte(decodeInt(str)), val) {
		return fmt.Errorf("event data does not match")
	}

	return nil
}

func (r timeRule) Validate(date interface{}) error {
	raw, ok := date.(string)
	if !ok {
		return fmt.Errorf("invalid type: %T, expected string", date)
	}

	bigDecimalDate, ok := new(big.Int).SetString(raw, 10)
	if !ok {
		return fmt.Errorf("failed to set string: %T", date)
	}

	parsed, err := time.Parse("060102", string(bigDecimalDate.Bytes()))
	if err != nil {
		return fmt.Errorf("invalid date string: %w", err)
	}

	if r.isEqualDate {
		if !datesEqual(r.point, parsed) {
			return errors.New("dates are not equal")
		}
		return nil
	}

	if r.isBefore && parsed.After(r.point) {
		return errors.New("date is too late")
	}

	if !r.isBefore && parsed.Before(r.point) {
		return errors.New("date is too early")
	}

	return nil
}

func datesEqual(one time.Time, another time.Time) bool {
	return one.Format(time.DateOnly) == another.Format(time.DateOnly)
}

func beforeDate(point time.Time) timeRule {
	return timeRule{
		point:       point,
		isBefore:    true,
		isEqualDate: false,
	}
}

func afterDate(point time.Time) timeRule {
	return timeRule{
		point:       point,
		isBefore:    false,
		isEqualDate: false,
	}
}

func equalDate(point time.Time) timeRule {
	return timeRule{
		point:       point,
		isBefore:    false,
		isEqualDate: true,
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

func validateOnOptSet(value, option any, rules val.Rule) error {
	return val.Validate(value, val.When(
		!val.IsEmpty(option),
		val.Required,
		rules,
	))
}
