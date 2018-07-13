package klvm

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	HALT = -1
	ADD  = 0
	SUB  = 1
	MULT = 2
	AND  = 3
	OR   = 4
	NOT  = 5
	PUSH = 6
	POP  = 7
)

type Register struct {
	Param1  int
	Param2  int
	Param3  int
	Param4  int
	Param5  int
	Param6  int
	Param7  int
	Param8  int
	Op1     int
	Op2     int
	Result  int
	Condi   int
	InsPtr  int
	DataPtr int
}

type Memory struct {
	nextData     int
	nextInst     int
	runtimeStack []int
	runtimeHeap  []int
	staticData   []int
	code         []int
}

type Machine struct {
	R Register
	M Memory
}

func RegInit() Register {
	return Register{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func MemoryInit() Memory {
	return Memory{0, 0, make([]int, 0, 1024), make([]int, 0, 1024), make([]int, 0, 1024), make([]int, 0, 1024)}
}

func MachineInit() Machine {
	return Machine{
		RegInit(),
		MemoryInit(),
	}
}

func (m *Machine) Load(fileName string) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(file)
	canRead := true
	for canRead {
		line, _, e := r.ReadLine()
		canRead = (e == nil)
		tokens := strings.Split(string(line), " ")
		val, e := strconv.Atoi(string(tokens[0]))
		m.M.code = append(m.M.code, val)
		for i := 1; i < len(tokens); i++ {
			val, _ = strconv.Atoi(string(tokens[0]))
			m.M.staticData = append(m.M.staticData, val)
		}
	}
}

func (m *Machine) Run() {

}
