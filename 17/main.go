package main

import (
	"math"
	"os"
	"strconv"
	"strings"

	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
)

func memAccess(registers []int, addr int) int {
	if 0 <= addr && addr <= 3 {
		return addr
	}
	return registers[addr-4]
}

func saveOutput(output []int) string {
	return strings.Join(slicestuff.Map(func(v int) string { return strconv.Itoa(v) }, output), ",")
}

func execProgram(registers, program []int) string {
	ip := 0
	out := []int{}

	for ip < len(program) {
		ins := program[ip]
		jumped := false

		switch ins {
		case 0:
			ip++
			if ip >= len(program) {
				break
			}
			registers[0] = registers[0] / int(math.Pow(2, float64(memAccess(registers, program[ip]))))
		case 1:
			ip++
			if ip >= len(program) {
				break
			}
			registers[1] = registers[1] ^ program[ip]
		case 2:
			ip++
			if ip >= len(program) {
				break
			}
			registers[1] = memAccess(registers, program[ip]) % 8
		case 3:
			if registers[0] != 0 {
				ip++
				if ip >= len(program) {
					break
				}
				ip = program[ip]
				jumped = true
			}
		case 4:
			ip++
			if ip >= len(program) {
				break
			}
			registers[1] = registers[1] ^ registers[2]
		case 5:
			ip++
			if ip >= len(program) {
				break
			}
			out = append(out, memAccess(registers, program[ip])%8)
		case 6:
			ip++
			if ip >= len(program) {
				break
			}
			registers[1] = registers[0] / int(math.Pow(2, float64(memAccess(registers, program[ip]))))
		case 7:
			ip++
			if ip >= len(program) {
				break
			}
			registers[2] = registers[0] / int(math.Pow(2, float64(memAccess(registers, program[ip]))))
		}

		if !jumped {
			ip++
		}
	}

	return saveOutput(out)
}

func valid(program []int, a int) bool {
	for i := 0; i < 11; i++ {
		if calculateResult(a/int(math.Pow(8, float64(i)))) != program[i] {
			return false
		}
	}
	return true
}

func calculateResult(A int) int {
	return ((((A % 8) ^ 3) ^ (A / (1 << ((A % 8) ^ 3)))) ^ 3) % 8
}

func main() {
	data, _ := os.ReadFile("input.txt")

	registers := stringstuff.GetNums(strings.Split(string(data), "\n\n")[0])
	program := stringstuff.GetNums(strings.Split(string(data), "\n\n")[1])

	output := execProgram(registers, program)

	println("part 1:", output)

	target := saveOutput(program)

	for i := int(math.Pow(8, 15)) + 1546749; i < int(math.Pow(8, 16)); i += 4194304 {
		if !valid(program, i) {
			continue
		}
		registers[0] = i
		if execProgram(registers, program) == target {
			println("part 2:", i)
      break
		}
	}
}
