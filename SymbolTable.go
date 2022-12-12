package main

var symbolTable map[string]int

func init() {
	symbolTable = make(map[string]int)
	symbolTable["R0"] = 0
	symbolTable["R1"] = 1
	symbolTable["R2"] = 2
	symbolTable["R3"] = 3
	symbolTable["R4"] = 4
	symbolTable["R5"] = 5
	symbolTable["R6"] = 6
	symbolTable["R7"] = 7
	symbolTable["R8"] = 8
	symbolTable["R9"] = 9
	symbolTable["R10"] = 10
	symbolTable["R11"] = 11
	symbolTable["R12"] = 12
	symbolTable["R13"] = 13
	symbolTable["R14"] = 14
	symbolTable["R15"] = 15
	symbolTable["SCREEN"] = 16384
	symbolTable["KBD"] = 24576
	symbolTable["SP"] = 0
	symbolTable["LCL"] = 1
	symbolTable["ARG"] = 2
	symbolTable["THIS"] = 3
	symbolTable["THAT"] = 4
}

// Adds <symbol, address> to the table
func AddEntry(symbol string, address int) {
	symbolTable[symbol] = address
}

// Returns true if symbol is contained in the symbol table
func Contains(symbol string) bool {
	if _, exists := symbolTable[symbol]; exists {
		return true
	}
	return false
}

// Returns the address assoicated with the symbol
func GetAddress(symbol string) int {
	return symbolTable[symbol]
}
