package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	//init
	fileName := os.Args[1]
	parser := CreateParser(fileName)

	//first pass
	lineCounter := 0
	for parser.Advance() {
		if parser.InstructionType() == L_INSTRUCTION {
			AddEntry(parser.Symbol(), lineCounter) //Adds new label to the symbol table
		} else { // increments lineCounter if the instruction is not L_INSTRUCTION
			lineCounter++
		}
	}

	//second pass
	outputFile, err := os.Create(strings.Split(fileName, ".asm")[0] + ".hack")
	//checks if there was an error while creating the output file
	if err != nil {
		panic(err)
	}
	parser.setInputFile(fileName) //reset the scanner to beginning of the input file
	symbolAddress := 16
	for parser.Advance() {
		binaryLine := ""
		switch parser.InstructionType() {
		case A_INSTRUCTION:
			{
				symbol, err := strconv.Atoi(parser.Symbol()) // checks if the symbol is a number or a variable name
				if err != nil {                              //symbol is a variable symbol i.e "sum"
					if !Contains(parser.Symbol()) { // the symbol doesn't exist in SymbolTable
						AddEntry(parser.Symbol(), symbolAddress)
						binaryLine += intToBinary(symbolAddress) //convert the symbol address into binary
						symbolAddress++
					} else { // the symbol exits in SymbolTable
						binaryLine += intToBinary(GetAddress(parser.Symbol())) //convert the symbol address into binary
					}
				} else {
					binaryLine += intToBinary(symbol) //convert the symbol number into binary
				}
				addBinaryLineToFile(outputFile, binaryLine)
			}
		case C_INSTRUCTION:
			{
				binaryLine += "111" + Comp(parser.Comp()) + Dest(parser.Dest()) + Jump(parser.Jump())
				addBinaryLineToFile(outputFile, binaryLine)
			}
		}
	}
}

// Adds leading 0's the 16-bit binary line (if needed) and appends it to the output Prog.hack file
func addBinaryLineToFile(outputFile *os.File, binaryLine string) {
	linePrefix := strings.Repeat("0", 16-len(binaryLine))
	binaryLine = linePrefix + binaryLine
	outputFile.WriteString(binaryLine + "\n")
}

// Converts num into its binary representation
func intToBinary(num int) string {
	return strconv.FormatInt(int64(num), 2)
}
