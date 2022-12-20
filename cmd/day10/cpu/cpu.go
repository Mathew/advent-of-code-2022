package cpu

type Command interface {
	Execute(int) int
	IsComplete() bool
}

type Noop struct {
	cycles int
}

func NewNoop() *Noop {
	return &Noop{1}
}

func (n *Noop) Execute(i int) int {
	n.cycles--
	return i
}

func (n Noop) IsComplete() bool { return n.cycles == 0 }

type Add struct {
	value  int
	cycles int
}

func NewAdd(value int) *Add {
	return &Add{value, 2}
}

func (a *Add) Execute(i int) int {
	a.cycles--
	if a.cycles == 0 {
		return i + a.value
	}

	return i
}

func (a Add) IsComplete() bool {
	return a.cycles == 0
}

func NewCPU(instructions []Command) func() int {
	register := 1
	commandIndex := 0

	return func() int {
		if commandIndex == len(instructions) {
			return register
		}

		value := register

		register = instructions[commandIndex].Execute(register)
		if instructions[commandIndex].IsComplete() {
			commandIndex += 1
		}
		return value
	}
}

func CycleCounter(cycleMeasurements []int, cpu func() int) func() int {
	cycle := 0
	total := 0
	cycI := 0

	return func() int {
		for cycle != cycleMeasurements[cycI] {
			cycle += 1
			register := cpu()

			if cycle == cycleMeasurements[cycI] {

				cycI += 1
				total += register * cycle
				return register * cycle
			}
		}

		return total
	}
}
