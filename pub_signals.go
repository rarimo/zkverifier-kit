package zkverifier_kit

// proofType defines public signals, their indexes and verification rules in ZKP
type proofType int

const (
	GlobalPassport proofType = iota
	GeorgianPassport
)

// pubSignalID is public signal identifiers (not index!) used to get the
// corresponding index from map[pubSignalID]int global variables. This is the
// only intended usage.
type pubSignalID int

const (
	Nullifier pubSignalID = iota
	BirthDate
	ExpirationDate
	Citizenship
	EventID
	EventData
	IdStateRoot
	Selector
	TimestampUpperBound
	IdentityCounterUpperBound
	BirthdateUpperBound
	ExpirationDateLowerBound
	PersonalNumberHash
	DocumentType
	CurrentDate
)

var (
	pubGlobalPassport = map[pubSignalID]int{
		Nullifier:                 0,
		BirthDate:                 1,
		ExpirationDate:            2,
		Citizenship:               6,
		EventID:                   9,
		EventData:                 10,
		IdStateRoot:               11,
		Selector:                  12,
		TimestampUpperBound:       14,
		IdentityCounterUpperBound: 16,
		BirthdateUpperBound:       18,
		ExpirationDateLowerBound:  19,
	}
	pubGeorgianPassport = map[pubSignalID]int{
		Nullifier:                 0,
		BirthDate:                 1,
		ExpirationDate:            2,
		Citizenship:               5,
		PersonalNumberHash:        8,
		DocumentType:              9,
		EventID:                   10,
		EventData:                 11,
		IdStateRoot:               12,
		Selector:                  13,
		CurrentDate:               14,
		TimestampUpperBound:       16,
		IdentityCounterUpperBound: 18,
		BirthdateUpperBound:       20,
		ExpirationDateLowerBound:  21,
	}
)

// PubSignalGetter is a structure to extract public signals from abstract ZKP in
// a convenient way. It is an alternative for Indexes to initialize once and
// reuse for the same proof type and signals.
type PubSignalGetter struct {
	ProofType proofType
	Signals   []string
}

// Get extracts public signal by its identifier. Returns empty string on invalid
// id or pub signals.
func (p *PubSignalGetter) Get(id pubSignalID) string {
	i, ok := Indexes(p.ProofType)[id]
	if !ok || len(p.Signals) <= i {
		return ""
	}
	return p.Signals[i]
}

// Indexes returns public signals indexes based on proof type provided. Use it
// when accessing public signals values in provided ZKP. Proof type must be
// supported by this package.
func Indexes(t proofType) map[pubSignalID]int {
	switch t {
	case GlobalPassport:
		return pubGlobalPassport
	case GeorgianPassport:
		return pubGeorgianPassport
	default:
		panic("unknown proof type")
	}
}
