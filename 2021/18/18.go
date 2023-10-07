package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

const NIL_NUMBER = -1

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	value  int
}

func main() {
	input := utils.ReadFile(2021, 18, "\n")

	fmt.Println("2021 Day 18")
	fmt.Println("\tPart 1:", calculate1(input))
	fmt.Println("\tPart 2:", calculate2(input))
}

func calculate1(input []string) int {
	result := parseString(input[0])

	for _, line := range input[1:] {
		result = sumNodes(result, parseString(line))
	}

	return calculateMagnitude(result)
}

func calculate2(input []string) int {
	max := 0

	for i1, line1 := range input {
		for i2, line2 := range input {
			if i1 != i2 {
				max = utils.Max(max, calculateMagnitude(sumNodes(parseString(line1), parseString(line2))))
			}
		}
	}

	return max
}

func parseString(line string) *Node {
	root := &Node{parent: nil, value: NIL_NUMBER}
	previousNode := root
	leftSide := true

	for _, char := range line[1:] {
		switch char {
		case '[':
			if leftSide {
				previousNode.left = &Node{parent: previousNode, value: NIL_NUMBER}
				previousNode = previousNode.left
			} else {
				previousNode.right = &Node{parent: previousNode, value: NIL_NUMBER}
				previousNode = previousNode.right
			}
			leftSide = true
		case ']':
			previousNode = previousNode.parent
		case ',':
			leftSide = false
		default:
			if leftSide {
				previousNode.left = &Node{parent: previousNode, value: utils.StringToInt(string(char))}
			} else {
				previousNode.right = &Node{parent: previousNode, value: utils.StringToInt(string(char))}
			}
		}
	}
	return root
}

func sumNodes(node1, node2 *Node) *Node {
	newNode := &Node{
		parent: nil,
		left:   node1,
		right:  node2,
		value:  NIL_NUMBER,
	}
	node1.parent = newNode
	node2.parent = newNode

	return reduce(newNode)
}

func reduce(node *Node) *Node {
	if nodeToExplode := getNodeAtDepth(node, 4); nodeToExplode != nil {
		explode(nodeToExplode)
		return reduce(node)
	} else if split(node) {
		return reduce(node)
	}
	return node
}

func getNodeAtDepth(node *Node, depth int) *Node {
	if depth == 0 && node.value == NIL_NUMBER {
		return node
	} else if node.value == NIL_NUMBER {
		if n := getNodeAtDepth(node.left, depth-1); n != nil {
			return n
		}
		if n := getNodeAtDepth(node.right, depth-1); n != nil {
			return n
		}
	}
	return nil
}

func explode(node *Node) {
	addLeft(node, node.left.value)
	addRight(node, node.right.value)

	node.left = nil
	node.right = nil
	node.value = 0
}

func addLeft(node *Node, number int) {
	if node.parent != nil {
		if node.parent.left.value != NIL_NUMBER {
			node.parent.left.value += number
		} else if node.parent.left != node {
			recursivelyAddToRightChild(node.parent.left, number)
		} else {
			addLeft(node.parent, number)
		}
	}
}

func recursivelyAddToRightChild(node *Node, number int) {
	if node.value != NIL_NUMBER {
		node.value += number
	} else {
		recursivelyAddToRightChild(node.right, number)
	}
}

func addRight(node *Node, number int) {
	if node.parent != nil {
		if node.parent.right.value != NIL_NUMBER {
			node.parent.right.value += number
		} else if node.parent.right != node {
			recursivelyAddToLeftChild(node.parent.right, number)
		} else {
			addRight(node.parent, number)
		}
	}
}

func recursivelyAddToLeftChild(node *Node, number int) {
	if node.value != NIL_NUMBER {
		node.value += number
	} else {
		recursivelyAddToLeftChild(node.left, number)
	}
}

func split(node *Node) bool {
	if node.value == NIL_NUMBER {
		return split(node.left) || split(node.right)
	} else if node.value > 9 {
		node.left = &Node{
			parent: node,
			value:  node.value / 2,
		}
		node.right = &Node{
			parent: node,
			value:  (node.value + 1) / 2,
		}
		node.value = NIL_NUMBER
		return true
	}
	return false
}

func calculateMagnitude(node *Node) int {
	if node.value >= 0 {
		return node.value
	}
	return 3*calculateMagnitude(node.left) + 2*calculateMagnitude(node.right)
}
