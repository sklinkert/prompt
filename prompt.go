package prompt

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
