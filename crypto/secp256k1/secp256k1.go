package secp256k1

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"

	"github.com/ChainSafe/gossamer/common"
	"github.com/ChainSafe/gossamer/crypto"
	secp256k1 "github.com/ethereum/go-ethereum/crypto"
)

// PrivateKeyLength is the fixed Private Key Length
const PrivateKeyLength = 32

// SignatureLength is the fixed Signature Length
const SignatureLength = 64

// MessageLength is the fixed Message Length
const MessageLength = 32

// Keypair holds the pub,pk keys
type Keypair struct {
	public  *PublicKey
	private *PrivateKey
}

// PublicKey struct for PublicKey
type PublicKey struct {
	key ecdsa.PublicKey
}

// PrivateKey struct for PrivateKey
type PrivateKey struct {
	key ecdsa.PrivateKey
}

// NewKeypair will returned a Keypair from a PrivateKey
func NewKeypair(pk ecdsa.PrivateKey) *Keypair {
	pub := pk.Public()

	return &Keypair{
		public:  &PublicKey{key: *pub.(*ecdsa.PublicKey)},
		private: &PrivateKey{key: pk},
	}
}

// NewKeypairFromPrivate will return a Keypair for a PrivateKey
func NewKeypairFromPrivate(priv *PrivateKey) (*Keypair, error) {
	pub, err := priv.Public()
	if err != nil {
		return nil, err
	}

	return &Keypair{
		public:  pub.(*PublicKey),
		private: priv,
	}, nil
}

// NewPrivateKey will return a PrivateKey for a []byte
func NewPrivateKey(in []byte) (*PrivateKey, error) {
	if len(in) != PrivateKeyLength {
		return nil, errors.New("input to create secp256k1 private key is not 32 bytes")
	}
	priv := new(PrivateKey)
	err := priv.Decode(in)
	return priv, err
}

// GenerateKeypair will generate a Keypair
func GenerateKeypair() (*Keypair, error) {
	priv, err := secp256k1.GenerateKey()
	if err != nil {
		return nil, err
	}

	return NewKeypair(*priv), nil
}

// Sign will sign
func (kp *Keypair) Sign(msg []byte) ([]byte, error) {
	if len(msg) != MessageLength {
		return nil, errors.New("invalid message length: not 32 byte hash")
	}

	return secp256k1.Sign(msg, &kp.private.key)
}

// Public returns the pub key
func (kp *Keypair) Public() crypto.PublicKey {
	return kp.public
}

// Private returns pk
func (kp *Keypair) Private() crypto.PrivateKey {
	return kp.private
}

// Verify a msg
func (k *PublicKey) Verify(msg, sig []byte) (bool, error) {
	if len(sig) != SignatureLength {
		return false, errors.New("invalid signature length")
	}

	if len(msg) != MessageLength {
		return false, errors.New("invalid message length: not 32 byte hash")
	}

	return secp256k1.VerifySignature(k.Encode(), msg, sig), nil
}

// Encode will encode to []byte
func (k *PublicKey) Encode() []byte {
	return secp256k1.CompressPubkey(&k.key)
}

// Decode will decode to PublicKey key field
func (k *PublicKey) Decode(in []byte) error {
	pub, err := secp256k1.DecompressPubkey(in)
	if err != nil {
		return err
	}
	k.key = *pub
	return nil
}

// Address will return PublicKey Address
func (k *PublicKey) Address() common.Address {
	return crypto.PublicKeyToAddress(k)
}

// Hex will return PublicKey Hex
func (k *PublicKey) Hex() string {
	enc := k.Encode()
	h := hex.EncodeToString(enc)
	return "0x" + h
}

// Sign a message
func (pk *PrivateKey) Sign(msg []byte) ([]byte, error) {
	if len(msg) != MessageLength {
		return nil, errors.New("invalid message length: not 32 byte hash")
	}

	return secp256k1.Sign(msg, &pk.key)
}

// Public will return pub key
func (pk *PrivateKey) Public() (crypto.PublicKey, error) {
	return &PublicKey{
		key: *(pk.key.Public().(*ecdsa.PublicKey)),
	}, nil
}

// Encode will encode
func (pk *PrivateKey) Encode() []byte {
	return secp256k1.FromECDSA(&pk.key)
}

// Decode will decode
func (pk *PrivateKey) Decode(in []byte) error {
	key := secp256k1.ToECDSAUnsafe(in)
	pk.key = *key
	return nil
}
