package prompt

// Well-known metadata keys
const (
	MetadataKeyLangIso6391              = "lang_iso_6391"
	MetadataKeySystemContext            = "system_context"
	MetadataKeyOutputMinRequiredWords   = "output_min_required_words"
	MetadataKeyContinuationInstructions = "continuation_instructions"
	MetadataKeyModelHighQuality         = "model_high_quality"
	MetadataKeyModelLargeTokens         = "model_large_tokens"
)

type Prompt struct {
	Sections []Section
	metadata map[string]any
}

func NewPrompt() *Prompt {
	return &Prompt{
		metadata: make(map[string]any),
	}
}

// SetMetadata sets an arbitrary metadata value
func (p *Prompt) SetMetadata(key string, value any) {
	if p.metadata == nil {
		p.metadata = make(map[string]any)
	}
	p.metadata[key] = value
}

// GetMetadata retrieves a metadata value
func (p *Prompt) GetMetadata(key string) (any, bool) {
	if p.metadata == nil {
		return nil, false
	}
	value, exists := p.metadata[key]
	return value, exists
}

// GetMetadataString retrieves a metadata value as string
func (p *Prompt) GetMetadataString(key string) string {
	value, exists := p.GetMetadata(key)
	if !exists {
		return ""
	}
	if str, ok := value.(string); ok {
		return str
	}
	return ""
}

// GetMetadataInt retrieves a metadata value as int
func (p *Prompt) GetMetadataInt(key string) int {
	value, exists := p.GetMetadata(key)
	if !exists {
		return 0
	}
	if i, ok := value.(int); ok {
		return i
	}
	return 0
}

// GetMetadataBool retrieves a metadata value as bool
func (p *Prompt) GetMetadataBool(key string) bool {
	value, exists := p.GetMetadata(key)
	if !exists {
		return false
	}
	if b, ok := value.(bool); ok {
		return b
	}
	return false
}

// GetAllMetadata returns a copy of all metadata
func (p *Prompt) GetAllMetadata() map[string]any {
	if p.metadata == nil {
		return make(map[string]any)
	}
	copy := make(map[string]any, len(p.metadata))
	for k, v := range p.metadata {
		copy[k] = v
	}
	return copy
}

// HasMetadata checks if a metadata key exists
func (p *Prompt) HasMetadata(key string) bool {
	_, exists := p.GetMetadata(key)
	return exists
}

// DeleteMetadata removes a metadata key
func (p *Prompt) DeleteMetadata(key string) {
	if p.metadata != nil {
		delete(p.metadata, key)
	}
}

// Backward compatible methods

func (p *Prompt) SetModelSuggestionHighQualityOutput() {
	p.SetMetadata(MetadataKeyModelHighQuality, true)
}

func (p *Prompt) SetContinuationInstructions(continuationInstructions string) {
	p.SetMetadata(MetadataKeyContinuationInstructions, continuationInstructions)
}

func (p *Prompt) SetOutputMinRequiredWords(outputMinRequiredWords int) {
	p.SetMetadata(MetadataKeyOutputMinRequiredWords, outputMinRequiredWords)
}

func (p *Prompt) SetModelSuggestionLargeTokenAmountRequired() {
	p.SetMetadata(MetadataKeyModelLargeTokens, true)
}

func (p *Prompt) SetLangIso6391(langIso6391 string) {
	p.SetMetadata(MetadataKeyLangIso6391, langIso6391)
}

func (p *Prompt) SetSystemContext(systemContext string) {
	p.SetMetadata(MetadataKeySystemContext, systemContext)
}

// Backward compatible getters

func (p *Prompt) GetModelSuggestionHighQualityOutput() bool {
	return p.GetMetadataBool(MetadataKeyModelHighQuality)
}

func (p *Prompt) GetContinuationInstructions() string {
	return p.GetMetadataString(MetadataKeyContinuationInstructions)
}

func (p *Prompt) GetOutputMinRequiredWords() int {
	return p.GetMetadataInt(MetadataKeyOutputMinRequiredWords)
}

func (p *Prompt) GetModelSuggestionLargeTokenAmountRequired() bool {
	return p.GetMetadataBool(MetadataKeyModelLargeTokens)
}

func (p *Prompt) GetLangIso6391() string {
	return p.GetMetadataString(MetadataKeyLangIso6391)
}

func (p *Prompt) GetSystemContext() string {
	return p.GetMetadataString(MetadataKeySystemContext)
}

// Deprecated: Use GetModelSuggestionHighQualityOutput and GetModelSuggestionLargeTokenAmountRequired
type ModelSuggestion struct {
	HighQualityOutput        bool
	LargeTokenAmountRequired bool
}

// GetModelSuggestion returns model suggestions in the deprecated format for backward compatibility
// Deprecated: Access metadata directly using GetModelSuggestionHighQualityOutput and GetModelSuggestionLargeTokenAmountRequired
func (p *Prompt) GetModelSuggestion() ModelSuggestion {
	return ModelSuggestion{
		HighQualityOutput:        p.GetModelSuggestionHighQualityOutput(),
		LargeTokenAmountRequired: p.GetModelSuggestionLargeTokenAmountRequired(),
	}
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
