package zkverifier_kit

type (
	// proofType defines public signals, their indexes and verification rules in ZKP
	proofType int

	// pubSignalID is public signal identifiers (not index!) used to get the
	// corresponding index from map[pubSignalID]int global variables. This is the
	// only intended usage.
	pubSignalID int
)

const (
	GlobalPassport proofType = iota
	GeorgianPassport
	PollParticipation
)

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

	ParticipationEventID
	NullifiersTreeRoot
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
		CurrentDate:               13,
		TimestampUpperBound:       15,
		IdentityCounterUpperBound: 17,
		BirthdateUpperBound:       19,
		ExpirationDateLowerBound:  20,
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
	pubPollParticipation = map[pubSignalID]int{
		Nullifier:            0,
		NullifiersTreeRoot:   1,
		ParticipationEventID: 2,
		EventID:              3,
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
// id, proof type or pub signals.
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
	case PollParticipation:
		return pubPollParticipation
	default:
		panic("unknown proof type")
	}
}

// PubSignalsCount returns the exact count of pub signals in proof. Use for
// validation on need to access specific fields, as Verifier already validates
// length.
func PubSignalsCount(t proofType) int {
	switch t {
	case GlobalPassport:
		return 22
	case GeorgianPassport:
		return 24
	case PollParticipation:
		return 4
	default:
		panic("unknown proof type")
	}
}
