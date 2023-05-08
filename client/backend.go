package client

// TODO prob delete, just a temp file to record data on that client backend

/// Client backend.
///
/// Manages the data layer.
///
/// # State Pruning
///
/// While an object from `state_at` is alive, the state
/// should not be pruned. The backend should internally reference-count
/// its state objects.
///
/// The same applies for live `BlockImportOperation`s: while an import operation building on a
/// parent `P` is alive, the state for `P` should not be pruned.
///
/// # Block Pruning
///
/// Users can pin blocks in memory by calling `pin_block`. When
/// a block would be pruned, its value is kept in an in-memory cache
/// until it is unpinned via `unpin_block`.
///
/// While a block is pinned, its state is also preserved.
///
/// The backend should internally reference count the number of pin / unpin calls.
