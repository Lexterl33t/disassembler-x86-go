package dissasembler

type Opcode struct {
	Mnemonic     byte
	Size         uint8
	Denomination string
	RmMod        bool
	DecodeFunc   func()
}

type RM_MOD_Opcode struct {
	Effective string
	Value     string
	SIB       bool
	Disp      bool
}

type Mod_RM struct {
	Mod byte
	Reg byte
	RM  byte
}

type SIB struct {
	Scale int
	Index int
	Base  int
}

var opcodes_table_32 map[byte]Opcode

var registers_8bits map[string]string = map[string]string{
	"eax": "al",
	"ecx": "cl",
	"ebx": "bl",
	"edx": "dl",
	"edi": "di",
	"esi": "si",
}

var rm_mod_32_table map[int]map[int]RM_MOD_Opcode = map[int]map[int]RM_MOD_Opcode{
	0x00: {
		0x00: {Effective: "eax", Value: "eax", SIB: false, Disp: false},
		0x01: {Effective: "ecx", Value: "eax", SIB: false, Disp: false},
		0x02: {Effective: "edx", Value: "eax", SIB: false, Disp: false},
		0x03: {Effective: "ebx", Value: "eax", SIB: false, Disp: false},
		0x04: {Effective: "", Value: "eax", SIB: true, Disp: false},
		0x05: {Effective: "disp32", Value: "eax", SIB: false, Disp: true},
		0x06: {Effective: "esi", Value: "eax", SIB: false, Disp: false},
		0x07: {Effective: "edi", Value: "eax", SIB: false, Disp: false},
		0x08: {Effective: "eax", Value: "ecx", SIB: false, Disp: false},
		0x09: {Effective: "ecx", Value: "ecx", SIB: false, Disp: false},
		0x0A: {Effective: "edx", Value: "ecx", SIB: false, Disp: false},
		0x0B: {Effective: "ebx", Value: "ecx", SIB: false, Disp: false},
		0x0C: {Effective: "", Value: "ecx", SIB: true, Disp: false},
		0x0D: {Effective: "disp32", Value: "ecx", SIB: false, Disp: true},
		0x0E: {Effective: "esi", Value: "ecx", SIB: false, Disp: false},
		0x0F: {Effective: "edi", Value: "ecx", SIB: false, Disp: false},
		0x10: {Effective: "eax", Value: "edx", SIB: false, Disp: false},
		0x11: {Effective: "ecx", Value: "edx", SIB: false, Disp: false},
		0x12: {Effective: "edx", Value: "edx", SIB: false, Disp: false},
		0x13: {Effective: "ebx", Value: "edx", SIB: false, Disp: false},
		0x14: {Effective: "", Value: "edx", SIB: true, Disp: false},
		0x15: {Effective: "disp32", Value: "edx", SIB: false, Disp: true},
		0x16: {Effective: "esi", Value: "edx", SIB: false, Disp: false},
		0x17: {Effective: "edi", Value: "edx", SIB: false, Disp: false},
		0x18: {Effective: "eax", Value: "ebx", SIB: false, Disp: false},
		0x19: {Effective: "ecx", Value: "ebx", SIB: false, Disp: false},
		0x1A: {Effective: "edx", Value: "ebx", SIB: false, Disp: false},
		0x1B: {Effective: "ebx", Value: "ebx", SIB: false, Disp: false},
		0x1C: {Effective: "", Value: "ebx", SIB: true, Disp: false},
		0x1D: {Effective: "disp32", Value: "ebx", SIB: false, Disp: true},
		0x1E: {Effective: "esi", Value: "ebx", SIB: false, Disp: false},
		0x1F: {Effective: "edi", Value: "ebx", SIB: false, Disp: false},
		0x20: {Effective: "eax", Value: "esp", SIB: false, Disp: false},
		0x21: {Effective: "ecx", Value: "esp", SIB: false, Disp: false},
		0x22: {Effective: "edx", Value: "esp", SIB: false, Disp: false},
		0x23: {Effective: "ebx", Value: "esp", SIB: false, Disp: false},
		0x24: {Effective: "", Value: "esp", SIB: true, Disp: false},
		0x25: {Effective: "disp32", Value: "esp", SIB: false, Disp: true},
		0x26: {Effective: "esi", Value: "esp", SIB: false, Disp: false},
		0x27: {Effective: "edi", Value: "esp", SIB: false, Disp: false},
		0x28: {Effective: "eax", Value: "ebp", SIB: false, Disp: false},
		0x29: {Effective: "ecx", Value: "ebp", SIB: false, Disp: false},
		0x2A: {Effective: "edx", Value: "ebp", SIB: false, Disp: false},
		0x2B: {Effective: "ebx", Value: "ebp", SIB: false, Disp: false},
		0x2C: {Effective: "", Value: "ebp", SIB: true, Disp: false},
		0x2D: {Effective: "disp32", Value: "ebp", SIB: false, Disp: true},
		0x2E: {Effective: "esi", Value: "ebp", SIB: false, Disp: false},
		0x2F: {Effective: "edi", Value: "ebp", SIB: false, Disp: false},
		0x30: {Effective: "eax", Value: "esi", SIB: false, Disp: false},
		0x31: {Effective: "ecx", Value: "esi", SIB: false, Disp: false},
		0x32: {Effective: "edx", Value: "esi", SIB: false, Disp: false},
		0x33: {Effective: "ebx", Value: "esi", SIB: false, Disp: false},
		0x34: {Effective: "", Value: "esi", SIB: true, Disp: false},
		0x35: {Effective: "disp32", Value: "esi", SIB: false, Disp: true},
		0x36: {Effective: "esi", Value: "esi", SIB: false, Disp: false},
		0x37: {Effective: "edi", Value: "esi", SIB: false, Disp: false},
		0x38: {Effective: "eax", Value: "edi", SIB: false, Disp: false},
		0x39: {Effective: "ecx", Value: "edi", SIB: false, Disp: false},
		0x3A: {Effective: "edx", Value: "edi", SIB: false, Disp: false},
		0x3B: {Effective: "ebx", Value: "edi", SIB: false, Disp: false},
		0x3C: {Effective: "", Value: "edi", SIB: true, Disp: false},
		0x3D: {Effective: "disp32", Value: "edi", SIB: false, Disp: true},
		0x3E: {Effective: "esi", Value: "edi", SIB: false, Disp: false},
		0x3F: {Effective: "edi", Value: "edi", SIB: false, Disp: false},
	},
	0x01: {},
	0x02: {},
	0x03: {
		0xC0: {Effective: "eax", Value: ""},
	},
}

