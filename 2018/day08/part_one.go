package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type node struct {
	index                     int
	quantityOfChildNodes      int
	quantityOfMetadataEntries int
	childNodes                []node
	metadata                  []int
}

// calculateSumOfMetadataEntries calculates the
// sum of all metadata entries of this node
// and all child nodes of this node.
func calculateSumOfMetadataEntries(node node) int {
	// Initial sum is 0
	sum := 0

	// Iterate over the child nodes of this node
	for _, v := range node.childNodes {
		// For each child, calculate the
		// sum of metadata entries and
		// add it to sum.
		sum = sum + calculateSumOfMetadataEntries(v)
	}

	// Iterate over the metadata of
	// this node and add its value
	// to sum.
	for _, v := range node.metadata {
		sum = sum + v
	}

	return sum
}

// Constructs a node where index is the
// index in splittedData of the first
// integer of this node.
func constructNode(index int, splittedData []int) node {
	// A node always consist of:
	// - A header, which is always exactly two integers:
	//   - Quantity of child nodes
	//   - Quantity of metadata entries
	// - Zero or more child nodes (as specified in the
	//	 header)
	// - One or more metadata entries (as specified in
	//   the header)

	// We start with creating a node with at least
	// the header. Then we know if there are child
	// nodes and metadata entries.
	node := node{
		index:                     index,
		quantityOfChildNodes:      splittedData[index],
		quantityOfMetadataEntries: splittedData[index+1]}

	offset := node.index + 2

	// If this node has child nodes,
	// we will apply some logic.
	for i := 0; i < node.quantityOfChildNodes; i++ {
		// Construct a node of this child node
		childNode := constructNode(offset, splittedData)

		// Append this child node to this node
		node.childNodes = append(node.childNodes, childNode)

		// Increment offset with the length
		// of this child node.
		offset = offset + calculateNodesLength(childNode)
	}

	// If this node has metadata entries,
	// we append each metadata entry to
	// this node.
	for i := 0; i < node.quantityOfMetadataEntries; i++ {
		node.metadata = append(node.metadata, splittedData[offset+i])
	}

	return node
}

// calculateNodesLength calculates the length of this node.
func calculateNodesLength(node node) int {
	// The length of each node must at least be 2,
	// because the header is always exactly two
	// integers.
	length := 2

	// If this node has child nodes, we will
	// calculate the length of each child
	// node and add it to length.
	for i := 0; i < node.quantityOfChildNodes; i++ {
		length = length + calculateNodesLength(node.childNodes[i])
	}

	// Also add the quantity of metadata
	// entries to length.
	length = length + node.quantityOfMetadataEntries

	return length
}

func main() {
	// Slurp the entire content of "input.txt"
	// into our memory.
	inputData, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// input.txt should look like a huge
	// string with integers, seperated
	// with a " ". For example:
	//
	// 9 1 2 1 2 1 1 1 3 3 1 3 1 3 4

	// Slice of all substrings in inputData
	// seperated by a " ".
	//
	// Note that inputData is of type []byte,
	// so we must first convert it to string.
	splittedRawData := strings.Split(string(inputData), " ")

	// splittedRawData only contains strings,
	// but we need ints. Instead of later
	// converting each string to an int,
	// we will do it immediatly and
	// store everything in this
	// slice.
	var splittedData []int

	// Convert every element in splittedRawData
	// to a string and append it to
	// splittedData.
	for _, v := range splittedRawData {
		i, err := strconv.Atoi(v)

		if err != nil {
			log.Fatal(err)
		}

		splittedData = append(splittedData, i)
	}

	rootNode := constructNode(0, splittedData)

	// Sum of all metadata entries.
	// This will be our final answer.
	sum := calculateSumOfMetadataEntries(rootNode)

	// Print the final answer
	fmt.Println("The sum of all metadata entries is", sum)
}
