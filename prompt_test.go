package prompt

import "testing"

func TestPrompt(t *testing.T) {
	p := NewPrompt()
	p.AddSection(Section{"intro1", []Instruction{"test1", "test2"}})
	p.AddSection(Section{"intro2", []Instruction{"test3", "test4"}})

	expected := "\nintro1:\n- test1\n- test2\n---\nintro2:\n- test3\n- test4\n---"
	actual := p.String()

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestPromptWordsCount(t *testing.T) {
	p := NewPrompt()
	p.AddSection(Section{"intro1", []Instruction{"test1", "test2"}})
	p.AddSection(Section{"intro2", []Instruction{"test3", "test4"}})

	expected := 4
	actual := p.WordCount()

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPromptTokenCount(t *testing.T) {
	p := NewPrompt()
	p.AddSection(Section{"intro1", []Instruction{"test1", "test2"}})
	p.AddSection(Section{"intro2", []Instruction{"test3", "test4", "test5"}})

	expected := 6
	actual := p.TokenCount()

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPromptEmpty(t *testing.T) {
	p := NewPrompt()
	if p.String() != "" {
		t.Errorf("Expected empty string for prompt with no sections, got %s", p.String())
	}
	if p.WordCount() != 0 {
		t.Errorf("Expected 0 word count, got %d", p.WordCount())
	}
	if p.TokenCount() != 0 {
		t.Errorf("Expected 0 token count, got %d", p.TokenCount())
	}
}

func TestPromptSetters(t *testing.T) {
	p := NewPrompt()

	// Test SetLangIso6391
	p.SetLangIso6391("en")
	if p.LangIso6391 != "en" {
		t.Errorf("Expected LangIso6391 to be 'en', got %s", p.LangIso6391)
	}

	// Test SetSystemContext
	p.SetSystemContext("test context")
	if p.SystemContext != "test context" {
		t.Errorf("Expected SystemContext to be 'test context', got %s", p.SystemContext)
	}

	// Test SetOutputMinRequiredWords
	p.SetOutputMinRequiredWords(500)
	if p.OutputMinRequiredWords != 500 {
		t.Errorf("Expected OutputMinRequiredWords to be 500, got %d", p.OutputMinRequiredWords)
	}

	// Test SetContinuationInstructions
	p.SetContinuationInstructions("continue here")
	if p.ContinuationInstructions != "continue here" {
		t.Errorf("Expected ContinuationInstructions to be 'continue here', got %s", p.ContinuationInstructions)
	}
}

func TestPromptModelSuggestions(t *testing.T) {
	p := NewPrompt()

	// Initially should be false
	if p.ModelSuggestion.HighQualityOutput {
		t.Error("Expected HighQualityOutput to be false initially")
	}
	if p.ModelSuggestion.LargeTokenAmountRequired {
		t.Error("Expected LargeTokenAmountRequired to be false initially")
	}

	// Test SetModelSuggestionHighQualityOutput
	p.SetModelSuggestionHighQualityOutput()
	if !p.ModelSuggestion.HighQualityOutput {
		t.Error("Expected HighQualityOutput to be true after setting")
	}

	// Test SetModelSuggestionLargeTokenAmountRequired
	p.SetModelSuggestionLargeTokenAmountRequired()
	if !p.ModelSuggestion.LargeTokenAmountRequired {
		t.Error("Expected LargeTokenAmountRequired to be true after setting")
	}
}

func TestPromptAddSections(t *testing.T) {
	p := NewPrompt()

	sections := []Section{
		{"intro1", []Instruction{"test1"}},
		{"intro2", []Instruction{"test2"}},
	}

	p.AddSections(sections)

	if len(p.Sections) != 2 {
		t.Errorf("Expected 2 sections, got %d", len(p.Sections))
	}
}

func TestSection(t *testing.T) {
	section := NewSection("Test intro")

	if section.Intro != "Test intro" {
		t.Errorf("Expected intro to be 'Test intro', got %s", section.Intro)
	}

	if len(section.Instructions) != 0 {
		t.Errorf("Expected 0 instructions, got %d", len(section.Instructions))
	}
}

func TestSectionAddInstruction(t *testing.T) {
	section := NewSection("Test")
	section.AddInstruction("instruction1")
	section.AddInstruction("instruction2")

	if len(section.Instructions) != 2 {
		t.Errorf("Expected 2 instructions, got %d", len(section.Instructions))
	}
}

func TestSectionString(t *testing.T) {
	tests := []struct {
		name     string
		section  Section
		expected string
	}{
		{
			name:     "Section with intro and instructions",
			section:  Section{"Test", []Instruction{"inst1", "inst2"}},
			expected: "Test:\n- inst1\n- inst2",
		},
		{
			name:     "Section with intro ending in colon",
			section:  Section{"Test:", []Instruction{"inst1"}},
			expected: "Test:\n- inst1",
		},
		{
			name:     "Section without intro",
			section:  Section{"", []Instruction{"inst1", "inst2"}},
			expected: "- inst1\n- inst2",
		},
		{
			name:     "Empty section",
			section:  Section{"", []Instruction{}},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.section.String()
			if actual != tt.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s", tt.expected, actual)
			}
		})
	}
}

func TestSectionWordsCount(t *testing.T) {
	section := Section{
		Intro: "ignored",
		Instructions: []Instruction{
			"one two three",
			"four five",
		},
	}

	expected := 5
	actual := section.WordsCount()

	if actual != expected {
		t.Errorf("Expected %d words, got %d", expected, actual)
	}
}

func TestSectionsString(t *testing.T) {
	sections := Sections{
		{"intro1", []Instruction{"test1"}},
		{"intro2", []Instruction{"test2"}},
	}

	expected := "\nintro1:\n- test1\n---\nintro2:\n- test2\n---"
	actual := sections.String()

	if actual != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}

func TestWordsCount(t *testing.T) {
	sections := []Section{
		{"intro1", []Instruction{"one two"}},
		{"intro2", []Instruction{"three four five"}},
	}

	expected := 5
	actual := WordsCount(sections)

	if actual != expected {
		t.Errorf("Expected %d words, got %d", expected, actual)
	}
}

func TestInstruction(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple instruction",
			input:    "test instruction",
			expected: "test instruction",
		},
		{
			name:     "Instruction with whitespace",
			input:    "  test instruction  ",
			expected: "test instruction",
		},
		{
			name:     "Instruction with triple dash",
			input:    "first part---second part",
			expected: "first part\n\nsecond part",
		},
		{
			name:     "Instruction with multiple triple dashes",
			input:    "one---two---three",
			expected: "one\n\ntwo\n\nthree",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instruction := NewInstruction(tt.input)
			actual := instruction.String()
			if actual != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, actual)
			}
		})
	}
}

