package vdf_test

import (
	"os"
	"testing"

	"github.com/pkierski/vdf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBinaryReEncode(t *testing.T) {
	in, err := os.ReadFile("testdata/UserGameStatsSchema_630.bin")
	require.NoError(t, err)

	var n vdf.Node
	err = n.UnmarshalBinary(in)
	require.NoError(t, err)

	out, err := n.MarshalBinary()
	require.NoError(t, err)

	assert.Equal(t, in, out)
}
