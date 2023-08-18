package main

import (
	"fmt"
	"automatas/structures"
	"strings"
	"os"
	"bufio"
)

var graph = structures.NewGraph() // The graph used to make the automatas

func main() {
	op := 0

	fmt.Println("This program was created by Luigi Salcedo,")
	fmt.Println("The source code is on GitHub.")
	fmt.Println()
	fmt.Println("GitHub -> Luigi Salcedo | Twitter -> @LuigiSalcedo96")
	fmt.Println()

	for op != 3 {

		fmt.Println("Automatas Builder\n")

		fmt.Println("[1]. Create a new automata.")
		fmt.Println("[2]. Load from file.")

		fmt.Println()

		fmt.Println("[3]. Exit.")

		fmt.Println()

		fmt.Print("Select an option: ")
		fmt.Scan(&op)

		switch op {
		case 1:
			createNodes()
			mainMenu()
		case 2:
			err := loadFromFile()
			if err == nil {
				mainMenu()
			} else {
				fmt.Println("\n", err, "\n")
			}
		case 3:
		default:
			fmt.Println("\nOption is not valid.\n")
		}
	}
}

func createNodes() {
	n := 0

	graph = structures.NewGraph()

	for n <= 0 {
		fmt.Print("\nEnter the number of nodes in the graph: ")
		fmt.Scan(&n)

		if n <= 0 {
			fmt.Println("\nThe number of nodes is not valid.")
		}
	}

	sliceTemp := make([]string, 0, n)

	for i := 0; i < n; i++ {
		sliceTemp = append(sliceTemp, fmt.Sprintf("q%d", i))
		graph.AddNode(sliceTemp[i])
	}

	fmt.Println("\nThe next nodes (states) have been created ->", sliceTemp)
}

func loadFromFile() error {
	fmt.Println("\nLoading from file\n")

	fmt.Println("Rember that the automatas files must be in the folder 'file' . . . \n")

	fileName := ""

	fmt.Print("Write de file name: ")
	fmt.Scan(&fileName)

	fileName = strings.Trim(fileName, " ")

	if !strings.HasSuffix(fileName, ".graph") {
		fileName += ".graph"
	}

	file, err := os.Open("files\\" + fileName)

	if err != nil {
		return err
	}

	graph = structures.NewGraph()

	return graph.LoadDataFromFile(file)
}

func mainMenu() {
	op := 0

	for op != 11 {
		fmt.Println("\nMain menu\n")

		fmt.Println("[1]. Add a new node.")
		fmt.Println("[2]. Modify a node.")
		fmt.Println("[3]. Create connection.")
		fmt.Println("[4]. Remove a connection.")
		fmt.Println("[5]. See the transitions table.")
		fmt.Println("[6]. Set initial state.")
		fmt.Println("[7]. Set a terminal state.")
		fmt.Println("[8]. Apply an input.")
		fmt.Println("[9]. Save in a file.")
		fmt.Println("[10]. Read inputs from file.")
		fmt.Println()
		fmt.Println("[11]. Exit.")
		fmt.Println()
		fmt.Print("Select an option: ")
		fmt.Scan(&op)

		switch op {
		case 1:
			addNode()
		case 2:
			modifyNode()
		case 3:
			createConnection() 
		case 4:
			removeConnection()
		case 5:
			transitionsTable()
		case 6:
			setInitialState()
		case 7:
			setTerminalState()
		case 8:
			applyInput()
		case 9:
			saveFile()
		case 10:
			readInputs() 
		case 11:
		default:
			fmt.Println("\nOption is not valid . . .")
		}
	}
}

func addNode() {
	fmt.Println("\nAdd a node\n")

	name := ""
	
	fmt.Print("New node name: ")	
	fmt.Scan(&name)

	name = strings.Trim(name, " ")

	ok := graph.AddNode(name)

	if !ok {
		fmt.Println("\n[Error]: Already exists a node using this name . . . ")
		return
	}

	fmt.Println("\n[OK]: Node added to the graph . . . ")
}

func modifyNode() {
	fmt.Println("\nModify a node\n")

	currentName := ""
	newName := ""

	fmt.Print("Write the current node name: ")
	fmt.Scan(&currentName)

	currentName = strings.Trim(currentName, " ")

	fmt.Print("Write the new node name: ")
	fmt.Scan(&newName)

	err := graph.UpdateNode(currentName, newName)

	if err == nil {
		fmt.Println("\nName has been updated from:", currentName, "to:", newName)
		return
	}

	fmt.Printf("\n[Error]: %v\n", err)
}

func createConnection() {
	fmt.Println("\nCreating a connection\n")

	var chars string

	sourceName := ""
	destinationName := ""

	fmt.Print("Write the source node name: ")
	fmt.Scan(&sourceName)

	sourceName = strings.Trim(sourceName, " ")

	_, ok := graph.Nodes[sourceName]

	if !ok {
		fmt.Printf("[Error]:There is no nodes named \"%v\"\n", sourceName)
		return
	}

	fmt.Print("Write the destination node name: ")
	fmt.Scan(&destinationName)

	destinationName = strings.Trim(destinationName, " ")

	_, ok = graph.Nodes[destinationName]

	if !ok {
		fmt.Printf("[Error]:There is no nodes named \"%v\"\n", destinationName)
		return
	}

	fmt.Print("Write the character (transition) for this connection: ")	
	fmt.Scan(&chars)

	for _, char := range chars {
		ok = graph.CreateConnection(sourceName, destinationName, uint8(char))

		if !ok {
			fmt.Printf("\n[Warning]: Maybe already exists a connection like %v - %c - %v . . . \n", sourceName, char, destinationName)
		}
	}

	fmt.Println("\nConnection has been created . . . ")
}