func TestInstructionWordCount(t *testing.T) {
	tests := []struct {
		name     string
		input    Instruction
		expected int
	}{
		{
			name:     "Single word",
			input:    "word",
			expected: 1,
		},
		{
			name:     "Multiple words",
			input:    "one two three four",
			expected: 4,
		},
		{
			name:     "Empty instruction",
			input:    "",
			expected: 0,
		},
		{
			name:     "Words with extra spaces",
			input:    "one  two   three",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.input.WordCount()
			if actual != tt.expected {
				t.Errorf("Expected %d words, got %d", tt.expected, actual)
			}
		})
	}
}

// Benchmark tests

func BenchmarkPromptString(b *testing.B) {
	p := NewPrompt()
	for i := 0; i < 10; i++ {
		section := NewSection("Section intro")
		for j := 0; j < 5; j++ {
			section.AddInstruction(NewInstruction("This is a test instruction with multiple words"))
		}
		p.AddSection(section)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.String()
	}
}

func BenchmarkPromptWordCount(b *testing.B) {
	p := NewPrompt()
	for i := 0; i < 10; i++ {
		section := NewSection("Section intro")
		for j := 0; j < 5; j++ {
			section.AddInstruction(NewInstruction("This is a test instruction with multiple words"))
		}
		p.AddSection(section)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.WordCount()
	}
}

func BenchmarkPromptTokenCount(b *testing.B) {
	p := NewPrompt()
	for i := 0; i < 10; i++ {
		section := NewSection("Section intro")
		for j := 0; j < 5; j++ {
			section.AddInstruction(NewInstruction("This is a test instruction with multiple words"))
		}
		p.AddSection(section)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.TokenCount()
	}
}

func BenchmarkNewInstruction(b *testing.B) {
	text := "  This is a test instruction with triple dashes---that need to be replaced  "

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewInstruction(text)
	}
}
