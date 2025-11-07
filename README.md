# Prompt

A lightweight, well-structured Go library for building and managing prompts for Large Language Models (LLMs). This library provides a clean builder pattern for creating complex, multi-section prompts with metadata to guide model selection and behavior.

[![Go Reference](https://pkg.go.dev/badge/github.com/sklinkert/prompt.svg)](https://pkg.go.dev/github.com/sklinkert/prompt)
[![Go Report Card](https://goreportcard.com/badge/github.com/sklinkert/prompt)](https://goreportcard.com/report/github.com/sklinkert/prompt)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Features

- **Builder Pattern**: Fluent API for constructing complex prompts
- **Structured Sections**: Organize prompts into logical sections with intros and instructions
- **Model Hints**: Provide suggestions for high-quality output or large token requirements
- **Flexible Metadata**: Generic key-value metadata system with type-safe getters and backward-compatible helpers
- **Word & Token Counting**: Built-in utilities for estimating prompt size
- **Zero Dependencies**: Uses only Go standard library

## Installation

```bash
go get github.com/sklinkert/prompt
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/sklinkert/prompt"
)

func main() {
    // Create a new prompt
    p := prompt.NewPrompt()

    // Add a section with instructions
    section := prompt.NewSection("Write a blog post about Go programming")
    section.AddInstruction(prompt.NewInstruction("Focus on concurrency patterns"))
    section.AddInstruction(prompt.NewInstruction("Include code examples"))
    section.AddInstruction(prompt.NewInstruction("Target audience: intermediate developers"))

    p.AddSection(section)

    // Set metadata
    p.SetLangIso6391("en")
    p.SetSystemContext("You are a technical writer specializing in Go")
    p.SetModelSuggestionHighQualityOutput()
    p.SetOutputMinRequiredWords(500)

    // Output the formatted prompt
    fmt.Println(p.String())

    // Check prompt size
    fmt.Printf("Word count: %d\n", p.WordCount())
    fmt.Printf("Estimated tokens: %d\n", p.TokenCount())
}
```

## API Documentation

### Prompt

The main structure for building prompts.

#### Creating a Prompt

```go
p := prompt.NewPrompt()
```

#### Adding Content

```go
// Add a single section
section := prompt.Section{
    Intro: "Task description",
    Instructions: []prompt.Instruction{
        prompt.NewInstruction("Requirement 1"),
        prompt.NewInstruction("Requirement 2"),
    },
}
p.AddSection(section)

// Add multiple sections at once
sections := []prompt.Section{section1, section2}
p.AddSections(sections)
```

#### Setting Metadata

**Generic Metadata (Recommended)**

Store arbitrary metadata with flexible key-value pairs:

```go
// Set custom metadata with any type
p.SetMetadata("temperature", 0.7)
p.SetMetadata("model_name", "gpt-4")
p.SetMetadata("max_tokens", 2000)
p.SetMetadata("use_streaming", true)

// Retrieve metadata
value, exists := p.GetMetadata("temperature")
if exists {
    temp := value.(float64)
}

// Type-safe retrieval
modelName := p.GetMetadataString("model_name")      // Returns "" if not found or wrong type
maxTokens := p.GetMetadataInt("max_tokens")         // Returns 0 if not found or wrong type
streaming := p.GetMetadataBool("use_streaming")     // Returns false if not found or wrong type

// Check existence
if p.HasMetadata("temperature") {
    // ...
}

// Get all metadata
allMeta := p.GetAllMetadata()

// Delete metadata
p.DeleteMetadata("temperature")
```

**Well-Known Metadata (Backward Compatible)**

For common metadata, convenience methods are available:

```go
// Set language (ISO 639-1 code)
p.SetLangIso6391("en")
lang := p.GetLangIso6391()

// Set system context for the LLM
p.SetSystemContext("You are a helpful assistant")
context := p.GetSystemContext()

// Set output requirements
p.SetOutputMinRequiredWords(300)
minWords := p.GetOutputMinRequiredWords()

// Set continuation instructions for multi-turn conversations
p.SetContinuationInstructions("Continue from where you left off")
instructions := p.GetContinuationInstructions()

// Or use generic metadata with well-known keys
p.SetMetadata("lang_iso_6391", "en")
p.SetMetadata("system_context", "You are a helpful assistant")
p.SetMetadata("output_min_required_words", 300)
p.SetMetadata("continuation_instructions", "Continue from where you left off")
p.SetMetadata("model_high_quality", true)
p.SetMetadata("model_large_tokens", true)
```

#### Model Suggestions

```go
// Suggest high-quality output is needed
p.SetModelSuggestionHighQualityOutput()
isHighQuality := p.GetModelSuggestionHighQualityOutput()

// Suggest large token amount will be required
p.SetModelSuggestionLargeTokenAmountRequired()
needsLargeTokens := p.GetModelSuggestionLargeTokenAmountRequired()
```

#### Output and Metrics

```go
// Get formatted prompt string
promptText := p.String()

// Get word count (counts actual words in instructions)
words := p.WordCount()

// Get estimated token count (uses 1.4x multiplier)
tokens := p.TokenCount()
```

### Section

Represents a logical section of the prompt with an intro and instructions.

```go
// Create a new section
section := prompt.NewSection("Task description")

// Add instructions dynamically
section.AddInstruction(prompt.NewInstruction("First instruction"))
section.AddInstruction(prompt.NewInstruction("Second instruction"))

// Get formatted section
formatted := section.String()

// Count words in section
words := section.WordsCount()
```

### Instruction

Represents a single instruction within a section.

```go
// Create an instruction (trims whitespace, replaces "---" with "\n\n")
instruction := prompt.NewInstruction("  Write clear code  ")

// Get string representation
text := instruction.String()

// Count words
words := instruction.WordCount()
```

## Usage Examples

### Example 1: Simple Content Generation

```go
p := prompt.NewPrompt()

section := prompt.NewSection("Generate a product description")
section.AddInstruction(prompt.NewInstruction("Product: Wireless headphones"))
section.AddInstruction(prompt.NewInstruction("Tone: Professional and engaging"))
section.AddInstruction(prompt.NewInstruction("Length: 150-200 words"))

p.AddSection(section)
p.SetLangIso6391("en")
```

### Example 2: Multi-Section Research Prompt

```go
p := prompt.NewPrompt()

// Research section
research := prompt.NewSection("Research the following topics")
research.AddInstruction(prompt.NewInstruction("History of cloud computing"))
research.AddInstruction(prompt.NewInstruction("Current market leaders"))

// Analysis section
analysis := prompt.NewSection("Analyze the findings")
analysis.AddInstruction(prompt.NewInstruction("Identify key trends"))
analysis.AddInstruction(prompt.NewInstruction("Compare different approaches"))

// Output section
output := prompt.NewSection("Format the output")
output.AddInstruction(prompt.NewInstruction("Use markdown formatting"))
output.AddInstruction(prompt.NewInstruction("Include citations"))

p.AddSections([]prompt.Section{research, analysis, output})
p.SetModelSuggestionHighQualityOutput()
p.SetModelSuggestionLargeTokenAmountRequired()
```

### Example 3: Code Generation with Context

```go
p := prompt.NewPrompt()

section := prompt.NewSection("Generate a REST API handler in Go")
section.AddInstruction(prompt.NewInstruction("Endpoint: POST /users"))
section.AddInstruction(prompt.NewInstruction("Validate email and password"))
section.AddInstruction(prompt.NewInstruction("Return JWT token on success"))
section.AddInstruction(prompt.NewInstruction("Use echo/v4 framework"))

p.AddSection(section)
p.SetSystemContext("You are an expert Go backend developer")
p.SetLangIso6391("en")
```

### Example 4: Continuation Instructions

```go
p := prompt.NewPrompt()

section := prompt.NewSection("Write a technical article")
section.AddInstruction(prompt.NewInstruction("Topic: Microservices architecture"))
section.AddInstruction(prompt.NewInstruction("Target: 2000 words"))

p.AddSection(section)
p.SetOutputMinRequiredWords(2000)
p.SetContinuationInstructions("Continue writing the article, maintaining the same style and tone. Pick up from where you left off without repeating content.")
```

### Example 5: Custom Metadata for LLM Configuration

```go
p := prompt.NewPrompt()

section := prompt.NewSection("Translate the following text")
section.AddInstruction(prompt.NewInstruction("Source: 'Hello, how are you today?'"))
section.AddInstruction(prompt.NewInstruction("Target language: Spanish"))

p.AddSection(section)

// Use generic metadata to store LLM configuration
p.SetMetadata("model", "gpt-4")
p.SetMetadata("temperature", 0.3)
p.SetMetadata("max_tokens", 150)
p.SetMetadata("top_p", 1.0)
p.SetMetadata("presence_penalty", 0.0)
p.SetMetadata("frequency_penalty", 0.0)

// Store custom application metadata
p.SetMetadata("user_id", "user-12345")
p.SetMetadata("request_id", "req-abc-123")
p.SetMetadata("cost_center", "translation-service")

// Later, retrieve metadata for API call
if p.HasMetadata("temperature") {
    temp := p.GetMetadata("temperature")
    // Use temp in your LLM API call
}

// Get all metadata for logging
allMeta := p.GetAllMetadata()
fmt.Printf("Request metadata: %+v\n", allMeta)
```

## Output Format

The `String()` method formats the prompt as follows:

```
intro1:
- instruction 1
- instruction 2
---
intro2:
- instruction 3
- instruction 4
---
```

Sections are separated by `---`, and instructions are formatted as bullet points under their section intro.

## Word and Token Counting

- **Word Count**: Counts actual words in all instructions using `strings.Fields()`
- **Token Count**: Estimates tokens by multiplying word count by 1.4 (a common approximation for English text)

These utilities help estimate prompt size for LLM context limits.

## Use Cases

This library is ideal for:

- Building dynamic prompts for OpenAI, Anthropic, Google, or other LLM APIs
- Creating reusable prompt templates
- Managing complex multi-section prompts
- Estimating prompt costs based on token usage
- Implementing prompt engineering patterns in production applications

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development

```bash
# Run tests
go test -v ./...

# Run tests with coverage
go test -v -race -coverprofile=coverage.out ./...

# View coverage report
go tool cover -html=coverage.out

# Run linters
go fmt ./...
go vet ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

Extracted from the [artitext](https://github.com/sklinkert/artitext) project and made available as a standalone library for the Go community.
