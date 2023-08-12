package structures

import (
	"fmt"
	"strings"
)

type Type uint8

type nodeGroup []*node

type node struct {
	name  	     string
	terminalNode bool
	adjacents 	 map[uint8]nodeGroup
}

func (n node) addAdjacent(node *node, character uint8) bool {
	listAdjacents, ok := n.applyInput(character)

	if ok {
		for _, adjacent := range listAdjacents {
			if adjacent.name == node.name {
				return false
			}
		}
	}
	
	n.adjacents[character] = append(listAdjacents, node)
	return true
}

func (n node) applyInput(character uint8) (nodeGroup, bool) {
	data, ok := n.adjacents[character]
	return data, ok
}

func newNode(name string) *node {
	nodeInstance := &node {
		name: name,
		adjacents: make(map[uint8]nodeGroup, 10),
	}

	return nodeInstance
}

func EqualsNodes(nodeA, nodeB *node) bool {
	return nodeA.name == nodeB.name
}

func (n node) String() string {
	return n.name
}

func (n node) toBytes() []byte {
	var sb strings.Builder
	for char, nodeGroup := range n.adjacents {
		for _, node := range nodeGroup {
			sb.WriteString(fmt.Sprintf("%v-%c-%v\n", n, char, node))
		}
	}
	return []byte(sb.String())
}