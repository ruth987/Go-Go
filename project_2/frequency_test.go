package wordutils

// import and use GetWordFrequency


import (
    "reflect"
    "testing"
)

func TestGetWordFrequency(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected map[string]int
    }{
        {
            name:  "Basic test",
            input: "The quick brown fox jumps over the lazy dog",
            expected: map[string]int{
                "the":   2,
                "quick": 1,
                "brown": 1,
                "fox":   1,
                "jumps": 1,
                "over":  1,
                "lazy":  1,
                "dog":   1,
            },
        },
        {
            name:  "With punctuation",
            input: "Hello, hello! How are you?",
            expected: map[string]int{
                "hello": 2,
                "how":   1,
                "are":   1,
                "you":   1,
            },
        },
        {
            name:     "Empty string",
            input:    "",
            expected: map[string]int{},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := GetWordFrequency(tt.input)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("GetWordFrequency() = %v, want %v", result, tt.expected)
            }
        })
    }
}
