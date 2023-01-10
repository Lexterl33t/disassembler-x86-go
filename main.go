package main

import (
	dissasembler "github.com/Lexterl33t/disassembler-x86-go/disassembler"
)

func main() {
	opcodes := []byte{
		0x00, 0xD8,
	}
	dissasembler.Disas32(opcodes)
}
