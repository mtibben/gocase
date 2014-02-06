// Package gocase converts strings between different casing schemes
package gocase

import "unicode"

var noop = func(a rune) rune { return a }

func snaker(s string, wordSeparator rune, firstRune func(rune) rune, firstRuneOfWord func(rune) rune, otherRunes func(rune) rune) string {
	useWordSeperator := wordSeparator != rune(0)
	newS := []rune{}

	// pops a rune off newS
	lastRuneIsWordSeparator := func() bool {
		if len(newS) > 0 {
			return newS[len(newS)-1] == wordSeparator
		}
		return true
	}

	prev := wordSeparator
	for _, cur := range s {
		isWordBoundary := (unicode.IsLower(prev) && unicode.IsUpper(cur)) || unicode.IsSpace(prev)

		if !unicode.IsLetter(cur) {
			// ignore
		} else if isWordBoundary {
			if useWordSeperator && !lastRuneIsWordSeparator() {
				newS = append(newS, wordSeparator)
			}
			newS = append(newS, firstRuneOfWord(cur))
		} else {
			newS = append(newS, otherRunes(cur))
		}

		prev = cur
	}

	if len(newS) > 0 {
		newS[0] = firstRune(newS[0])
	}

	return string(newS)
}

// ToSnake converts strings to snake_case.
// In Snake case, punctuation is ignored and spaces are replaced by single underscores.
//
// Example: may_the_force_be_with_you
func ToSnake(s string) string {
	return snaker(s, '_', unicode.ToLower, unicode.ToLower, unicode.ToLower)
}

// ToUpperSnake converts strings to uppercase snake case.
// Also known as SCREAMING_SNAKE_CASE.
//
// Example: MAY_THE_FORCE_BE_WITH_YOU
func ToUpperSnake(s string) string {
	return snaker(s, '_', unicode.ToUpper, unicode.ToUpper, unicode.ToUpper)
}

// ToSpinal converts strings to spinal case.
// Spinal case is similar to snake case, except it uses a `-` instead of `_`.
//
// Example: may-the-force-be-with-you
func ToSpinal(s string) string {
	return snaker(s, '-', unicode.ToLower, unicode.ToLower, unicode.ToLower)
}

// ToTrain converts strings to train case.
// Train case is similar to spinal case, except each word is capitalised.
//
// Example: May-The-Force-Be-With-You
func ToTrain(s string) string {
	return snaker(s, '-', unicode.ToUpper, unicode.ToUpper, noop)
}

// ToCamel converts strings to CamelCase.
// In CamelCase, the first letter of each word is capitalized, spaces and punctuation removed.
// Also known as "upper camel case", "Pascal case" or "Bumpy case".
//
// Example: MayTheForceBeWithYou
func ToCamel(s string) string {
	return snaker(s, rune(0), unicode.ToUpper, unicode.ToUpper, noop)
}

// ToLowerCamel converts strings to lower camel case
// Lower camel case describes a variation to CamelCase, as in "camelCase" (or "iPod" or "eBay"),
// in which the very first letter is in lower case.
//
// Example: mayTheForceBeWithYou
func ToLowerCamel(s string) string {
	return snaker(s, rune(0), unicode.ToLower, unicode.ToUpper, noop)
}
