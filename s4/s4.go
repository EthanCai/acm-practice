package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")
	matrixNum, _ := strconv.Atoi(text)

	i := 0
	for i < matrixNum {
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n")
		matrixRows, _ := strconv.Atoi(text)

		matrix := readMatrix(reader, matrixRows)
		// fmt.Printf("matrix is %v\n", matrix)

		key := findKey(matrix)
		if key == nil || key.value == "" {
			fmt.Printf("0 0\n")
		} else {
			fmt.Printf("%v %v\n", key.row+1, key.column+1)
		}

		i++
	}
}

type Matrix struct {
	goNode *Node

	rows    int
	columns int

	nodes [][]*Node
}

type Node struct {
	value string

	row    int // 从0开始
	column int // 从0开始

	in  []*Node
	out *Node
}

const (
	valueOfGoNode = "GO"
	splitChars    = " "
)

func readMatrix(reader *bufio.Reader, matrixRows int) *Matrix {
	matrix := &Matrix{
		nodes: make([][]*Node, matrixRows),
		rows:  matrixRows,
	}

	// load input
	for i := 0; i < matrixRows; i++ {
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n")

		nodeValues := strings.Split(text, splitChars)
		row := make([]*Node, len(nodeValues))
		matrix.columns = len(nodeValues)

		for j, value := range nodeValues {
			newNode := &Node{
				value:  value,
				row:    i,
				column: j,
			}

			row[j] = newNode

			if valueOfGoNode == value {
				matrix.goNode = newNode
			}
		}

		matrix.nodes[i] = row
	}

	// build edges
	for i := 0; i < matrix.rows; i++ {
		for j := 0; j < matrix.columns; j++ {

			current := matrix.nodes[i][j]

			direction := current.value[0:1]
			steps, _ := strconv.Atoi(current.value[1:2])

			target := &Node{}

			switch direction {
			case "U":
				target = matrix.nodes[i-steps][j]
			case "D":
				target = matrix.nodes[i+steps][j]
			case "R":
				target = matrix.nodes[i][j+steps]
			case "L":
				target = matrix.nodes[i][j-steps]
			}

			current.out = target
			if target.in != nil {
				target.in = append(target.in, current)
			} else {
				target.in = []*Node{current}
			}
		}
	}

	return matrix
}

func findKey(matrix *Matrix) *Node {

	if len(matrix.goNode.in) <= 0 {
		return nil
	}

	maskMatrix := make([][]bool, matrix.rows)
	// init mask matrix
	for i := 0; i < matrix.rows; i++ {
		maskMatrix[i] = make([]bool, matrix.columns)
	}

	var keyNode *Node
	checkNode(maskMatrix, matrix.goNode, &keyNode)

	// fmt.Printf("keyNode is %v %v %v\n", keyNode.row, keyNode.column, keyNode.value)
	return keyNode
}

// keyNode 必须是指针的指针
func checkNode(maskMatrix [][]bool, currentNode *Node, keyNode **Node) {

	// fmt.Printf("check node %v %v %v\n", currentNode.row, currentNode.column, currentNode.value)

	if maskMatrix[currentNode.row][currentNode.column] {
		return
	}

	maskMatrix[currentNode.row][currentNode.column] = true
	if checkMaskMatrix(maskMatrix) {
		*keyNode = currentNode

		// fmt.Printf("keyNode is %v %v %v\n", currentNode.row, currentNode.column, currentNode.value)
		return
	}

	if len(currentNode.in) <= 0 {
		return
	}

	for _, n := range currentNode.in {
		checkNode(maskMatrix, n, keyNode)
	}
}

func checkMaskMatrix(maskMatrix [][]bool) bool {
	rows := len(maskMatrix)

	for i := 0; i < rows; i++ {

		columns := len(maskMatrix[i])

		for j := 0; j < columns; j++ {
			if maskMatrix[i][j] == false {
				return false
			}
		}
	}

	return true
}
