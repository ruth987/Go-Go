package wordutils

import (
    "strings"
    "unicode"
)

func IsPalindrome(text string) bool {
    text = strings.ToLower(text)
    
    var cleaned strings.Builder
    for _, r := range text {
        if unicode.IsLetter(r) || unicode.IsNumber(r) {
            cleaned.WriteRune(r)
        }
    }
    
    cleanedText := cleaned.String()
    length := len(cleanedText)
    
    for i := 0; i < length/2; i++ {
        if cleanedText[i] != cleanedText[length-1-i] {
            return false
        }
    }
    
    return true
}