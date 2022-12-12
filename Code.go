package main

var destMap map[string]string
var jumpMap map[string]string
var compMap map[string]string

func init() {
	destMap = make(map[string]string)
	jumpMap = make(map[string]string)
	compMap = make(map[string]string)

	//Initiazlize destMap
	destMap["nil"] = "000"
	destMap["M"] = "001"
	destMap["D"] = "010"
	destMap["DM"] = "011"
	destMap["MD"] = "011"
	destMap["A"] = "100"
	destMap["AM"] = "101"
	destMap["AD"] = "110"
	destMap["ADM"] = "111"

	//Initialize jumpMap
	jumpMap["nil"] = "000"
	jumpMap["JGT"] = "001"
	jumpMap["JEQ"] = "010"
	jumpMap["JGE"] = "011"
	jumpMap["JLT"] = "100"
	jumpMap["JNE"] = "101"
	jumpMap["JLE"] = "110"
	jumpMap["JMP"] = "111"

	//Initialize compMap
	compMap["0"] = "0101010"
	compMap["1"] = "0111111"
	compMap["-1"] = "0111010"
	compMap["D"] = "0001100"
	compMap["A"] = "0110000"
	compMap["M"] = "1110000"
	compMap["!D"] = "0001101"
	compMap["!A"] = "0110001"
	compMap["!M"] = "1110001"
	compMap["-D"] = "0001111"
	compMap["-A"] = "0110011"
	compMap["-M"] = "1110011"
	compMap["D+1"] = "0011111"
	compMap["A+1"] = "0110111"
	compMap["M+1"] = "1110111"
	compMap["D-1"] = "0001110"
	compMap["A-1"] = "0110010"
	compMap["M-1"] = "1110010"
	compMap["D+A"] = "0000010"
	compMap["D+M"] = "1000010"
	compMap["D-A"] = "0010011"
	compMap["D-M"] = "1010011"
	compMap["A-D"] = "0000111"
	compMap["M-D"] = "1000111"
	compMap["D&A"] = "0000000"
	compMap["D&M"] = "1000000"
	compMap["D|A"] = "0010101"
	compMap["D|M"] = "1010101"
}

// Returns the binary code of the dest mnemonic
func Dest(dest string) string {
	return destMap[dest]
}

// Returns the binary code of the comp mnemonic
func Comp(comp string) string {
	return compMap[comp]
}

// Returns the binary code of the jump mnemonic
func Jump(jump string) string {
	return jumpMap[jump]
}
