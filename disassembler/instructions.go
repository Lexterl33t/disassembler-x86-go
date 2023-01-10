package dissasembler

import (
	"fmt"
)

func (p *Pointer) push() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}
	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) add() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	p.NextOpcode()

	rm_mod := ParseRMOD(p.Opcodes[p.IP])

	rm_mod_opcode, ok := rm_mod_32_table[int(rm_mod.Mod)]

	if !ok {
		panic("Unknow mode")
		return
	}

	mod, ok := rm_mod_opcode[int(p.Opcodes[p.IP])]

	if !ok {
		panic("Unknow rm mode byte")
		return
	}

	dbit := ParseDbit(opcode.Mnemonic)

	if dbit == 1 {
		mod.Effective, mod.Value = mod.Value, mod.Effective
	}

	if !mod.SIB || !mod.Disp {
		if opcode.Size == 8 {
			mod.Effective = registers_8bits[mod.Effective]
			mod.Value = registers_8bits[mod.Value]
		}

		fmt.Println(opcode.Denomination, mod.Effective, ",", mod.Value)
	}

	p.NextOpcode()
}

func (p *Pointer) hlt() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}
	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) pop() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) es() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) daa() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) cs() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) das() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) ss() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) aaa() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) ds() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) aas() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) inc() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) dec() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) fs() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) gs() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) insb() {
	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) insw() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) outsb() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) outsw() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) nop() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) xchg() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) cbw() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) cwd() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) wait() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) lahf() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) sahf() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) movsb() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) movsw() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) cmpsb() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) cmpsw() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) stosb() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) stosw() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) lodsb() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) lodsw() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}
	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) scasb() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) scasw() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) retn() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) leave() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) retf() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) int3() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) into() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) iret() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) xlat() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) in() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) out() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) lock() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) repne() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}
	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) rep() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) htl() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}
	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) cmc() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) clc() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}
	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) stc() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) cli() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) sti() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) cld() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}

func (p *Pointer) std() {
	opcode, ok := opcodes_table_32[p.Opcodes[p.IP]]

	if !ok {
		panic("Unknow opcode")
		return
	}

	if !opcode.RmMod {
		fmt.Println(opcode.Denomination)
		p.NextOpcode()
		return
	}

	fmt.Println("Hello world")
	p.NextOpcode()
}
