package main

import (
	"bufio"
	"os"
	"strings"
)

const A_INSTRUCTION int = 0
const C_INSTRUCTION int = 1
const L_INSTRUCTION int = 2

type Parser struct {
	scanner                  *bufio.Scanner
	currentInstruction       string
	currentInstructionType   int
	currentInstructionSymbol string
}

// Opens the input file and gets ready to parse it using Scanner
func CreateParser(fileName string) *Parser {
	p := &Parser{}
	p.setInputFile(fileName)
	return p
}

// Initializes the scanner with the input file
func (p *Parser) setInputFile(fileName string) {
	f, err := os.Open(fileName)
	checkErr(err)
	scanner := bufio.NewScanner(f)
	p.scanner = scanner
}

// Returns true if the input file has more lines to read
func (p *Parser) HasMoreLines() bool {
	return p.scanner.Scan()
}

/*
Advances the scanner to the next readable line and
parses the line according to its instruction type

Returns true if a line was successfully parsed
*/
func (p *Parser) Advance() bool {
	for p.HasMoreLines() {
		line := strings.TrimSpace(p.scanner.Text())
		if len(p.scanner.Text()) != 0 && !strings.HasPrefix(line, "/") { // skip lines that are empty or comments
			line = strings.Split(line, "/")[0] // remove comments from the line
			line = strings.TrimSpace(line)
			p.currentInstruction = line
			switch firstChar := line[0]; firstChar {
			case '@':
				{
					p.currentInstructionType = A_INSTRUCTION
					p.currentInstructionSymbol = line[1:]
				}
			case '(':
				{
					p.currentInstructionType = L_INSTRUCTION
					p.currentInstructionSymbol = line[1 : len(line)-1]
				}
			default:
				{
					p.currentInstructionType = C_INSTRUCTION
					p.currentInstructionSymbol = "nil"
				}
			}
			return true
		}
	}
	return false
}

// Returns the type of the current instruction
func (p *Parser) InstructionType() int {
	return p.currentInstructionType
}

/*
If the current instruction is (xxx), returns the symbol xxx.
If the current instruction @xxx, returns the symbol or decimal xxx (as a string)
*/
func (p *Parser) Symbol() string {
	return p.currentInstructionSymbol
}

// Returns the symbolic dest part of the current C-instruction
func (p *Parser) Dest() string {
	s := strings.Split(p.currentInstruction, "=")[0]
	if s != p.currentInstruction {
		return s
	} else { // implying '=' is not contained in the line
		return "nil"
	}
}

// Returns the symbolic comp part of the current C-instruction
func (p *Parser) Comp() string {
	s := strings.Split(p.currentInstruction, ";")[0]
	if strings.Contains(s, "=") {
		return strings.Split(s, "=")[1]
	}
	return s
}

// Returns the symbolic jump part of the current C-instruction
func (p *Parser) Jump() string {
	if strings.Contains(p.currentInstruction, ";") {
		return strings.Split(p.currentInstruction, ";")[1]
	}
	return "nil"
}

// Terminates the program if there was an error while opening the input file
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
