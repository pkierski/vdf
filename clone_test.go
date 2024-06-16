package vdf_test

import (
	"testing"

	"github.com/pkierski/vdf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClone(t *testing.T) {
	str := `foo {
	a va
	c {
		d vd
		e {
			f vf
		}
		g {
			h vh
		}
		i vi
	}
	e ve
	}`
	n := vdf.Node{}
	err := n.UnmarshalText([]byte(str))
	require.NoError(t, err)

	toClone := n.FirstByName("c")
	cloned := toClone.Clone()

	assert.True(t, cloned.Equal(toClone))
}
