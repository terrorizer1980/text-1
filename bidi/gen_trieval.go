// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

type class int

const (
	classLeftToRight        class = iota // Class L
	classRightToLeft                     // Class R
	classEuropeanNumber                  // Class EN
	classEuropeanSeparator               // Class ES
	classEuropeanTerminator              // Class ET
	classArabicNumber                    // Class AN
	classCommonSeparator                 // Class CS
	classParagraphSeparator              // Class B
	classSegmentSeparator                // Class S
	classWhiteSpace                      // Class WS
	classOtherNeutral                    // Class ON
	classBoundaryNeutral                 // Class BN
	classNonspacingMark                  // Class NSM
	classArabicLetter                    // Class AL
	classControl                         // Control LRO - PDI

	numClass

	classLeftToRightOverride   // Class LRO
	classRightToLeftOverride   // Class RLO
	classLeftToRightEmbedding  // Class LRE
	classRightToLeftEmbedding  // Class RLE
	classPopDirectionalFormat  // Class PDF
	classLeftToRightIsolate    // Class LRI
	classRightToLeftIsolate    // Class RLI
	classFirstStrongIsolate    // Class FSI
	classPopDirectionalIsolate // Class PDI
)

var controlToClass = map[rune]class{
	0x202D: classLeftToRightOverride,
	0x202E: classRightToLeftOverride,
	0x202A: classLeftToRightEmbedding,
	0x202B: classRightToLeftEmbedding,
	0x202C: classPopDirectionalFormat,
	0x2066: classLeftToRightIsolate,
	0x2067: classRightToLeftIsolate,
	0x2068: classFirstStrongIsolate,
	0x2069: classPopDirectionalIsolate,
}

// A trie entry has the following bits:
// 7..5  XOR mask for brackets
// 4     1: Bracket open, 0: Bracket close
// 3..0  class type
type entry uint8

const (
	openMask     = 0x10
	xorMaskShift = 5
)

func (e entry) isBracket() bool            { return e&0xF0 != 0 }
func (e entry) isOpen() bool               { return e&openMask != 0 }
func (e entry) reverseBracket(r rune) rune { return xorMasks[e>>xorMaskShift] ^ r }
func (e entry) class(r rune) class {
	c := class(e & 0x0F)
	if c == classControl {
		return controlToClass[r]
	}
	return c
}