package vdf

func (n *Node) Clone() *Node {
	return n.cloneInternal(false)
}

func (n *Node) cloneInternal(withNext bool) *Node {
	if n == nil {
		return nil
	}

	c := &Node{
		condition: n.condition,
		name:      n.name,
		value:     n.value,
	}

	if withNext {
		c.next = n.next.cloneInternal(true)
	}

	if n.child == nil {
		return c
	}

	var cprev *Node
	dstChild := &c.child
	for srcChild := n.child; srcChild != nil; srcChild = srcChild.next {
		*dstChild = srcChild.cloneInternal(true)
		(*dstChild).parent = c
		(*dstChild).prev = cprev
		cprev = *dstChild
		dstChild = &cprev.next
	}

	return c
}