func LoadTable(p *Pointer) {

	opcodes_table_32 = map[byte]Opcode{
		0x06: {Mnemonic: 0x06, Size: 8, Denomination: "push es", RmMod: false, DecodeFunc: p.push},
		0x07: {Mnemonic: 0x07, Size: 8, Denomination: "pop es", RmMod: false, DecodeFunc: p.pop},
		0x0E: {Mnemonic: 0x0E, Size: 8, Denomination: "push cs", RmMod: false, DecodeFunc: p.push},
		0x16: {Mnemonic: 0x16, Size: 8, Denomination: "push ss", RmMod: false, DecodeFunc: p.push},
		0x17: {Mnemonic: 0x17, Size: 8, Denomination: "pop ss", RmMod: false, DecodeFunc: p.pop},
		0x1E: {Mnemonic: 0x1E, Size: 8, Denomination: "push ds", RmMod: false, DecodeFunc: p.push},
		0x1F: {Mnemonic: 0x1F, Size: 8, Denomination: "pop ds", RmMod: false, DecodeFunc: p.pop},
		0x26: {Mnemonic: 0x26, Size: 8, Denomination: "es:", RmMod: false, DecodeFunc: p.es},
		0x27: {Mnemonic: 0x27, Size: 8, Denomination: "daa", RmMod: false, DecodeFunc: p.daa},
		0x2E: {Mnemonic: 0x2E, Size: 8, Denomination: "cs:", RmMod: false, DecodeFunc: p.cs},
		0x2F: {Mnemonic: 0x2F, Size: 8, Denomination: "das", RmMod: false, DecodeFunc: p.das},
		0x36: {Mnemonic: 0x36, Size: 8, Denomination: "ss:", RmMod: false, DecodeFunc: p.ss},
		0x37: {Mnemonic: 0x37, Size: 8, Denomination: "aaa", RmMod: false, DecodeFunc: p.aaa},
		0x3E: {Mnemonic: 0x3E, Size: 8, Denomination: "ds:", RmMod: false, DecodeFunc: p.ds},
		0x3F: {Mnemonic: 0x3F, Size: 8, Denomination: "aas", RmMod: false, DecodeFunc: p.aas},
		0x40: {Mnemonic: 0x40, Size: 8, Denomination: "inc eax", RmMod: false, DecodeFunc: p.inc},
		0x41: {Mnemonic: 0x41, Size: 8, Denomination: "inc ecx", RmMod: false, DecodeFunc: p.inc},
		0x42: {Mnemonic: 0x42, Size: 8, Denomination: "inc edx", RmMod: false, DecodeFunc: p.inc},
		0x43: {Mnemonic: 0x43, Size: 8, Denomination: "inc ebx", RmMod: false, DecodeFunc: p.inc},
		0x44: {Mnemonic: 0x44, Size: 8, Denomination: "inc esp", RmMod: false, DecodeFunc: p.inc},
		0x45: {Mnemonic: 0x45, Size: 8, Denomination: "inc ebp", RmMod: false, DecodeFunc: p.inc},
		0x46: {Mnemonic: 0x46, Size: 8, Denomination: "inc esi", RmMod: false, DecodeFunc: p.inc},
		0x47: {Mnemonic: 0x47, Size: 8, Denomination: "inc edi", RmMod: false, DecodeFunc: p.inc},
		0x48: {Mnemonic: 0x48, Size: 8, Denomination: "dec eax", RmMod: false, DecodeFunc: p.dec},
		0x49: {Mnemonic: 0x49, Size: 8, Denomination: "dec ecx", RmMod: false, DecodeFunc: p.dec},
		0x4A: {Mnemonic: 0x4A, Size: 8, Denomination: "dec edx", RmMod: false, DecodeFunc: p.dec},
		0x4B: {Mnemonic: 0x4B, Size: 8, Denomination: "dec ebx", RmMod: false, DecodeFunc: p.dec},
		0x4C: {Mnemonic: 0x4C, Size: 8, Denomination: "dec esp", RmMod: false, DecodeFunc: p.dec},
		0x4D: {Mnemonic: 0x4D, Size: 8, Denomination: "dec ebp", RmMod: false, DecodeFunc: p.dec},
		0x4E: {Mnemonic: 0x4E, Size: 8, Denomination: "dec esi", RmMod: false, DecodeFunc: p.dec},
		0x4F: {Mnemonic: 0x4F, Size: 8, Denomination: "dec edi", RmMod: false, DecodeFunc: p.dec},
		0x50: {Mnemonic: 0x50, Size: 8, Denomination: "push eax", RmMod: false, DecodeFunc: p.push},
		0x51: {Mnemonic: 0x51, Size: 8, Denomination: "push ecx", RmMod: false, DecodeFunc: p.push},
		0x52: {Mnemonic: 0x52, Size: 8, Denomination: "push edx", RmMod: false, DecodeFunc: p.push},
		0x53: {Mnemonic: 0x53, Size: 8, Denomination: "push ebx", RmMod: false, DecodeFunc: p.push},
		0x54: {Mnemonic: 0x54, Size: 8, Denomination: "push esp", RmMod: false, DecodeFunc: p.push},
		0x55: {Mnemonic: 0x55, Size: 8, Denomination: "push ebp", RmMod: false, DecodeFunc: p.push},
		0x56: {Mnemonic: 0x56, Size: 8, Denomination: "push esi", RmMod: false, DecodeFunc: p.push},
		0x57: {Mnemonic: 0x57, Size: 8, Denomination: "push edi", RmMod: false, DecodeFunc: p.push},
		0x58: {Mnemonic: 0x58, Size: 8, Denomination: "pop eax", RmMod: false, DecodeFunc: p.pop},
		0x59: {Mnemonic: 0x59, Size: 32, Denomination: "pop ecx", RmMod: false, DecodeFunc: p.pop},
		0x5A: {Mnemonic: 0x5A, Size: 32, Denomination: "pop edx", RmMod: false, DecodeFunc: p.pop},
		0x5B: {Mnemonic: 0x5B, Size: 32, Denomination: "pop ebx", RmMod: false, DecodeFunc: p.pop},
		0x5C: {Mnemonic: 0x5C, Size: 32, Denomination: "pop esp", RmMod: false, DecodeFunc: p.pop},
		0x5D: {Mnemonic: 0x5D, Size: 32, Denomination: "pop ebp", RmMod: false, DecodeFunc: p.pop},
		0x5F: {Mnemonic: 0x5F, Size: 32, Denomination: "pop esi", RmMod: false, DecodeFunc: p.pop},
		0x60: {Mnemonic: 0x60, Size: 0, Denomination: "pusha", RmMod: false, DecodeFunc: p.push},
		0x61: {Mnemonic: 0x61, Size: 0, Denomination: "popa", RmMod: false, DecodeFunc: p.pop},
		0x64: {Mnemonic: 0x64, Size: 8, Denomination: "fs:", RmMod: false, DecodeFunc: p.fs},
		0x65: {Mnemonic: 0x65, Size: 8, Denomination: "gs:", RmMod: false, DecodeFunc: p.gs},
		0x6C: {Mnemonic: 0x6C, Size: 8, Denomination: "insb", RmMod: false, DecodeFunc: p.insb},
		0x6D: {Mnemonic: 0x6D, Size: 32, Denomination: "insw", RmMod: false, DecodeFunc: p.insw},
		0x6E: {Mnemonic: 0x6E, Size: 8, Denomination: "outsb", RmMod: false, DecodeFunc: p.outsb},
		0x6F: {Mnemonic: 0x6F, Size: 32, Denomination: "outsw", RmMod: false, DecodeFunc: p.outsw},
		0x90: {Mnemonic: 0x90, Size: 0, Denomination: "nop", RmMod: false, DecodeFunc: p.nop},
		0x91: {Mnemonic: 0x91, Size: 32, Denomination: "xchg eax, ecx", RmMod: false, DecodeFunc: p.xchg},
		0x92: {Mnemonic: 0x92, Size: 32, Denomination: "xchg eax, edx", RmMod: false, DecodeFunc: p.xchg},
		0x93: {Mnemonic: 0x93, Size: 32, Denomination: "xchg eax, ebx", RmMod: false, DecodeFunc: p.xchg},
		0x94: {Mnemonic: 0x94, Size: 32, Denomination: "xchg eax, esp", RmMod: false, DecodeFunc: p.xchg},
		0x95: {Mnemonic: 0x95, Size: 32, Denomination: "xchg eax, ebp", RmMod: false, DecodeFunc: p.xchg},
		0x96: {Mnemonic: 0x96, Size: 32, Denomination: "xchg, eax, esi", RmMod: false, DecodeFunc: p.xchg},
		0x97: {Mnemonic: 0x97, Size: 32, Denomination: "xchg eax, edi", RmMod: false, DecodeFunc: p.xchg},
		0x98: {Mnemonic: 0x98, Size: 32, Denomination: "cbw", RmMod: false, DecodeFunc: p.cbw},
		0x99: {Mnemonic: 0x99, Size: 32, Denomination: "cwd", RmMod: false, DecodeFunc: p.cwd},
		0x9B: {Mnemonic: 0x9B, Size: 0, Denomination: "wait", RmMod: false, DecodeFunc: p.wait},
		0x9C: {Mnemonic: 0x9C, Size: 32, Denomination: "pushf", RmMod: false, DecodeFunc: p.push},
		0x9D: {Mnemonic: 0x9D, Size: 32, Denomination: "popf", RmMod: false, DecodeFunc: p.pop},
		0x9E: {Mnemonic: 0x9E, Size: 32, Denomination: "sahf", RmMod: false, DecodeFunc: p.sahf},
		0x9F: {Mnemonic: 0x9F, Size: 32, Denomination: "lahf", RmMod: false, DecodeFunc: p.lahf},
		0xA4: {Mnemonic: 0xA4, Size: 8, Denomination: "movsb", RmMod: false, DecodeFunc: p.movsb},
		0xA5: {Mnemonic: 0xA5, Size: 32, Denomination: "movsw", RmMod: false, DecodeFunc: p.movsw},
		0xA6: {Mnemonic: 0xA6, Size: 8, Denomination: "cmpsb", RmMod: false, DecodeFunc: p.cmpsb},
		0xA7: {Mnemonic: 0xA7, Size: 32, Denomination: "cmpsw", RmMod: false, DecodeFunc: p.cmpsw},
		0xAA: {Mnemonic: 0xAA, Size: 8, Denomination: "stosb", RmMod: false, DecodeFunc: p.stosb},
		0xAB: {Mnemonic: 0xAB, Size: 32, Denomination: "stosw", RmMod: false, DecodeFunc: p.stosw},
		0xAC: {Mnemonic: 0xAC, Size: 8, Denomination: "lodsb", RmMod: false, DecodeFunc: p.lodsb},
		0xAD: {Mnemonic: 0xAD, Size: 32, Denomination: "lodsw", RmMod: false, DecodeFunc: p.lodsw},
		0xAE: {Mnemonic: 0xAE, Size: 8, Denomination: "scasb", RmMod: false, DecodeFunc: p.scasb},
		0xAF: {Mnemonic: 0xAF, Size: 32, Denomination: "scasw", RmMod: false, DecodeFunc: p.scasw},
		0xC3: {Mnemonic: 0xC3, Size: 32, Denomination: "retn", RmMod: false, DecodeFunc: p.retn},
		0xC9: {Mnemonic: 0xC9, Size: 32, Denomination: "leave", RmMod: false, DecodeFunc: p.leave},
		0xCB: {Mnemonic: 0xCB, Size: 32, Denomination: "retf", RmMod: false, DecodeFunc: p.retf},
		0xCC: {Mnemonic: 0xCC, Size: 16, Denomination: "int3", RmMod: false, DecodeFunc: p.int3},
		0xCE: {Mnemonic: 0xCE, Size: 8, Denomination: "into", RmMod: false, DecodeFunc: p.into},
		0xCF: {Mnemonic: 0xCF, Size: 32, Denomination: "iret", RmMod: false, DecodeFunc: p.iret},
		0xD7: {Mnemonic: 0xD7, Size: 32, Denomination: "xlat", RmMod: false, DecodeFunc: p.xlat},
		0xEC: {Mnemonic: 0xEC, Size: 8, Denomination: "in al, dx", RmMod: false, DecodeFunc: p.in},
		0xED: {Mnemonic: 0xED, Size: 32, Denomination: "in eax, dx", RmMod: false, DecodeFunc: p.in},
		0xEE: {Mnemonic: 0xEE, Size: 8, Denomination: "out dx, al", RmMod: false, DecodeFunc: p.out},
		0xEF: {Mnemonic: 0xEF, Size: 32, Denomination: "out dx, eax", RmMod: false, DecodeFunc: p.out},
		0xF0: {Mnemonic: 0xF0, Size: 8, Denomination: "lock", RmMod: false, DecodeFunc: p.lock},
		0xF2: {Mnemonic: 0xF2, Size: 8, Denomination: "repne", RmMod: false, DecodeFunc: p.repne},
		0xF3: {Mnemonic: 0xF3, Size: 8, Denomination: "rep", RmMod: false, DecodeFunc: p.rep},
		0xF4: {Mnemonic: 0xF4, Size: 8, Denomination: "hlt", RmMod: false, DecodeFunc: p.hlt},
		0xF5: {Mnemonic: 0xF5, Size: 8, Denomination: "cmc", RmMod: false, DecodeFunc: p.cmc},
		0xF8: {Mnemonic: 0xF8, Size: 8, Denomination: "clc", RmMod: false, DecodeFunc: p.clc},
		0xF9: {Mnemonic: 0xF9, Size: 8, Denomination: "stc", RmMod: false, DecodeFunc: p.stc},
		0xFA: {Mnemonic: 0xFA, Size: 8, Denomination: "cli", RmMod: false, DecodeFunc: p.cli},
		0xFB: {Mnemonic: 0xFB, Size: 8, Denomination: "sti", RmMod: false, DecodeFunc: p.sti},
		0xFC: {Mnemonic: 0xFC, Size: 8, Denomination: "cld", RmMod: false, DecodeFunc: p.cld},
		0xFD: {Mnemonic: 0xFD, Size: 8, Denomination: "std", RmMod: false, DecodeFunc: p.std},

		// RM MOD opcode

		0x00: {Mnemonic: 0x00, Size: 8, Denomination: "ADD", RmMod: true, DecodeFunc: p.add},
	}
}