func removeConnection() {
	fmt.Println("\nModify a connection\n")

	sourceName := ""
	destinationName := ""

	var char string

	fmt.Print("Write the source node name of the connection: ")
	fmt.Scan(&sourceName)

	fmt.Print("Write the destination node name of the connection: ")
	fmt.Scan(&destinationName)

	fmt.Print("Write the character of the connection: ")
	fmt.Scan(&char)

	err := graph.RemoveConnection(sourceName, destinationName, char[0])

	if err != nil {
		fmt.Println("\n[Error]:", err)
		return
	}

	fmt.Println("\n[OK]: Connection has been removed . . . ")

}

func setInitialState() {
	fmt.Println("\nSetting the initial state\n")

	name := ""

	fmt.Print("Write the node that will be the initial: ")
	fmt.Scan(&name)

	name = strings.Trim(name, " ")

	ok := graph.SetInitialNode(name)

	if !ok {
		fmt.Println("\n[Error]: Please, verify the name of the initial node . . . ")
		return 
	}

	fmt.Println("\nInitial state defined. The graph is ready to get inputs . . .")
}

func setTerminalState() {
	fmt.Println("\nSetting a terminal state\n")

	name := ""

	fmt.Print("Write the node that will be a terminal state: ")

	fmt.Scan(&name)

	name = strings.Trim(name, " ")

	ok := graph.SetTerminalNode(name)

	if !ok {
		fmt.Println("\n[Error]: Please, verify the name of the terminal node . . . ")
		return
	}

	fmt.Println("\nTerminal state defined . . .")
}

func applyInput() {

	if graph.InitialNode == nil {
		fmt.Println("[Error]: You have to set an initial state if you want to do this . . .")
		return
	}

	fmt.Println("\nApply an input\n")

	input := ""

	fmt.Print("Please, write the input: ")
	fmt.Scan(&input)

	output := graph.ApplyInput(input)

	fmt.Printf("\nInput: %s\nOutput: %v\n", input, output)
}

func transitionsTable() {
	fmt.Println("\n\n")
	for _, char := range graph.Alphabet {
		fmt.Printf("\t\t%c", char)
	}
	fmt.Println()
	for node, _ := range graph.Nodes {
		fmt.Print(node)
		for _,char := range graph.Alphabet {
			fmt.Print("\t\t", graph.GetAdjacents(node, char))
		}
		fmt.Println()
	}
	fmt.Println("\n")
}

func saveFile() {
	fmt.Println("\nSaving in a file.\n")
	
	fileName := ""

	fmt.Print("Write the file name to save this automata: ")
	fmt.Scan(&fileName)

	fileName = strings.Trim(fileName, " ")

	if !strings.HasSuffix(fileName, ".graph") {
		fileName += ".graph"
	}

	file, err := os.Create("files\\" + fileName)

	if os.IsExist(err) {
		op := 0
		fmt.Println("\nAlready exists a file named like this . . .\n")
		for op != 1 && op != 2 {
			fmt.Print("Would you like override this file? [1. Yes - 2. No]: ")
			fmt.Scan(&op)
		}

		if op == 2 {
			fmt.Println("\nThe automata was not saved . . .")
			return
		}

		file, err = os.Open("files\\" + fileName)
	} 

	if err != nil {
		fmt.Println("\n[Error]:", err)
		return
	}

	fileData := graph.ToBytes()

	writer := bufio.NewWriter(file)
	writer.Write([]byte(fileData))
	writer.Flush()
	file.Close()
	fmt.Println("\n[OK]: The automata was saved . . .")
}

func readInputs() {
	fmt.Println("\nReading inputs from file\n")
	fmt.Println("The inputs have to be in 'inputs' like an .txt file . . . ")

	fileName := ""

	fmt.Print("\nWrite the file name: ")
	fmt.Scan(&fileName)

	fileName = strings.Trim(fileName, " ")

	if !strings.HasSuffix(fileName, ".txt") {
		fileName += ".txt"
	}

	file, err := os.Open("inputs/" + fileName)

	if err != nil {
		fmt.Printf("\n[ERROR]:%v\n", err)
		return 
	}

	scanner := bufio.NewScanner(file)

	data := make(chan string, 10)

	go readFile(scanner, data)

	for input := range data {
		fmt.Printf("\nInput: %v -> Output: %v\n", input, graph.ApplyInput(input))
	}
	file.Close()
}

func readFile(sc *bufio.Scanner, data chan<- string) {
	defer close(data)

	for sc.Scan(){ 
		data <- string(sc.Bytes())
	}
}
