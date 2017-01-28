// Copyright (c) 2016, Joel Scoble. All rights reserved.
//
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package int2word converts integers to their word equivelants.
package int2word

import "strings"

const negative = "minus "

var (
	small = [...]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}
	tens  = [...]string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	scale = [...]string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion"}
)

// Lower returns the word representation of an int as words. This is an alias to
// Word
func Lower(i int) string {
	return Word(i)
}

// Title returns the word representation of an int using title case.
func Title(i int) string {
	return strings.ToTitle(Word(i))
}

// Capitalized
func Capitalized(i int) string {
	s := Word(i)
	return string(s[0]-0x20) + s[1:] // make the first char UPPER
}

func Word(v int) string {
	// shortcut on small numbers
	if v == 0 {
		return "zero"
	}

	var neg bool
	if v < 0 {
		neg = true
		v = -v
	}

	var parts []int
	var r int
	// break into chunks of 3 digits
	for v > 0 {
		r = v % 1000
		v = v / 1000
		parts = append(parts, r)
	}

	var words string
	// If there is just 1 part, process and return.
	if len(parts) == 1 {
		words = hundreds(parts[0])
		if neg {
			words = "minus " + words
		}
		return words
	}

	// Otherwise build the string by processing the parts in reverse order since
	// the slice parts are right to left.
	//words := make([]string, 0, len(parts))
	l := len(parts) - 1
	for i := l; i >= 0; i-- {
		s := hundreds(parts[i])
		if s == "" {
			continue
		}
		/*		if scale[i] == "" {
					words = append(words, s)
					continue
				}
		*/
		if i == l {
			if neg {
				words = "minus " + s + " " + scale[i]
			} else {
				words = s + " " + scale[i]
			}
			continue
		}
		if scale[i] == "" {
			words += " " + s
		} else {
			words += " " + s + " " + scale[i]
		}

		//		words = append(words, fmt.Sprintf("%s %s", s, scale[i]))
	}
	//	return strings.Join(words, " ")
	return words
}

func hundreds(v int) string {
	// If it's <= 20, just return the value.
	if v <= 20 {
		return small[v]
	}

	// Break the number up to their individual digits.
	var r int8
	parts := make([]int8, 0, 3)
	var s string
	for v > 0 {
		r = int8(v % 10)
		v = v / 10
		parts = append(parts, r)
	}
	var hasTens bool
	// Process in reverse order since the parts are currently right to left.
	for i := len(parts) - 1; i >= 0; i-- {
		switch i {
		case 2: // hundreds part
			if parts[i] == 0 { // skip if its 0
				continue
			}
			s = small[parts[i]] + " hundred"
		case 1: // tens parts
			switch parts[i] {
			case 0:
				continue
			case 1: // If this a teens number treat it like one.
				i = i*10 + int(parts[0])
				if s == "" {
					s = small[i]
				} else {
					s += " " + small[i]
				}
				return s
			}
			if s == "" { // Add the tens part.
				s = tens[parts[i]]
			} else {
				s += " " + tens[parts[i]]
			}
			hasTens = true
		case 0:
			if parts[i] == 0 {
				continue
			}
			if hasTens {
				s += "-" + small[parts[i]] // Add the ones part.
			} else {
				s += " " + small[parts[i]]
			}
		}
	}
	return s
}
