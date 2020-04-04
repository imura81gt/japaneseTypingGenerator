package ja

import "testing"

func TestHiraganaToHebon(t *testing.T) {
	testCases := []struct {
		keyword  string
		expected string
	}{
		{keyword: "たいきばんせい", expected: "taikibannsei"},
		{keyword: "すーぱーぐれーとごるふぁー", expected: "su-pa-gure-togorufa-"},
		{keyword: "すもももももももものうち", expected: "sumomomomomomomomonouchi"},
		{keyword: "わっしゃー", expected: "wassha-"},
		// TestHiraganaToHebon cannot convert belows
		{keyword: "漢字", expected: "漢字"},
		{keyword: "abc", expected: "abc"},
		{keyword: "ABC", expected: "ABC"},
		{keyword: "ｂｃｄ", expected: "ｂｃｄ"},
		{keyword: "012345", expected: "012345"},
		{keyword: "０１２３４５", expected: "０１２３４５"},
		{keyword: "ﾊﾝｶｸｶﾀｶﾅ", expected: "ﾊﾝｶｸｶﾀｶﾅ"},
	}
	for _, tC := range testCases {
		t.Run(tC.keyword, func(t *testing.T) {
			result := HiraganaToHebon(tC.keyword)
			t.Log(result)
			if result != tC.expected {
				t.Errorf("expected: %s but result: %s", tC.expected, result)
			}
		})
	}
}
