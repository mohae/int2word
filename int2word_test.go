package int2word

import "testing"

var tests = []struct {
	number int
	lower  string
}{
	{0, "zero"},
	{1, "one"},
	{2, "two"},
	{3, "three"},
	{4, "four"},
	{5, "five"},
	{6, "six"},
	{7, "seven"},
	{8, "eight"},
	{9, "nine"},
	{10, "ten"},
	{11, "eleven"},
	{12, "twelve"},
	{13, "thirteen"},
	{14, "fourteen"},
	{15, "fifteen"},
	{16, "sixteen"},
	{17, "seventeen"},
	{18, "eighteen"},
	{19, "nineteen"},
	{20, "twenty"},
	{21, "twenty-one"},
	{-1, "minus one"},
	{-2, "minus two"},
	{-3, "minus three"},
	{-4, "minus four"},
	{-5, "minus five"},
	{-6, "minus six"},
	{-7, "minus seven"},
	{-8, "minus eight"},
	{-9, "minus nine"},
	{-10, "minus ten"},
	{-11, "minus eleven"},
	{-12, "minus twelve"},
	{-13, "minus thirteen"},
	{-14, "minus fourteen"},
	{-15, "minus fifteen"},
	{-16, "minus sixteen"},
	{-17, "minus seventeen"},
	{-18, "minus eighteen"},
	{-19, "minus nineteen"},
	{-20, "minus twenty"},
	{-21, "minus twenty-one"},
	{22, "twenty-two"},
	{23, "twenty-three"},
	{24, "twenty-four"},
	{25, "twenty-five"},
	{26, "twenty-six"},
	{27, "twenty-seven"},
	{28, "twenty-eight"},
	{29, "twenty-nine"},
	{30, "thirty"},
	{31, "thirty-one"},
	{40, "forty"},
	{42, "forty-two"},
	{50, "fifty"},
	{53, "fifty-three"},
	{60, "sixty"},
	{64, "sixty-four"},
	{70, "seventy"},
	{75, "seventy-five"},
	{80, "eighty"},
	{86, "eighty-six"},
	{90, "ninety"},
	{97, "ninety-seven"},
	{100, "one hundred"},
	{101, "one hundred one"},
	{102, "one hundred two"},
	{103, "one hundred three"},
	{104, "one hundred four"},
	{105, "one hundred five"},
	{106, "one hundred six"},
	{107, "one hundred seven"},
	{108, "one hundred eight"},
	{109, "one hundred nine"},
	{110, "one hundred ten"},
	{111, "one hundred eleven"},
	{112, "one hundred twelve"},
	{113, "one hundred thirteen"},
	{114, "one hundred fourteen"},
	{115, "one hundred fifteen"},
	{116, "one hundred sixteen"},
	{117, "one hundred seventeen"},
	{118, "one hundred eighteen"},
	{119, "one hundred nineteen"},
	{120, "one hundred twenty"},
	{121, "one hundred twenty-one"},
	{232, "two hundred thirty-two"},
	{344, "three hundred forty-four"},
	{455, "four hundred fifty-five"},
	{566, "five hundred sixty-six"},
	{677, "six hundred seventy-seven"},
	{788, "seven hundred eighty-eight"},
	{899, "eight hundred ninety-nine"},
	{942, "nine hundred forty-two"},
	{1000, "one thousand"},
	{1001, "one thousand one"},
	{1002, "one thousand two"},
	{1003, "one thousand three"},
	{1004, "one thousand four"},
	{1005, "one thousand five"},
	{1006, "one thousand six"},
	{1007, "one thousand seven"},
	{1008, "one thousand eight"},
	{1009, "one thousand nine"},
	{1010, "one thousand ten"},
	{1011, "one thousand eleven"},
	{1012, "one thousand twelve"},
	{1013, "one thousand thirteen"},
	{1014, "one thousand fourteen"},
	{1015, "one thousand fifteen"},
	{1016, "one thousand sixteen"},
	{1017, "one thousand seventeen"},
	{1018, "one thousand eighteen"},
	{1019, "one thousand nineteen"},
	{1020, "one thousand twenty"},
	{2121, "two thousand one hundred twenty-one"},
	{3232, "three thousand two hundred thirty-two"},
	{4342, "four thousand three hundred forty-two"},
	{5455, "five thousand four hundred fifty-five"},
	{6342, "six thousand three hundred forty-two"},
	{7765, "seven thousand seven hundred sixty-five"},
	{8567, "eight thousand five hundred sixty-seven"},
	{9065, "nine thousand sixty-five"},
	{10001, "ten thousand one"},
	{10010, "ten thousand ten"},
	{10011, "ten thousand eleven"},
	{20000, "twenty thousand"},
	{20011, "twenty thousand eleven"},
	{20103, "twenty thousand one hundred three"},
	{21230, "twenty-one thousand two hundred thirty"},
	{23400, "twenty-three thousand four hundred"},
	{34642, "thirty-four thousand six hundred forty-two"},
	{41232, "forty-one thousand two hundred thirty-two"},
	{56716, "fifty-six thousand seven hundred sixteen"},
	{68350, "sixty-eight thousand three hundred fifty"},
	{79021, "seventy-nine thousand twenty-one"},
	{88327, "eighty-eight thousand three hundred twenty-seven"},
	{97642, "ninety-seven thousand six hundred forty-two"},
	{3203120, "three million two hundred three thousand one hundred twenty"},
	{10000001, "ten million one"},
	{643000352, "six hundred forty-three million three hundred fifty-two"},
	{12387459201, "twelve billion three hundred eighty-seven million four hundred fifty-nine thousand two hundred one"},
	{94740000498920, "ninety-four trillion seven hundred forty billion four hundred ninety-eight thousand nine hundred twenty"},
}

