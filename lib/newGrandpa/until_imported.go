// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package newGrandpa

// BlockUntilImported Something that needs to be withheld until specific blocks are available.
//
// For example a GRANDPA commit message which is not of any use without the corresponding block
// that it commits on.
type BlockUntilImported struct {
	// this is actually a trait
}

// Check if a new incoming item needs awaiting until a block(s) is imported.
func (s *BlockUntilImported) needs_waiting() {}

// called when the wait has completed. The canonical number is passed through
// for further checks.
func (s *BlockUntilImported) wait_completed() {}

// DiscardWaitOrReady Describes whether a given [`BlockUntilImported`] (a) should be discarded, (b) is waiting for
// specific blocks to be imported or (c) is ready to be used.
//
// A reason for discarding a [`BlockUntilImported`] would be if a referenced block is perceived
// under a different number than specified in the message.
type DiscardWaitOrReady struct {
	// this is an enum
}

// Metrics2 Prometheus metrics for the `UntilImported` queue.
// At a given point in time there can be more than one `UntilImported` queue. One can not register a
// metric twice, thus queues need to share the same Prometheus metrics instead of instantiating
// their own ones.
//
// When a queue is dropped it might still contain messages. In order for those to not distort the
// Prometheus metrics, the `Metric` struct cleans up after itself within its `Drop` implementation
// by subtracting the local_waiting_messages (the amount of messages left in the queue about to
// be dropped) from the global_waiting_messages gauge.
type Metrics2 struct {
	//global_waiting_messages: Gauge<U64>,
	//local_waiting_messages: u64,
}

// UntilImported Buffering incoming messages until blocks with given hashes are imported.
type UntilImported struct {
	//import_notifications: Fuse<TracingUnboundedReceiver<BlockImportNotification<Block>>>,
	//block_sync_requester: BlockSyncRequester,
	//status_check: BlockStatus,
	//incoming_messages: Fuse<I>,
	//ready: VecDeque<M::Blocked>,
	///// Interval at which to check status of each awaited block.
	//check_pending: Pin<Box<dyn Stream<Item = Result<(), std::io::Error>> + Send>>,
	///// Mapping block hashes to their block number, the point in time it was
	///// first encountered (Instant) and a list of GRANDPA messages referencing
	///// the block hash.
	//pending: HashMap<Block::Hash, (NumberFor<Block>, Instant, Vec<M>)>,
	//
	///// Queue identifier for differentiation in logs.
	//identifier: &'static str,
	///// Prometheus metrics.
	//metrics: Option<Metrics>,
}

// Create a new `UntilImported` wrapper.
func (s *UntilImported) new() {}

func (s *UntilImported) poll_next() {}

func warn_authority_wrong_target() {}

func (s *BlockUntilImported) needs_waiting2() {}

func (s *BlockUntilImported) wait_completed2() {}

// UntilVoteTargetImported Helper type definition for the stream which waits until vote targets for
// signed messages are imported.
type UntilVoteTargetImported struct{}

// BlockGlobalMessage This blocks a global message import, i.e. a commit or catch up messages,
// until all blocks referenced in its votes are known.
//
// This is used for compact commits and catch up messages which have already
// been checked for structural soundness (e.g. valid signatures).
//
// We use the `Arc`'s reference count to implicitly count the number of outstanding blocks that we
// are waiting on for the same message (i.e. other `BlockGlobalMessage` instances with the same
// `inner`).
type BlockGlobalMessage struct{}

func (s *BlockGlobalMessage) needs_waiting() {}

func (s *BlockGlobalMessage) wait_completed() {}

// UntilGlobalMessageBlocksImported A stream which gates off incoming global messages, i.e. commit and catch up
// messages, until all referenced block hashes have been imported.
type UntilGlobalMessageBlocksImported struct{}
