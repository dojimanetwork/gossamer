// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package stress

import (
	"context"
	"testing"
	"time"

	libutils "github.com/dojimanetwork/gossamer/lib/utils"
	"github.com/dojimanetwork/gossamer/tests/utils"
	"github.com/dojimanetwork/gossamer/tests/utils/config"
	"github.com/dojimanetwork/gossamer/tests/utils/node"
	"github.com/dojimanetwork/gossamer/tests/utils/rpc"

	"github.com/dojimanetwork/gossamer/internal/log"
	"github.com/stretchr/testify/require"
)

func TestNetwork_MaxPeers(t *testing.T) {
	numNodes := 9 // 9 block producers
	genesisPath := libutils.GetWestendDevRawGenesisPath(t)
	utils.Logger.Patch(log.SetLevel(log.Info))
	tomlConfig := config.Default()
	tomlConfig.ChainSpec = genesisPath
	nodes := node.MakeNodes(t, numNodes, tomlConfig)
	ctx, cancel := context.WithCancel(context.Background())
	nodes.InitAndStartTest(ctx, t, cancel)

	// wait for nodes to connect
	time.Sleep(time.Second * 10)

	for i, node := range nodes {
		const getPeersTimeout = time.Second
		getPeersCtx, getPeersCancel := context.WithTimeout(ctx, getPeersTimeout)
		peers, err := rpc.GetPeers(getPeersCtx, node.RPCPort())
		getPeersCancel()

		require.NoError(t, err)

		t.Logf("node %d: peer count=%d", i, len(peers))
		require.LessOrEqual(t, len(peers), 5)
	}
}
