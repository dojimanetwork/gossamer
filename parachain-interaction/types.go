package parachaininteraction

import (
	"fmt"

	"github.com/ChainSafe/gossamer/lib/common"
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
	NewValidationCode *validationCode `scale:"3"`
	// The head-data produced as a result of execution.
	HeadData headData `scale:"4"`
	// The number of messages processed from the DMQ.
	ProcessedDownwardMessages uint32 `scale:"5"`
	// The mark which specifies the block number up to which all inbound HRMP messages are processed.
	HrmpWatermark uint32 `scale:"6"`
}
