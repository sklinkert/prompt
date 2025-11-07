package prompt

type Prompt struct {
	Sections                 []Section
	ModelSuggestion          ModelSuggestion
	LangIso6391              string
	SystemContext            string
	OutputMinRequiredWords   int
	ContinuationInstructions string
}

type ModelSuggestion struct {
	HighQualityOutput        bool
	LargeTokenAmountRequired bool
}

func NewPrompt() *Prompt {
	return &Prompt{}
}

func (p *Prompt) SetModelSuggestionHighQualityOutput() {
	p.ModelSuggestion.HighQualityOutput = true
}

func (p *Prompt) SetContinuationInstructions(continuationInstructions string) {
	p.ContinuationInstructions = continuationInstructions
}

func (p *Prompt) SetOutputMinRequiredWords(outputMinRequiredWords int) {
	p.OutputMinRequiredWords = outputMinRequiredWords
}

func (p *Prompt) SetModelSuggestionLargeTokenAmountRequired() {
	p.ModelSuggestion.LargeTokenAmountRequired = true
}

func (p *Prompt) SetLangIso6391(langIso6391 string) {
	p.LangIso6391 = langIso6391
}

func (p *Prompt) SetSystemContext(systemContext string) {
	p.SystemContext = systemContext
}

func (p *Prompt) AddSection(section Section) {
	p.Sections = append(p.Sections, section)
}

func (p *Prompt) AddSections(sections []Section) {
	p.Sections = append(p.Sections, sections...)
}

func (p *Prompt) String() string {
	var output string

	for _, section := range p.Sections {
		output += "\n" + section.String() + "\n---"
	}

	return output
}

func (p *Prompt) WordCount() int {
	var count int

	for _, section := range p.Sections {
		count += section.WordsCount()
	}

	return count
}

// TokenCount - returns the number of tokens in the prompt derived from word count
func (p *Prompt) TokenCount() int {
	var count int

	for _, section := range p.Sections {
		count += int(float64(section.WordsCount()) * 1.4)
	}

	return count
}
