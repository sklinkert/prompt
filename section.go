package prompt

type Section struct {
	Intro        string
	Instructions []Instruction
}

type Sections []Section

func NewSection(intro string) Section {
	return Section{intro, []Instruction{}}
}

func (s *Section) AddInstruction(instruction Instruction) {
	s.Instructions = append(s.Instructions, instruction)
}

func (s *Section) String() string {
	var output string

	if s.Intro != "" {
		// if last char of intro is not ':', add ':'
		if s.Intro[len(s.Intro)-1] != ':' {
			s.Intro += ":"
		}

		output += s.Intro + "\n"
	}

	for _, instruction := range s.Instructions {
		output += "- " + string(instruction) + "\n"
	}

	if output != "" {
		output = output[:len(output)-1]
	}

	return output
}

func (s *Section) WordsCount() int {
	var count int

	for _, instruction := range s.Instructions {
		count += instruction.WordCount()
	}

	return count
}

func (ss Sections) String() string {
	var output string

	for _, section := range ss {
		output += "\n" + section.String() + "\n---"
	}

	return output
}

func WordsCount(sections []Section) int {
	var count int

	for _, section := range sections {
		count += section.WordsCount()
	}

	return count
}
