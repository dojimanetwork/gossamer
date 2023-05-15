package grandpa

// DelayedKinds Kinds of delays for pending changes.
type DelayedKinds interface {
	Finalized | Best
}

type DelayKind struct {
	value interface{}
}

func SetDelayKind[T DelayedKinds](delayKind *DelayKind, val T) {
	delayKind.value = val
}

func NewDelayKind[T DelayedKinds](val T) DelayKind {
	delayKind := DelayKind{}
	SetDelayKind(&delayKind, val)
	return delayKind
}

// Finalized Depth in finalized chain.
type Finalized struct{}

// Best Depth in best chain. The median last finalized block is calculated at the time the
// change was signaled.
type Best struct {
	medianLastFinalized uint
}
