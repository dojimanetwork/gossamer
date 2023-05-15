package grandpa

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDelayKind(t *testing.T) {
	finalizedKind := Finalized{}
	delayKind := NewDelayKind(finalizedKind)
	_, isFinalizedType := delayKind.value.(Finalized)
	require.True(t, isFinalizedType)

	medLastFinalized := uint(3)
	bestKind := Best{medianLastFinalized: medLastFinalized}
	delayKind = NewDelayKind(bestKind)
	best, isBestType := delayKind.value.(Best)
	require.True(t, isBestType)
	require.Equal(t, medLastFinalized, best.medianLastFinalized)
}
