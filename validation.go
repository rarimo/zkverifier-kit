package zkverifier_kit

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"time"

	"github.com/cosmos/btcutil/bech32"
	"github.com/ethereum/go-ethereum/common"
	val "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	eventDataRule struct {
		wantAddr string
		wantRaw  any
	}

	timeRule struct {
		point       time.Time
		isBefore    bool
		isEqualDate bool
	}
)

func (r eventDataRule) Validate(data interface{}) error {
	str, ok := data.(string)
	if !ok {
		return fmt.Errorf("invalid type: %T, expected string", data)
	}
	decodedData := []byte(decodeInt(str))

	if r.wantAddr == "" {
		return val.Validate(decodedData, val.In(r.wantRaw))
	}

	if isEthereumAddress(decodedData) {
		return val.Validate(common.BytesToAddress(decodedData).String(), val.In(r.wantRaw))
	}

	addr, err := bech32.Encode("rarimo", decodedData)
	if err != nil {
		return fmt.Errorf("invalid bech32 address: %w", err)
	}

	return val.Validate(addr, val.In(r.wantAddr))
}

func isEthereumAddress(data []byte) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(hex.EncodeToString(data))
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

	parsed, err := time.Parse("020106", string(bigDecimalDate.Bytes()))
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

func matchesData(raw any) eventDataRule {
	return eventDataRule{
		wantRaw:  raw,
		wantAddr: "",
	}
}

func matchesAddress(addr string) eventDataRule {
	return eventDataRule{
		wantAddr: addr,
		wantRaw:  nil,
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
