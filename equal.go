package vdf

import (
	"reflect"
)

func (n *Node) Equal(r *Node) bool {
	n = n.notNil()
	r = r.notNil()

	if n.name != r.name {
		return false
	}

	if !reflect.DeepEqual(n.value, r.value) {
		return false
	}

	for nc, rc := n.child, r.child; nc != nil && rc != nil; nc, rc = nc.next, rc.next {
		if !nc.Equal(rc) {
			return false
		}
	}

	return true
}
