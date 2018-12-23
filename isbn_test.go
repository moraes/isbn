// Copyright 2015 Rodrigo Moraes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package isbn

import (
	"testing"
)

type book struct {
	isbn10 string
	isbn13 string
	valid  bool
}

var books = []book{
	// Calvin and Hobbes, 1987
	{"0836220889", "9780836220889", true},
	// Something Under the Bed Is Drooling, 1988
	{"0836218256", "9780836218251", true},
	// Yukon Ho!, 1989
	{"0836218353", "9780836218350", true},
	// Weirdos from Another Planet!, 1990
	{"1449407102", "9781449407100", true},
	// Scientific Progress Goes 'Boink', 1991
	{"0836218787", "9780836218787", true},
	// Attack of the Deranged Mutant Killer Monster Snow Goons, 1992
	{"0836218833", "9780836218831", true},
	// The Days are Just Packed, 1993
	{"0836217357", "9780836217353", true},
	// Invalid: too many characters
	{"08362208891", "97808362208891", false},
	{"08362182562", "97808362182512", false},
	{"08362183533", "97808362183503", false},
	{"08362186204", "97804391374924", false},
	{"08362187875", "97808362187875", false},
	{"08362188336", "97808362188316", false},
	{"08362173577", "97808362173537", false},
	// Invalid: too few characters
	{"083622088", "978083622088", false},
	{"083621825", "978083621825", false},
	{"083621835", "978083621835", false},
	{"083621862", "978043913749", false},
	{"083621878", "978083621878", false},
	{"083621883", "978083621883", false},
	{"083621735", "978083621735", false},
	// Invalid: bad check digit
	{"0836220888", "9780836220880", false},
	{"0836218255", "9780836218252", false},
	{"0836218352", "9780836218351", false},
	{"0836218629", "9780439137493", false},
	{"0836218786", "9780836218788", false},
	{"0836218832", "9780836218832", false},
	{"0836217356", "9780836217354", false},
}

func TestISBN(t *testing.T) {
	for _, v := range books {
		shouldbe, shouldnotbe := "valid", "invalid"
		if v.valid {
			d10, err := CheckDigit10(v.isbn10)
			if err != nil || d10 != v.isbn10[len(v.isbn10)-1:] {
				t.Errorf("CheckDigit10: failed to calculate check digit for %s: got %s, expected %s (error: %v)", v.isbn10, d10, v.isbn10[len(v.isbn10)-1:], err)
			}
			d13, err := CheckDigit13(v.isbn13)
			if err != nil || d13 != v.isbn13[len(v.isbn13)-1:] {
				t.Errorf("CheckDigit13: failed to calculate check digit for %s: got %s, expected %s (error: %v)", v.isbn13, d13, v.isbn13[len(v.isbn13)-1:], err)
			}
			to13, err := To13(v.isbn10)
			if err != nil || to13 != v.isbn13 {
				t.Errorf("To13: failed to convert %s from ISBN-10 to ISBN-13: got %s, expected %s (error: %v)", v.isbn10, to13, v.isbn13, err)
			}
		} else {
			shouldbe, shouldnotbe = "invalid", "valid"
		}
		if Validate(v.isbn10) != v.valid {
			t.Errorf("Validate: %s should be %s, got %s", v.isbn10, shouldbe, shouldnotbe)
		}
		if Validate(v.isbn13) != v.valid {
			t.Errorf("Validate: %s should be %s, got %s", v.isbn13, shouldbe, shouldnotbe)
		}
		if Validate10(v.isbn10) != v.valid {
			t.Errorf("Validate10: %s should be %s, got %s", v.isbn10, shouldbe, shouldnotbe)
		}
		if Validate13(v.isbn13) != v.valid {
			t.Errorf("Validate13: %s should be %s, got %s", v.isbn13, shouldbe, shouldnotbe)
		}
	}
}
