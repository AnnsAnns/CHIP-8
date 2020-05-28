package main

import "fmt"
import "io/ioutil"

var opcode uint16 // Should be short-ish?
var memory[4096] byte // Should be char-ish?

// https://en.wikipedia.org/wiki/CHIP-8#Virtual_machine_description
// 4096 Total, 512 CHIP-8 Interpreter, 0-256 Display, 257-257+96 Call Stack and other stuff
// Rest for the games

var V[16] byte // CPU Registers
var I uint // Index Register I
var pc uint = 0x200 // Program Counter (PC)
var gfx[64 * 32] byte // resolution of CHIP-8

var delay_timer byte
var sound_timer byte

var stack[16] uint16 //16 levels of Stack
var sp uint // Stack Pointer

var key[16] byte // Keypad

func rom_load() {
	program, err := ioutil.ReadFile("TETRIS")
	if err != nil {
		fmt.Println("Error Loading ROM")
		panic(err)
	}

	for i, bit := range program {
		memory[512+i] = bit
	}
}

func emu_init() {

}

func emu_cycle() {
	opcode = uint16(memory[pc] << 8 | memory[pc+1])
}

func decode_opcodes(opcode uint16) {

	switch (opcode & 0xF000) {
	case 0x0000:
		switch(opcode & 0x000F) {
		case 0x0000: // Clears the screen

		break

		case 0x000E: // Returns from subroutine

		break

		default:
			fmt.Printf("Unknown opcode [0x0000]: 0x%X\n", opcode)
		}
	break

	}
}

func main() {


}