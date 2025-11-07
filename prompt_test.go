package prompt

import "testing"

func TestPrompt(t *testing.T) {
	p := NewPrompt()
	p.AddSection(Section{"intro1", []Instruction{"test1", "test2"}, []DataBlock{}})
	p.AddSection(Section{"intro2", []Instruction{"test3", "test4"}, []DataBlock{}})

	expected := "\nintro1:\n- test1\n- test2\n---\nintro2:\n- test3\n- test4\n---"
	actual := p.String()

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestPromptWordsCount(t *testing.T) {
	p := NewPrompt()
	p.AddSection(Section{"intro1", []Instruction{"test1", "test2"}, []DataBlock{}})
	p.AddSection(Section{"intro2", []Instruction{"test3", "test4"}, []DataBlock{}})

	expected := 4
	actual := p.WordCount()

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPromptTokenCount(t *testing.T) {
	p := NewPrompt()
	p.AddSection(Section{"intro1", []Instruction{"test1", "test2"}, []DataBlock{}})
	p.AddSection(Section{"intro2", []Instruction{"test3", "test4", "test5"}, []DataBlock{}})

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

func TestPromptAddSections(t *testing.T) {
	p := NewPrompt()

	sections := []Section{
		{"intro1", []Instruction{"test1"}, []DataBlock{}},
		{"intro2", []Instruction{"test2"}, []DataBlock{}},
	}

	p.AddSections(sections)

	if len(p.Sections) != 2 {
		t.Errorf("Expected 2 sections, got %d", len(p.Sections))
	}
}

func TestPromptGenericMetadata(t *testing.T) {
	p := NewPrompt()

	// Test setting and getting generic metadata
	p.SetMetadata("custom_key", "custom_value")
	value, exists := p.GetMetadata("custom_key")
	if !exists {
		t.Error("Expected custom_key to exist")
	}
	if value != "custom_value" {
		t.Errorf("Expected 'custom_value', got %v", value)
	}

	// Test non-existent key
	_, exists = p.GetMetadata("non_existent")
	if exists {
		t.Error("Expected non_existent key to not exist")
	}

	// Test different data types
	p.SetMetadata("int_value", 42)
	p.SetMetadata("bool_value", true)
	p.SetMetadata("float_value", 3.14)

	intVal, exists := p.GetMetadata("int_value")
	if !exists || intVal != 42 {
		t.Errorf("Expected int_value to be 42, got %v", intVal)
	}

	boolVal, exists := p.GetMetadata("bool_value")
	if !exists || boolVal != true {
		t.Errorf("Expected bool_value to be true, got %v", boolVal)
	}

	floatVal, exists := p.GetMetadata("float_value")
	if !exists || floatVal != 3.14 {
		t.Errorf("Expected float_value to be 3.14, got %v", floatVal)
	}
}

func TestPromptMetadataTypedGetters(t *testing.T) {
	p := NewPrompt()

	// Test GetMetadataString
	p.SetMetadata("string_key", "test_string")
	if p.GetMetadataString("string_key") != "test_string" {
		t.Errorf("Expected 'test_string', got %s", p.GetMetadataString("string_key"))
	}

	// Test GetMetadataString with non-existent key
	if p.GetMetadataString("non_existent") != "" {
		t.Error("Expected empty string for non-existent key")
	}

	// Test GetMetadataString with wrong type
	p.SetMetadata("int_key", 42)
	if p.GetMetadataString("int_key") != "" {
		t.Error("Expected empty string for wrong type")
	}

	// Test GetMetadataInt
	p.SetMetadata("int_key", 123)
	if p.GetMetadataInt("int_key") != 123 {
		t.Errorf("Expected 123, got %d", p.GetMetadataInt("int_key"))
	}

	// Test GetMetadataInt with non-existent key
	if p.GetMetadataInt("non_existent") != 0 {
		t.Error("Expected 0 for non-existent key")
	}

	// Test GetMetadataInt with wrong type
	p.SetMetadata("string_key", "not_an_int")
	if p.GetMetadataInt("string_key") != 0 {
		t.Error("Expected 0 for wrong type")
	}

	// Test GetMetadataBool
	p.SetMetadata("bool_key", true)
	if !p.GetMetadataBool("bool_key") {
		t.Error("Expected true for bool_key")
	}

	// Test GetMetadataBool with non-existent key
	if p.GetMetadataBool("non_existent") {
		t.Error("Expected false for non-existent key")
	}

	// Test GetMetadataBool with wrong type
	p.SetMetadata("string_key", "not_a_bool")
	if p.GetMetadataBool("string_key") {
		t.Error("Expected false for wrong type")
	}
}

func TestPromptMetadataHelpers(t *testing.T) {
	p := NewPrompt()

	// Test HasMetadata
	p.SetMetadata("test_key", "test_value")
	if !p.HasMetadata("test_key") {
		t.Error("Expected test_key to exist")
	}
	if p.HasMetadata("non_existent") {
		t.Error("Expected non_existent to not exist")
	}

	// Test GetAllMetadata
	p.SetMetadata("key1", "value1")
	p.SetMetadata("key2", 42)
	p.SetMetadata("key3", true)

	allMetadata := p.GetAllMetadata()
	if len(allMetadata) != 4 { // test_key + key1 + key2 + key3
		t.Errorf("Expected 4 metadata entries, got %d", len(allMetadata))
	}
	if allMetadata["key1"] != "value1" {
		t.Error("Expected key1 to be 'value1'")
	}

	// Test that GetAllMetadata returns a copy (mutations don't affect original)
	allMetadata["new_key"] = "new_value"
	if p.HasMetadata("new_key") {
		t.Error("Expected modifications to returned map to not affect original")
	}

	// Test DeleteMetadata
	p.DeleteMetadata("key1")
	if p.HasMetadata("key1") {
		t.Error("Expected key1 to be deleted")
	}
	if !p.HasMetadata("key2") {
		t.Error("Expected key2 to still exist")
	}

	// Test deleting non-existent key (should not panic)
	p.DeleteMetadata("non_existent")
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
			section:  Section{"Test", []Instruction{"inst1", "inst2"}, []DataBlock{}},
			expected: "Test:\n- inst1\n- inst2",
		},
		{
			name:     "Section with intro ending in colon",
			section:  Section{"Test:", []Instruction{"inst1"}, []DataBlock{}},
			expected: "Test:\n- inst1",
		},
		{
			name:     "Section without intro",
			section:  Section{"", []Instruction{"inst1", "inst2"}, []DataBlock{}},
			expected: "- inst1\n- inst2",
		},
		{
			name:     "Empty section",
			section:  Section{"", []Instruction{}, []DataBlock{}},
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
		DataBlocks: []DataBlock{},
	}

	expected := 5
	actual := section.WordsCount()

	if actual != expected {
		t.Errorf("Expected %d words, got %d", expected, actual)
	}
}

func TestSectionsString(t *testing.T) {
	sections := Sections{
		{"intro1", []Instruction{"test1"}, []DataBlock{}},
		{"intro2", []Instruction{"test2"}, []DataBlock{}},
	}

	expected := "\nintro1:\n- test1\n---\nintro2:\n- test2\n---"
	actual := sections.String()

	if actual != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}

func TestWordsCount(t *testing.T) {
	sections := []Section{
		{"intro1", []Instruction{"one two"}, []DataBlock{}},
		{"intro2", []Instruction{"three four five"}, []DataBlock{}},
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

// Tests for JSON/XML/HTML data blocks

func TestSectionAddJSONData(t *testing.T) {
	section := NewSection("API Request")
	section.AddInstruction("Use the following JSON structure")

	data := map[string]any{
		"name":  "John Doe",
		"email": "john@example.com",
		"age":   30,
	}

	err := section.AddJSONData("Example", data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(section.DataBlocks) != 1 {
		t.Errorf("Expected 1 data block, got %d", len(section.DataBlocks))
	}

	if section.DataBlocks[0].Type != "json" {
		t.Errorf("Expected type 'json', got '%s'", section.DataBlocks[0].Type)
	}

	if section.DataBlocks[0].Label != "Example" {
		t.Errorf("Expected label 'Example', got '%s'", section.DataBlocks[0].Label)
	}

	output := section.String()
	if !contains(output, "```json") {
		t.Error("Expected output to contain '```json'")
	}
	if !contains(output, "John Doe") {
		t.Error("Expected output to contain 'John Doe'")
	}
}

func TestSectionAddRawJSON(t *testing.T) {
	section := NewSection("Data Example")
	jsonStr := `{"key": "value", "number": 42}`

	section.AddRawJSON("Raw JSON", jsonStr)

	if len(section.DataBlocks) != 1 {
		t.Errorf("Expected 1 data block, got %d", len(section.DataBlocks))
	}

	if section.DataBlocks[0].Type != "json" {
		t.Errorf("Expected type 'json', got '%s'", section.DataBlocks[0].Type)
	}

	if section.DataBlocks[0].Content != jsonStr {
		t.Errorf("Expected content to match input JSON")
	}

	output := section.String()
	if !contains(output, "```json") {
		t.Error("Expected output to contain '```json'")
	}
	if !contains(output, jsonStr) {
		t.Error("Expected output to contain the JSON string")
	}
}

func TestSectionAddXMLData(t *testing.T) {
	type Person struct {
		Name  string `xml:"name"`
		Email string `xml:"email"`
		Age   int    `xml:"age"`
	}

	section := NewSection("XML Example")
	section.AddInstruction("Use this XML structure")

	person := Person{
		Name:  "Jane Smith",
		Email: "jane@example.com",
		Age:   25,
	}

	err := section.AddXMLData("Person Data", person)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(section.DataBlocks) != 1 {
		t.Errorf("Expected 1 data block, got %d", len(section.DataBlocks))
	}

	if section.DataBlocks[0].Type != "xml" {
		t.Errorf("Expected type 'xml', got '%s'", section.DataBlocks[0].Type)
	}

	output := section.String()
	if !contains(output, "```xml") {
		t.Error("Expected output to contain '```xml'")
	}
	if !contains(output, "Jane Smith") {
		t.Error("Expected output to contain 'Jane Smith'")
	}
}

func TestSectionAddRawXML(t *testing.T) {
	section := NewSection("XML Data")
	xmlStr := `<person><name>Bob</name><age>35</age></person>`

	section.AddRawXML("Sample XML", xmlStr)

	if len(section.DataBlocks) != 1 {
		t.Errorf("Expected 1 data block, got %d", len(section.DataBlocks))
	}

	if section.DataBlocks[0].Type != "xml" {
		t.Errorf("Expected type 'xml', got '%s'", section.DataBlocks[0].Type)
	}

	if section.DataBlocks[0].Content != xmlStr {
		t.Errorf("Expected content to match input XML")
	}

	output := section.String()
	if !contains(output, "```xml") {
		t.Error("Expected output to contain '```xml'")
	}
	if !contains(output, xmlStr) {
		t.Error("Expected output to contain the XML string")
	}
}

func TestSectionAddRawHTML(t *testing.T) {
	section := NewSection("HTML Template")
	htmlStr := `<div class="container"><h1>Hello World</h1><p>This is a paragraph.</p></div>`

	section.AddRawHTML("Example HTML", htmlStr)

	if len(section.DataBlocks) != 1 {
		t.Errorf("Expected 1 data block, got %d", len(section.DataBlocks))
	}

	if section.DataBlocks[0].Type != "html" {
		t.Errorf("Expected type 'html', got '%s'", section.DataBlocks[0].Type)
	}

	if section.DataBlocks[0].Content != htmlStr {
		t.Errorf("Expected content to match input HTML")
	}

	output := section.String()
	if !contains(output, "```html") {
		t.Error("Expected output to contain '```html'")
	}
	if !contains(output, htmlStr) {
		t.Error("Expected output to contain the HTML string")
	}
}

func TestSectionMultipleDataBlocks(t *testing.T) {
	section := NewSection("Multiple Formats")
	section.AddInstruction("Example with multiple data formats")

	section.AddRawJSON("JSON Example", `{"key": "value"}`)
	section.AddRawXML("XML Example", `<root><item>value</item></root>`)
	section.AddRawHTML("HTML Example", `<p>Hello</p>`)

	if len(section.DataBlocks) != 3 {
		t.Errorf("Expected 3 data blocks, got %d", len(section.DataBlocks))
	}

	output := section.String()
	if !contains(output, "```json") {
		t.Error("Expected output to contain '```json'")
	}
	if !contains(output, "```xml") {
		t.Error("Expected output to contain '```xml'")
	}
	if !contains(output, "```html") {
		t.Error("Expected output to contain '```html'")
	}
}

func TestSectionDataBlockWithoutLabel(t *testing.T) {
	section := NewSection("Test")
	section.AddRawJSON("", `{"test": true}`)

	output := section.String()
	if !contains(output, "```json") {
		t.Error("Expected output to contain '```json'")
	}
	// Should not have a label line before the code block
	if contains(output, ":\n```json") && !contains(output, "Test:\n") {
		t.Error("Should not have label line when label is empty")
	}
}

func TestSectionInstructionsAndDataBlocks(t *testing.T) {
	section := NewSection("Complete Example")
	section.AddInstruction("First instruction")
	section.AddInstruction("Second instruction")
	section.AddRawJSON("Data", `{"id": 123}`)

	output := section.String()

	expectedParts := []string{
		"Complete Example:",
		"- First instruction",
		"- Second instruction",
		"Data:",
		"```json",
		`{"id": 123}`,
		"```",
	}

	for _, part := range expectedParts {
		if !contains(output, part) {
			t.Errorf("Expected output to contain '%s'", part)
		}
	}
}

func TestPromptWithDataBlocks(t *testing.T) {
	p := NewPrompt()

	section1 := NewSection("API Example")
	section1.AddInstruction("Send a POST request")
	section1.AddRawJSON("Request Body", `{"action": "create"}`)

	section2 := NewSection("Response Example")
	section2.AddRawJSON("Expected Response", `{"status": "success"}`)

	p.AddSection(section1)
	p.AddSection(section2)

	output := p.String()

	if !contains(output, "API Example:") {
		t.Error("Expected output to contain first section intro")
	}
	if !contains(output, "Response Example:") {
		t.Error("Expected output to contain second section intro")
	}
	if !contains(output, "```json") {
		t.Error("Expected output to contain JSON code blocks")
	}
}

func TestAddJSONDataError(t *testing.T) {
	section := NewSection("Test")

	// Create a value that cannot be marshaled to JSON (channels cannot be marshaled)
	invalidData := make(chan int)

	err := section.AddJSONData("Invalid", invalidData)
	if err == nil {
		t.Error("Expected error when marshaling invalid data")
	}
}

func TestAddXMLDataError(t *testing.T) {
	section := NewSection("Test")

	// Create a value that cannot be marshaled to XML (channels cannot be marshaled)
	invalidData := make(chan int)

	err := section.AddXMLData("Invalid", invalidData)
	if err == nil {
		t.Error("Expected error when marshaling invalid data")
	}
}

// Helper function for substring checks
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
