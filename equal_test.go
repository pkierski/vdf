package vdf_test

import (
	"fmt"
	"testing"

	"github.com/pkierski/vdf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEqualSpecialCases(t *testing.T) {

	emptyA := &vdf.Node{}
	emptyB := &vdf.Node{}

	for _, c := range []struct {
		a, b *vdf.Node
	}{
		{nil, nil},
		{nil, emptyA},
		{emptyA, nil},
		{emptyA, emptyA},
		{emptyA, emptyB},
	} {
		assert.True(t, c.a.Equal(c.b))
	}
}

func TestEqual(t *testing.T) {
	for _, c := range []struct {
		a, b string
	}{
		{`a b`, `a b`},
		{`a {b c}`, `a {b c}`},
		{`a {b {d e} c f}`, `a {b {d e} c f}`},
	} {
		t.Run(fmt.Sprintf("%v != %v", c.a, c.b), func(t *testing.T) {
			na := &vdf.Node{}
			err := na.UnmarshalText([]byte(c.a))
			require.NoError(t, err)
			nb := &vdf.Node{}
			err = nb.UnmarshalText([]byte(c.b))
			require.NoError(t, err)
			assert.True(t, na.Equal(nb))
		})
	}

}

func TestNotEqual(t *testing.T) {
	for _, c := range []struct {
		a, b string
	}{
		{`a1 b`, `a b`},
		{`a b1`, `a b`},
		{`a {b c1}`, `a {b c}`},
		{`a {b {d1 e} c f}`, `a {b {d e} c f}`},
		{`a {b {d e1} c f}`, `a {b {d e} c f}`},
	} {
		t.Run(fmt.Sprintf("%v != %v", c.a, c.b), func(t *testing.T) {
			na := &vdf.Node{}
			err := na.UnmarshalText([]byte(c.a))
			require.NoError(t, err)
			nb := &vdf.Node{}
			err = nb.UnmarshalText([]byte(c.b))
			require.NoError(t, err)
			assert.False(t, na.Equal(nb))
		})
	}

}
