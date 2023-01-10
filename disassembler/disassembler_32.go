package dissasembler

import ()

/*

Direct address. The instruction has no ModR/M byte. The address
 of the operand is coded in the instruction: no base register,
 index register, or scaling factor can be applied.
 for ex : far JMP(EA)

*/

type Pointer struct {
	Opcodes []byte
	IP      int
}

func NewPointer(opcodes []byte) *Pointer {
	return &Pointer{
		Opcodes: opcodes,
		IP:      0,
	}
}

func ParseRMOD(rm_byte byte) Mod_RM {

	var mod_rm Mod_RM

	mod_rm.Mod = ((1 << 2) - 1) & (rm_byte >> (7 - 1))
	mod_rm.Reg = ((1 << 3) - 1) & (rm_byte >> (3 - 1))
	mod_rm.RM = ((1 << 3) - 1) & (rm_byte >> (1 - 1))

	return mod_rm
}

func ParseDbit(opcode byte) byte {
	return ((opcode >> 6) & 0xF) << 7
}

func (p *Pointer) NextOpcode() {
	p.IP++
}

func Disas32(opcodes []byte) {

	var size_opcode int = len(opcodes)

	var pointer *Pointer = NewPointer(opcodes)

	LoadTable(pointer)

	for pointer.IP < size_opcode {
		if opcode, ok := opcodes_table_32[pointer.Opcodes[pointer.IP]]; ok {
			opcode.DecodeFunc()
		} else {
			panic("Unknow opcode")
		}
	}

}
