package grandpa

import (
	"github.com/ChainSafe/gossamer/dot/types"
	"github.com/ChainSafe/gossamer/lib/keystore"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestApi struct {
	genesisAuthorities []types.Authority
}

func NewTestApi(genesisAuthorities []types.Authority) TestApi {
	return TestApi{genesisAuthorities: genesisAuthorities}
}

type GrandpaPeer struct {
	// TODO fill
	// type GrandpaPeer = Peer<PeerData, GrandpaBlockImport>;
}

type GrandpaTestNet struct {
	peers      []GrandpaPeer
	testConfig TestApi
}

func NewGrandpaTestNet(testConfig TestApi, numAuthorities, numFull int) GrandpaTestNet {
	net := GrandpaTestNet{
		peers:      make([]GrandpaPeer, 0, numAuthorities+numFull),
		testConfig: testConfig,
	}
	for i := 0; i < numAuthorities; i++ {
		net.addAuthorityPeer()
	}

	for i := 0; i < numFull; i++ {
		net.addFullPeer()
	}

	return net
}

func (gtn GrandpaTestNet) addAuthorityPeer() {
	// TODO impl
}

func (gtn GrandpaTestNet) addFullPeer() {
	// TODO impl
}

func makeIds(t *testing.T, peers []keystore.KeyPair) (authorityList []types.Authority) {
	t.Helper()
	for i, peer := range peers {
		authorityList[i] = types.Authority{
			Key:    peer.Public(),
			Weight: 1,
		}
	}
	return authorityList
}

func TestFinalize3VotersNoObservers(t *testing.T) {
	ks, err := keystore.NewEd25519Keyring()
	require.NoError(t, err)
	peers := []keystore.KeyPair{ks.Alice(), ks.Bob(), ks.Charlie()}
	voters := makeIds(t, peers)

	testApi := NewTestApi(voters)
	_ = NewGrandpaTestNet(testApi, 3, 0)
}
