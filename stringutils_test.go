// Package goarabic contains utility functions for working with strings.
package goarabic

import (
	"testing"
	"unicode/utf8"
)

// Reverse returns its argument string reversed rune-wise left to right.
func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Crowdbotics", "scitobdworC"},
		{"Hello, 世界", "界世 ,olleH"},
		{"نص عربي", "يبرع صن"},
		{"نَصٌ عَربِيٌّ", "ٌّيِبرَع ٌصَن"},
		{"نَصٌ عَربِيٌّ!", "!ٌّيِبرَع ٌصَن"},
		{"نَصٌ example, عَربِيٌّ!", "!ٌّيِبرَع ,elpmaxe ٌصَن"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestRemoveTashkeel(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"نَصٌ عَربِيٌّ", "نص عربي"},
		{"نص عربي", "نص عربي"},
		{"", ""},
	}
	for _, c := range cases {
		got := RemoveTashkeel(c.in)
		if got != c.want {
			t.Errorf("RemoveTashkeel(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSmartLength(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"نَصٌ عَربِيٌّ", 7},
		{"نص عربي", 7},
		{"Hello, world", 12},
		{"Hello, 世界", 9},
		{"", 0},
	}
	for _, c := range cases {
		got := SmartLength(&c.in)
		if got != c.want {
			t.Errorf("SmartLength(...) got %d, want %d", got, c.want)
		}
	}
}

func TestToGlyph(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"تجربة النص العربي", "\ufe97\ufea0\ufeae\ufe91\ufe94 \u0627\ufedf\ufee8\ufeba \u0627\ufedf\ufecc\ufeae\ufe91\ufef2"},
		{"", ""},
	}
	for _, c := range cases {
		got := ToGlyph(c.in)
		if got != c.want {
			t.Errorf("ToGlyph(...) got %q, want %+q", got, c.want)
		}
	}
}

func TestRemoveTatweel(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"نـــص عــربــي", "نص عربي"},
		{"نـــَصٌ عَـربي", "نَصٌ عَربي"},
		{"", ""},
	}
	for _, c := range cases {
		got := RemoveTatweel(c.in)
		if got != c.want {
			t.Errorf("RemoveTatweel(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestRemoveAllNonArabicChars(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"عــربـabcـينـــص", "عــربــينـــص"},
		{"عــربــينwo%%rd_ـــصa", "عــربــينـــص"},
		{"", ""},
	}
	for _, c := range cases {
		got := RemoveAllNonArabicChars(c.in)
		if got != c.want {
			t.Errorf("RemoveAllNonArabicChars(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSafeBWToAr(t *testing.T) {
	ans := "هَذَا بَيْتٌ"
	ar, err := SafeBWToAr("haVaA bayotN")
	if err != nil {
		t.Errorf("SafeBWToAr(\"haVaA bayotN\") raised error: %v", err)
	} else if !utf8.ValidString(ar) {
		t.Error("SafeBWToAr(\"haVaA bayotN\") is an invalid string")
	} else if ar != ans {
		t.Errorf("SafeBWToAr(\"haVaA bayotN\") = (%s); want (%s)", ar, ans)
		t.Error("Returned result:\n")
		for _, v := range ar {
			t.Errorf("%U\n", v)
		}
		t.Error("Answer:\n")
		for _, v := range ans {
			t.Errorf("%U\n", v)
		}
	}
}
