package prompt

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type DataBlock struct {
	Label   string
	Content string
	Type    string // "json", "xml", or "html"
}

type Section struct {
	Intro        string
	Instructions []Instruction
	DataBlocks   []DataBlock
}

type Sections []Section

func NewSection(intro string) Section {
	return Section{
		Intro:        intro,
		Instructions: []Instruction{},
		DataBlocks:   []DataBlock{},
	}
}

func (s *Section) AddInstruction(instruction Instruction) {
	s.Instructions = append(s.Instructions, instruction)
}

// AddJSONData marshals the provided data to JSON and adds it as a data block
func (s *Section) AddJSONData(label string, data any) error {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data to JSON: %w", err)
	}
	s.DataBlocks = append(s.DataBlocks, DataBlock{
		Label:   label,
		Content: string(jsonBytes),
		Type:    "json",
	})
	return nil
}

// AddRawJSON adds pre-formatted JSON string as a data block
func (s *Section) AddRawJSON(label string, jsonString string) {
	s.DataBlocks = append(s.DataBlocks, DataBlock{
		Label:   label,
		Content: jsonString,
		Type:    "json",
	})
}

// AddXMLData marshals the provided data to XML and adds it as a data block
func (s *Section) AddXMLData(label string, data any) error {
	xmlBytes, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data to XML: %w", err)
	}
	s.DataBlocks = append(s.DataBlocks, DataBlock{
		Label:   label,
		Content: string(xmlBytes),
		Type:    "xml",
	})
	return nil
}

// AddRawXML adds pre-formatted XML string as a data block
func (s *Section) AddRawXML(label string, xmlString string) {
	s.DataBlocks = append(s.DataBlocks, DataBlock{
		Label:   label,
		Content: xmlString,
		Type:    "xml",
	})
}

// AddRawHTML adds pre-formatted HTML string as a data block
func (s *Section) AddRawHTML(label string, htmlString string) {
	s.DataBlocks = append(s.DataBlocks, DataBlock{
		Label:   label,
		Content: htmlString,
		Type:    "html",
	})
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

	// Add data blocks after instructions
	for _, block := range s.DataBlocks {
		// Add blank line before data block if there are instructions
		if len(s.Instructions) > 0 {
			output += "\n"
		}

		// Add label if provided
		if block.Label != "" {
			output += block.Label + ":\n"
		}

		// Add code fence with content
		output += "```" + block.Type + "\n"
		output += block.Content + "\n"
		output += "```\n"
	}

	if output != "" && output[len(output)-1] == '\n' {
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
