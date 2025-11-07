package prompt

import "strings"

type Instruction string

func NewInstruction(instruction string) Instruction {
	instruction = strings.TrimSpace(instruction)
	instruction = strings.ReplaceAll(instruction, "---", "\n\n")

	return Instruction(instruction)
}

func (i Instruction) String() string {
	return string(i)
}

func (i Instruction) WordCount() int {
	return len(strings.Fields(i.String()))
}
