package scrabble

import "strings"
import "unicode"
import "regexp"

func CapitalUse(word string) bool {
    if len(word) < 1 {
        return false
    }
    allUpper := regexp.MustCompile(`^[A-Z]+$`)
    allLower := regexp.MustCompile(`^[a-z]+$`)
    capitalized := regexp.MustCompile(`^[A-Z][a-z]+$`)
    

    return capitalized.MatchString(word) || allUpper.MatchString(word) || allLower.MatchString(word)
}


func CapitalUse2(word string) bool {
    if len(word) < 1 {
        return false
    }
    
    if len(word) == 1 {
        return unicode.IsUpper(rune(word[0])) || unicode.IsLower(rune(word[0])) 
    }
    
    if strings.ToUpper(word) == word || strings.ToLower(word) == word {
        return true
    }
        
    return unicode.IsUpper(rune(word[0])) && strings.ToLower(word[1:]) == word[1:]
}