func TestLower(t *testing.T) {
	for _, test := range tests {
		s := Lower(test.number)
		if s != test.lower {
			t.Errorf("%d: got %q want %q", test.number, s, test.lower)
		}
	}
}

func TestHundreds(t *testing.T) {
	hTests := []struct {
		v int
		s string
	}{
		{0, ""},
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
		{6, "six"},
		{7, "seven"},
		{8, "eight"},
		{9, "nine"},
		{10, "ten"},
		{11, "eleven"},
		{12, "twelve"},
		{13, "thirteen"},
		{14, "fourteen"},
		{15, "fifteen"},
		{16, "sixteen"},
		{17, "seventeen"},
		{18, "eighteen"},
		{19, "nineteen"},
		{20, "twenty"},
		{21, "twenty-one"},
		{22, "twenty-two"},
		{23, "twenty-three"},
		{24, "twenty-four"},
		{25, "twenty-five"},
		{26, "twenty-six"},
		{27, "twenty-seven"},
		{28, "twenty-eight"},
		{29, "twenty-nine"},
		{30, "thirty"},
		{31, "thirty-one"},
		{40, "forty"},
		{42, "forty-two"},
		{50, "fifty"},
		{53, "fifty-three"},
		{60, "sixty"},
		{64, "sixty-four"},
		{70, "seventy"},
		{75, "seventy-five"},
		{80, "eighty"},
		{86, "eighty-six"},
		{90, "ninety"},
		{97, "ninety-seven"},
		{100, "one hundred"},
		{101, "one hundred one"},
		{102, "one hundred two"},
		{103, "one hundred three"},
		{104, "one hundred four"},
		{105, "one hundred five"},
		{106, "one hundred six"},
		{107, "one hundred seven"},
		{108, "one hundred eight"},
		{109, "one hundred nine"},
		{110, "one hundred ten"},
		{111, "one hundred eleven"},
		{112, "one hundred twelve"},
		{113, "one hundred thirteen"},
		{114, "one hundred fourteen"},
		{115, "one hundred fifteen"},
		{116, "one hundred sixteen"},
		{117, "one hundred seventeen"},
		{118, "one hundred eighteen"},
		{119, "one hundred nineteen"},
		{120, "one hundred twenty"},
		{121, "one hundred twenty-one"},
		{232, "two hundred thirty-two"},
		{344, "three hundred forty-four"},
		{405, "four hundred five"},
		{411, "four hundred eleven"},
		{442, "four hundred forty-two"},
		{450, "four hundred fifty"},
		{566, "five hundred sixty-six"},
		{677, "six hundred seventy-seven"},
		{788, "seven hundred eighty-eight"},
		{899, "eight hundred ninety-nine"},
		{942, "nine hundred forty-two"},
	}
	for _, test := range hTests {
		s := hundreds(test.v)
		if s != test.s {
			t.Errorf("got %q; want %q", s, test.s)
		}
	}
}
