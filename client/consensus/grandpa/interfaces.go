package grandpa

import (
	"encoding/json"
	"github.com/ChainSafe/gossamer/dot/network"
	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

// ForkTree A tree data structure that stores several nodes across multiple branches.
//
// Top-level branches are called roots. The tree has functionality for
// finalizing nodes, which means that node is traversed, and all competing
// branches are pruned. It also guarantees that nodes in the tree are finalized
// in order. Each node is uniquely identified by its hash but can be ordered by
// its number. In order to build the tree an external function must be provided
// when interacting with the tree to establish a node's ancestry.
type ForkTree interface {
	Import(hash common.Hash, number uint, change PendingChange, isDescendentOf IsDescendentOf) error
}

// Network is the interface required by GRANDPA for the network
type Network interface {
	GossipMessage(msg network.NotificationsMessage)
	SendMessage(to peer.ID, msg network.NotificationsMessage) error
	RegisterNotificationsProtocol(sub protocol.ID,
		messageID byte,
		handshakeGetter network.HandshakeGetter,
		handshakeDecoder network.HandshakeDecoder,
		handshakeValidator network.HandshakeValidator,
		messageDecoder network.MessageDecoder,
		messageHandler network.NotificationsMessageHandler,
		batchHandler network.NotificationsMessageBatchHandler,
		maxSize uint64,
	) error
}

// Telemetry is the telemetry client to send telemetry messages.
type Telemetry interface {
	SendMessage(msg json.Marshaler)
}
