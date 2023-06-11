// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package rpc

import (
	"fmt"

	"github.com/dojimanetwork/gossamer/dot/rpc/modules"
	"github.com/dojimanetwork/gossamer/dot/types"
	"github.com/dojimanetwork/gossamer/lib/common"
)

// headerResponseToHeader converts a *ChainBlockHeaderResponse to a *types.Header
func headerResponseToHeader(rpcHeader modules.ChainBlockHeaderResponse) (header *types.Header, err error) {
	parentHash, err := common.HexToHash(rpcHeader.ParentHash)
	if err != nil {
		return nil, fmt.Errorf("malformed parent hash: %w", err)
	}

	nb, err := common.HexToBytes(rpcHeader.Number)
	if err != nil {
		return nil, fmt.Errorf("malformed number hex string: %w", err)
	}

	number := common.BytesToUint(nb)

	stateRoot, err := common.HexToHash(rpcHeader.StateRoot)
	if err != nil {
		return nil, fmt.Errorf("malformed state root: %w", err)
	}

	extrinsicsRoot, err := common.HexToHash(rpcHeader.ExtrinsicsRoot)
	if err != nil {
		return nil, fmt.Errorf("malformed extrinsic root: %w", err)
	}

	digest, err := rpcLogsToDigest(rpcHeader.Digest.Logs)
	if err != nil {
		return nil, fmt.Errorf("malformed digest logs: %w", err)
	}

	header = types.NewHeader(parentHash, stateRoot, extrinsicsRoot, number, digest)
	return header, nil
}
