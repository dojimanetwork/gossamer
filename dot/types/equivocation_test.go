// Copyright 2022 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package types

import (
	"testing"

	"github.com/dojimanetwork/gossamer/lib/common"
	"github.com/dojimanetwork/gossamer/pkg/scale"
	"github.com/stretchr/testify/require"
)

func TestEquivocationProof(t *testing.T) {
	// To get these bytes run
	// https://github.com/paritytech/substrate/blob/17c07af0b953b84dbe89341294e98e586f9b4591/frame/babe/src/tests.rs#L932
	expectedEncoding := common.MustHexToBytes("0xdef12e42f3e487e9b14095aa8d5cc16a33491f1b50dadcf8811d1480f3fa86270b0000000000000043fd935464ab466417a2d3b51b750b3047acc94708aa8e69bb01d19e7ba841f428cee631e4d752a4de8130431b63246d695dcc87af881316251bc6d35651f9508a03170a2e7597b7b7e3d84c05391d139a62b157e78786d8c082f29dcf4c1113140c06424142453402000000000a0000000000000004424142456902010cdef12e42f3e487e9b14095aa8d5cc16a33491f1b50dadcf8811d1480f3fa862701000000000000003a3d45dc55b57bf542f4c6ff41af080ec675317f4ed50ae1d2713bf9f892692d010000000000000054c71c235773b82115f0744252369c13414fd0e8bad3e8feff462c6a4bb58a0f0100000000000000c6e9d02ce38de7b255382f804a64f9bc74aad5597f51fde6bb53c0b8a76c22ba054241424501015881750a61f36303470033d7a9c4d5654ee4d11983ba73008cbe4af8e0361e62b1e67b58236a4258f17ceed53d11e204528238a412eab6ce3476e9d3eb42c18143fd935464ab466417a2d3b51b750b3047acc94708aa8e69bb01d19e7ba841f428cee631e4d752a4de8130431b63246d695dcc87af881316251bc6d35651f9508a03170a2e7597b7b7e3d84c05391d139a62b157e78786d8c082f29dcf4c1113140c06424142453402000000000a0000000000000004424142456902010cdef12e42f3e487e9b14095aa8d5cc16a33491f1b50dadcf8811d1480f3fa862701000000000000003a3d45dc55b57bf542f4c6ff41af080ec675317f4ed50ae1d2713bf9f892692d010000000000000054c71c235773b82115f0744252369c13414fd0e8bad3e8feff462c6a4bb58a0f0100000000000000c6e9d02ce38de7b255382f804a64f9bc74aad5597f51fde6bb53c0b8a76c22ba05424142450101e4df6a034d5057b1eace2dd4918f2357c5ab0413615596ebee5129fb0fcf146a087c8b3b65d55f76ebf91a77504e334be9b6a36cb836adf58cfd1756b149b689") //nolint:lll

	decodedProof := BabeEquivocationProof{
		FirstHeader:  *NewEmptyHeader(),
		SecondHeader: *NewEmptyHeader(),
	}

	err := scale.Unmarshal(expectedEncoding, &decodedProof)
	require.NoError(t, err)

	actualEncoding, err := scale.Marshal(decodedProof)
	require.NoError(t, err)
	require.Equal(t, expectedEncoding, actualEncoding)
}
