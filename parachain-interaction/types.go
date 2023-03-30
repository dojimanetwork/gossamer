package parachaininteraction

import (
	"bytes"
	"fmt"

	"github.com/ChainSafe/gossamer/dot/types"
	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/ChainSafe/gossamer/pkg/scale"
)

// signature could be one of Ed25519 signature, Sr25519 signature or ECDSA/SECP256k1 signature.
type signature [64]byte

func (s signature) String() string { return fmt.Sprintf("0x%x", s[:]) }

// collatorID is the collator's relay-chain account ID
type collatorID []byte

// collatorSignature is the signature on a candidate's block data signed by a collator.
type collatorSignature signature

// validationCodeHash is the blake2-256 hash of the validation code bytes.
type validationCodeHash common.Hash

// candidateDescriptor is a unique descriptor of the candidate receipt.
type CandidateDescriptor struct {
	// The ID of the para this is a candidate for.
	ParaID uint32 `scale:"1"`

	// RelayParent is the hash of the relay-chain block this should be executed in
	// the context of.
	// NOTE: the fact that the hash includes this value means that code depends
	// on this for deduplication. Removing this field is likely to break things.
	RelayParent common.Hash `scale:"2"`

	// Collator is the collator's relay-chain account ID
	Collator collatorID `scale:"3"`

	// PersistedValidationDataHash is the blake2-256 hash of the persisted validation data. This is extra data derived from
	// relay-chain state which may vary based on bitfields included before the candidate.
	// Thus it cannot be derived entirely from the relay-parent.
	PersistedValidationDataHash common.Hash `scale:"4"`

	// PovHash is the hash of the `pov-block`.
	PovHash common.Hash `scale:"5"`
	// ErasureRoot is the root of a block's erasure encoding Merkle tree.
	ErasureRoot common.Hash `scale:"6"`

	// Signature on blake2-256 of components of this receipt:
	// The parachain index, the relay parent, the validation data hash, and the `pov_hash`.
	Signature collatorSignature `scale:"7"`

	// ParaHead is the hash of the para header that is being generated by this candidate.
	ParaHead common.Hash `scale:"8"`
	// ValidationCodeHash is the blake2-256 hash of the validation code bytes.
	ValidationCodeHash validationCodeHash `scale:"9"`
}

func (cd CandidateDescriptor) CheckCollatorSignature() error {
	// payload
	var payload [132]byte
	copy(payload[0:32], cd.RelayParent.ToBytes())
	buffer := bytes.NewBuffer(nil)
	encoder := scale.NewEncoder(buffer)
	err := encoder.Encode(cd.ParaID)
	if err != nil {
		return fmt.Errorf("encoding parachain id: %w", err)
	}
	copy(payload[32:32+buffer.Len()], buffer.Bytes())
	copy(payload[36:68], cd.PersistedValidationDataHash.ToBytes())
	copy(payload[68:100], cd.PovHash.ToBytes())
	copy(payload[100:132], common.Hash(cd.ValidationCodeHash).ToBytes())

	// collator public key cd.Collator
	// type of signature or just the verify function

}

type CandidateReceipt struct {
	descriptor      CandidateDescriptor `scale:"1"`
	commitmentsHash common.Hash         `scale:"2"`
}

// candidateCommitments are Commitments made in a `CandidateReceipt`. Many of these are outputs of validation.
type candidateCommitments struct {
	// Messages destined to be interpreted by the Relay chain itself.
	UpwardMessages []upwardMessage `scale:"1"`
	// Horizontal messages sent by the parachain.
	HorizontalMessages []outboundHrmpMessage `scale:"2"`
	// New validation code.
	NewValidationCode *ValidationCode `scale:"3"`
	// The head-data produced as a result of execution.
	HeadData headData `scale:"4"`
	// The number of messages processed from the DMQ.
	ProcessedDownwardMessages uint32 `scale:"5"`
	// The mark which specifies the block number up to which all inbound HRMP messages are processed.
	HrmpWatermark uint32 `scale:"6"`
}

// An assumption being made about the state of an occupied core.
type OccupiedCoreAssumption scale.VaryingDataType

// Set will set a VaryingDataTypeValue using the underlying VaryingDataType
func (o *OccupiedCoreAssumption) Set(val scale.VaryingDataTypeValue) (err error) {
	// cast to VaryingDataType to use VaryingDataType.Set method
	vdt := scale.VaryingDataType(*o)
	err = vdt.Set(val)
	if err != nil {
		return fmt.Errorf("setting value to varying data type: %w", err)
	}
	// store original ParentVDT with VaryingDataType that has been set
	*o = OccupiedCoreAssumption(vdt)
	return nil
}

// Value will return value from underying VaryingDataType
func (o *OccupiedCoreAssumption) Value() (scale.VaryingDataTypeValue, error) {
	vdt := scale.VaryingDataType(*o)
	return vdt.Value()
}

// Included means the candidate occupying the core was made available and included to free the core.
type Included scale.VaryingDataType //skipcq

// Index returns VDT index
func (Included) Index() uint { //skipcq
	return 0
}

func (Included) String() string { //skipcq
	return "Included"
}

// TimedOut means the candidate occupying the core timed out and freed the core without advancing the para.
type TimedOut scale.VaryingDataType //skipcq

// Index returns VDT index
func (TimedOut) Index() uint { //skipcq
	return 1
}

func (TimedOut) String() string { //skipcq
	return "TimedOut"
}

// Free means the core was not occupied to begin with.
type Free scale.VaryingDataType //skipcq

// Index returns VDT index
func (Free) Index() uint { //skipcq
	return 2
}

func (Free) String() string { //skipcq
	return "Free"
}

// The validation data provides information about how to create the inputs for validation of a candidate.
// This information is derived from the chain state and will vary from para to para, although some
// fields may be the same for every para.
//
// Since this data is used to form inputs to the validation function, it needs to be persisted by the
// availability system to avoid dependence on availability of the relay-chain state.
//
// Furthermore, the validation data acts as a way to authorize the additional data the collator needs
// to pass to the validation function. For example, the validation function can check whether the incoming
// messages (e.g. downward messages) were actually sent by using the data provided in the validation data
// using so called MQC heads.
//
// Since the commitments of the validation function are checked by the relay-chain, secondary checkers
// can rely on the invariant that the relay-chain only includes para-blocks for which these checks have
// already been done. As such, there is no need for the validation data used to inform validators and
// collators about the checks the relay-chain will perform to be persisted by the availability system.
//
// The `PersistedValidationData` should be relatively lightweight primarily because it is constructed
// during inclusion for each candidate and therefore lies on the critical path of inclusion.
type PersistedValidationData struct {
	ParentHead             headData
	RelayParentNumber      uint32
	RelayParentStorageRoot types.Header
	MaxPovSize             uint32
}
