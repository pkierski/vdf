package vdf_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pkierski/vdf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestText(t *testing.T) {
	for _, name := range []string{
		"hello",
		"hello_eof",
		"hello_quotes",
		"cond",
	} {
		name := name // shadow

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			in, err := os.ReadFile(filepath.Join("testdata", "in_"+name+".txt"))
			require.NoError(t, err)
			out, err := os.ReadFile(filepath.Join("testdata", "out_"+name+".txt"))
			require.NoError(t, err)

			var n vdf.Node
			err = n.UnmarshalText(in)
			require.NoError(t, err)

			out1, err := n.MarshalText()
			require.NoError(t, err)

			assert.Equal(t, in, out1)

			n.ClearFormatting()

			out2, err := n.MarshalText()
			require.NoError(t, err)
			assert.Equal(t, out, out2)
		})
	}
}
