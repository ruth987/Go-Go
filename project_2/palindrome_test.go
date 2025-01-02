package wordutils

import "testing"

func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected bool
    }{
        {
            name:     "Simple palindrome",
            input:    "racecar",
            expected: true,
        },
        {
            name:     "Palindrome with spaces",
            input:    "A man a plan a canal Panama",
            expected: true,
        },
        {
            name:     "Palindrome with punctuation",
            input:    "Was it a car or a cat I saw?",
            expected: true,
        },
        {
            name:     "Not a palindrome",
            input:    "hello world",
            expected: false,
        },
        {
            name:     "Empty string",
            input:    "",
            expected: true,
        },
        {
            name:     "Single character",
            input:    "a",
            expected: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := IsPalindrome(tt.input)
            if result != tt.expected {
                t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, result, tt.expected)
            }
        })
    }
}