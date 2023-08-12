package structures

import (
	"errors"
	"slices"
	"strings"
	"fmt"
	"os"
	"bufio"
)


// This type represent if the input is valid or not
type InputResult bool

func (ir InputResult) String() string {
	if bool(ir) == false {
		return "Rejected"
	}

	return "Accepted"
}

/*
	Graph: struct

	This is the graph representation of an automata like a graph. 

	InitialNode -> Represent the initial state

	Nodes -> States group. The node names are using like key to get them.
*/
type Graph struct {
	InitialNode    *node
	Nodes 		   map[string]*node
	Alphabet	   []uint8
}

func (graph *Graph) SetInitialNode(name string) bool {
	node, ok := graph.Nodes[name]

	if !ok {
		return false
	}

	graph.InitialNode = node
	return true
}

func (graph *Graph) SetTerminalNode(name string) bool {
	node, ok := graph.Nodes[name]

	if !ok {
		return false
	}

	node.terminalNode = true
	return true
}

func (graph *Graph) AddNode(name string) bool {
	node, ok := graph.Nodes[name]

	if ok {
		return false
	}

	node = newNode(name)
	graph.Nodes[name] = node
	return true
}

// This function is used by graphs to creat an arc between two nodes using a character
func (graph *Graph) CreateConnection(source, destination string, character uint8) bool {
	nodeSource, ok := graph.Nodes[source]

	if !ok {
		return false
	}

	nodeDestination, ok := graph.Nodes[destination]

	if !ok {
		return false
	}

	ok = nodeSource.addAdjacent(nodeDestination, character)

	if ok {
		if !slices.Contains(graph.Alphabet, character) {
			graph.Alphabet = append(graph.Alphabet, character)
		}
		return true
	}
	return false
	
}

func (graph *Graph) RemoveConnection(sourceName, destinationName string, char uint8) error {
	sourceNode, ok := graph.Nodes[sourceName]

	if !ok {
		return errors.New("There is not a source node named " + sourceName + " in this graph . . .")
	}

	destinationNode, ok := graph.Nodes[destinationName]

	if !ok {
		return errors.New("There is not a destination node named " + destinationName + " in this graph . . .")
	}

	arcDestination, ok := sourceNode.adjacents[char]

	if !ok {
		return errors.New("The node "+ sourceNode.name + " doesn't have an arc with that character . . .")
	}

	ok = false

	index := 0

	for i, adj := range arcDestination {
		if EqualsNodes(adj, destinationNode) {
			ok = true
			index = i
			break
		}
	}

	if !ok {
		return errors.New("The node " + sourceNode.name + " doesn't have a connection to " + destinationNode.name + " using that char . . . ")
	}

	sourceNode.adjacents[char] = append(arcDestination[:index], arcDestination[index+1:]...)

	return nil
}

// This funcion change de name of a node
func (graph *Graph) UpdateNode(currentName, newName string) error {
	nodeSource, ok := graph.Nodes[currentName]

	if !ok {
		return errors.New("No nodes with the current name: " + currentName)
	}

	_, ok = graph.Nodes[newName]

	if ok {
		return errors.New("Already exist a node with the new name: " + newName)
	}

	nodeSource.name = newName
	return nil
}

// This function is used to get the result with a given input
func (graph *Graph) ApplyInput(input string) InputResult {
	states := make([]*node, 0, 25)
	states = append(states, graph.InitialNode)

	for _, character := range []uint8(input) {
		nextStates := make([]*node, 0, 25)

		for _, node := range states {
			nodeAdjacents, ok := node.applyInput(character)

			if !ok {
				break
			}
			nextStates = append(nextStates, nodeAdjacents...)
		}

		if len(nextStates) == 0 {
			return InputResult(false)
		}

		states = make([]*node, 0, len(nextStates))
		states = append(states, nextStates...)
	}

	for _, state := range states {
		if state.terminalNode {
			return InputResult(true)
		}
	}

	return InputResult(false)
}

// This function is used to get the adjacents of a node
func (graph *Graph) GetAdjacents(name string, char uint8) nodeGroup {
	node, ok := graph.Nodes[name]

	if !ok {
		return nodeGroup{}
	}

	return node.adjacents[char]
}

// This function convert the data on memory of the graph in []byte to write it in a file
func (graph *Graph) ToBytes() string {
	var sb strings.Builder
	for name, _ := range graph.Nodes {
		sb.WriteString(name + ",")
	}
	sb.WriteString("\n")
	for _, char := range graph.Alphabet {
		sb.WriteString(fmt.Sprintf("%c", char))
	}
	for _, node := range graph.Nodes {
		sb.WriteString("\n")
		sb.Write(node.toBytes())
	} 
	if graph.InitialNode != nil {
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf(">%v", graph.InitialNode.String()))
	}
	for _, node := range graph.Nodes {
		if node.terminalNode {
			sb.WriteString("\n")
			sb.WriteString("*" + node.name)
		}
	}
	return sb.String()
}

// This function is used to get the graph data from file
func (graph *Graph) LoadDataFromFile(file *os.File) error {
	defer func() error {
		file.Close()
		if err := recover(); err != nil {
			return errors.New(fmt.Sprintf("Someting went wrong -> %v", err))
		}
		return nil
	}()
	
	reader := bufio.NewScanner(file)

	reader.Scan()

	nodesNames := strings.Split(string(reader.Bytes()), ",")

	if len(nodesNames) < 1{
		panic("It's look like a not valid .graph file . . . ")
	}

	for i := 0; i < len(nodesNames)-1; i++ {
		graph.AddNode(nodesNames[i])
	}

	fmt.Println("\nNodes loaded:", nodesNames)

	reader.Scan()

	alphabetChars := strings.Split(string(reader.Bytes()), "")

	graph.Alphabet = make([]uint8, 0, len(alphabetChars))

	for _, char := range alphabetChars {
		graph.Alphabet = append(graph.Alphabet, char[0])
	}

	fmt.Println("Alphabet loaded:", graph.Alphabet)

	var ok bool

	for reader.Scan() {
		line := string(reader.Bytes())
		
		// When a connection is detected
		if strings.Contains(line, "-") {
			connection := strings.Split(line, "-")
			ok = graph.CreateConnection(connection[0], connection[2], connection[1][0]) 
			if !ok {
				panic("The node " + line + " is not valid in this graph . . .")
			}
			fmt.Printf("Connection created: %v - %c -> %v\n", connection[0], byte(connection[1][0]), connection[2])
			continue
		}

		// When a configuration is detected
		if strings.Contains(line, ">") {
			ok = graph.SetInitialNode(strings.Trim(line, ">"))
			if !ok {
				panic("The node " + line + " is not valid in this graph . . .")
			}
			fmt.Println("Initial state detected:", line)
			continue
		}

		if strings.Contains(line, "*") {
			ok = graph.SetTerminalNode(strings.Trim(line, "*"))
			if !ok {
				panic("The node " + line + " is not valid in this graph . . .")
			}
			fmt.Println("Terminal state detected:", line)
		}
	}
	panic(nil)
}

// This function is used to create graph elements
func NewGraph() *Graph {
	graph := &Graph {
		InitialNode: nil,
		Nodes: make(map[string]*node),
		Alphabet: make([]uint8, 0, 10),
	}

	return graph
}