package main

import "io/ioutil"
import "fmt"

type Chip8 struct {
	memory[4096] byte // 4096 bytes of pure Memory

	//0x0 - 0x1FF Chip8 Interpreter
	//0x050-0x0A0 Fonts
	//0x200-0xFFF Program and RAM

	opcode uint // Place for the opcode
	V[16] byte // CPU registers
	I uint // Index register
	pc uint // Program Counter
	gfx[64*32] byte // Graphics at a staggering 64x32 pixels

	delay_timer int
	sound_timer int

	stack[16] uint
	stackpointer uint

	key[16] uint8 // Keys

	draw_flag bool
}

func (chip8 *Chip8) emu_init() {
	//Clears all variables
	chip8.pc = 0x200
	chip8.opcode = 0
	chip8.I = 0
	chip8.stackpointer = 0

	// Clear GFX
	for i := range chip8.gfx {
		chip8.gfx[i] = 0
	}

	// Clear Stack
	for i := range chip8.stack {
		chip8.stack[i] = 0
	}

	// Clear registers V0-VF
	for i := range chip8.V {
		chip8.V[i] = 0
	}

	// Clear Memory
	for i := range chip8.memory {
		chip8.memory[i] = 0
	}

	fmt.Println("Cleared all variables!")

	// Load Fontset

	fmt.Println("Loaded font!")

	// Load Game
	game, err := ioutil.ReadFile("TETRIS")
	if err != nil {
		fmt.Println("Loading game failed!")
		panic(err)
	}

	for i, b := range game {
		chip8.memory[512 + i] = b
	}

	fmt.Println("Loaded game!")
}

func (chip8 *Chip8) emu_cycle() {
	chip8.opcode = uint(chip8.memory[chip8.pc]) << 8 | uint(chip8.memory[chip8.pc + 1]) // Get next opcode

	// Decode opcode, pc += 2 -> next cycle, pc += 4 -> skip cycle
	switch (chip8.opcode & 0xF000) {
	
	case 0x0000:
		switch (chip8.opcode & 0x000F) {
		case 0x0000: // 0x00E0: Clears the screen
		break

		case 0x000E: // 0x00EE: Returns from subroutine
		chip8.pc = uint(chip8.memory[chip8.stackpointer]) << 8 | uint(chip8.memory[chip8.stackpointer + 1])
		chip8.stackpointer -= 1
		chip8.pc += 2 // Might be wrong here
		break
		}
		break
	case 0x1000: // 1NNN: Jump to location nnn.
		chip8.pc = chip8.opcode & 0x0FFF
		break
	case 0x2000: // 2NNN: Calls subroutine at adress NNN
		chip8.stack[chip8.stackpointer] = chip8.pc
		chip8.stackpointer += 1
		chip8.pc = chip8.opcode & 0x0FFF
		break
	case 0x3000: // 3XKK Skip next instruction if Vx = kk.
		if chip8.V[chip8.opcode & 0x0F00] == byte(chip8.opcode & 0x00FF) {
			chip8.pc += 4
		} else {
			chip8.pc += 2
		}
		break
	case 0x4000: // 4XKK Skip next instruction if NOT Vx = kk.
		if chip8.V[chip8.opcode & 0x0F00] != byte(chip8.opcode & 0x00FF) {
			chip8.pc += 4
		} else {
			chip8.pc += 2
		}
		break
	case 0x5000: // 5xy0: Skip next instruction if Vx = Vy.
		if chip8.V[chip8.opcode & 0x0F00] == chip8.V[chip8.opcode & 0x00F0] {
			chip8.pc += 4
		} else {
			chip8.pc += 2
		}
		break
	case 0x6000: // 6xkk: Set Vx = kk
		chip8.V[chip8.opcode & 0x0F00] = byte(chip8.opcode & 0x00FF)
		chip8.pc += 2
		break
	case 0x7000: // 7xkk: Vx = Vx + kk.
		chip8.V[chip8.opcode & 0x0F00] += byte(chip8.opcode & 0x00FF)
		chip8.pc += 2
		break
	case 0x8000: 
		switch (chip8.opcode & 0x000F) {
		case 0x0000: // 8XY0: Set Vx = Vy.
			chip8.V[chip8.opcode & 0x0F00] = chip8.V[chip8.opcode & 0x00F0]
			chip8.pc += 2
			break
		case 0x0001: // 8XY1: set Vx = Vx OR Vy.
		break
		case 0x0002: // 8XY2: Set Vx = Vx AND Vy.
		break
		case 0x0003: // 8XY3: Set Vx = Vx XOR Vy.
		break
		case 0x0004: // 8XY4: Set Vx = Vx + Vy, set VF = carry.
		break
		case 0x0005: // 8XY5: Set Vx = Vx - Vy, set VF = NOT borrow.
		break
		case 0x0006: // 8XY6: Set Vx = Vx SHR 1.
		break
		case 0x0007: // 8XY7: Set Vx = Vy - Vx, set VF = NOT borrow.
		break
		case 0x000E: // 8XY8: Set Vx = Vx SHL 1.
		break
		}
		break
	case 0xA000: // ANNN: Sets I to the adress NNN
		chip8.I = chip8.opcode & 0x0FFF
		chip8.pc += 2
		break

	//default:
	//	fmt.Printf("Unimplemented Opcode 0x%X", chip8.opcode)
	//	panic(chip8.opcode)
	}

	//Updates Timers
	if chip8.delay_timer > 0 {
		chip8.delay_timer -= 1
	}

	if chip8.sound_timer > 0 {
		if chip8.sound_timer == 1 {
			fmt.Println("Imagine that this is beeping right now!")
			chip8.sound_timer -= 1
		}
	}
}

func (chip8 *Chip8) opcode_handling() {

}

func (chip8 *Chip8) set_keys() {
// Check for newly pressed keys
}

func main() {

	// Graphics
	// Input

	var emu Chip8
	emu.emu_init()

	for {
		emu.emu_cycle()

		if emu.draw_flag {
			//Graphics should be drawn
		}

		emu.set_keys() // Check for new keypresses by the user
	}
}