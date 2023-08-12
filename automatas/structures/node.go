package structures

import (
	"fmt"
	"strings"
)

/*
NodeGroup is a type that define a group (as slice) that the nodes that are
pointed by the node owner of group.
*/
type nodeGroup []*node

/*
Node are the minimal structures to implements connections in the automata, these
represent to the states.
*/
type node struct {
	name  	     string
	terminalNode bool
	adjacents 	 map[uint8]nodeGroup
}

/*
Thhis node function add a node to the adjacents group using a character as key to get
the group. If the connection didn't exist It will return true.
*/
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

/*
This node function return the result of apply a character input.
*/
func (n node) applyInput(character uint8) (nodeGroup, bool) {
	data, ok := n.adjacents[character]
	return data, ok
}

// This function is used to create nodes.
func newNode(name string) *node {
	nodeInstance := &node {
		name: name,
		adjacents: make(map[uint8]nodeGroup, 10),
	}

	return nodeInstance
}

// Semantical function used to compare if two nodes are equals (They have the same name).
func EqualsNodes(nodeA, nodeB *node) bool {
	return nodeA.name == nodeB.name
}

func (n node) String() string {
	return n.name
}

// Convert the info of the node into []byte to write it on a .graph file.
func (n node) toBytes() []byte {
	var sb strings.Builder
	for char, nodeGroup := range n.adjacents {
		for _, node := range nodeGroup {
			sb.WriteString(fmt.Sprintf("%v-%c-%v\n", n, char, node))
		}
	}
	return []byte(sb.String())
}