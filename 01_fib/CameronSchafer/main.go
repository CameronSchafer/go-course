package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//base fibonacci function
func fib(n int){
	sequence := calculateNormalFib(7)
	printFibSequence(sequence)
}

//calculates the normal fibonacci sequence 
//returns the sequence as an integer array.
func calculateNormalFib(x int) []int{
	var calcd_sequence []int
	//starting values.
	n1 := 0
	n2 := 1
	
	//loop until end of sequence.
	for count := 0; count < x; count++{
		calcd_sequence = append(calcd_sequence, n2)	//store n2 into array
		n1,n2 = n2, calcNextInSequence(n1,n2)	//calc next value in the sequence + assign n1 to the old value of n2.
	}

	return calcd_sequence
}

func calculateNegaFib(x int) []int{
	var calcd_sequence []int
	return calcd_sequence
}

//function calculates the next number in the fibonacci sequence.
func calcNextInSequence(n1 int, n2 int) int{
	nextInSequence := n1 + n2
	return nextInSequence
}

//loop through and print the sequence.
func printFibSequence(sequence []int){
	for _, num := range sequence {
        fmt.Fprintln(out, num)
	}
}

func main() {
	fib(7)
}
