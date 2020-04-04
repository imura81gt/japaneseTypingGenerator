package ja

import "testing"

func TestHiragana(t *testing.T) {
	testCases := []struct {
		keyword  string
		expected string
	}{
		{keyword: "タイキバンセイ", expected: "たいきばんせい"},
		{keyword: "スーパーグレートゴルファー", expected: "すーぱーぐれーとごるふぁー"},
		{keyword: "すもももももももものうち", expected: "すもももももももものうち"},
		// KatakanaToHiragana cannot convert belows
		{keyword: "漢字", expected: "漢字"},
		{keyword: "abc", expected: "abc"},
		{keyword: "ABC", expected: "abc"},
		{keyword: "ｂｃｄ", expected: "ｂｃｄ"},
		{keyword: "012345", expected: "012345"},
		{keyword: "０１２３４５", expected: "０１２３４５"},
		{keyword: "ﾊﾝｶｸｶﾀｶﾅ", expected: "ﾊﾝｶｸｶﾀｶﾅ"},
	}
	for _, tC := range testCases {
		t.Run(tC.keyword, func(t *testing.T) {
			result := KatakanaToHiragana(tC.keyword)
			t.Log(result)
			if result != tC.expected {
				t.Errorf("expected: %s but result: %s", tC.expected, result)
			}
		})
	}
}
