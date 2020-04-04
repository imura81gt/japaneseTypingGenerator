package ja

import "testing"

func TestYomigana(t *testing.T) {
	testCases := []struct {
		keyword  string
		expected string
	}{
		{keyword: "大器晩成", expected: "たいきばんせい"},
		{keyword: "スーパーグレートゴルファー", expected: "すーぱーぐれーとごるふぁー"},
		{keyword: "すもももももももものうち", expected: "すもももももももものうち"},
	}
	for _, tC := range testCases {
		t.Run(tC.keyword, func(t *testing.T) {
			result := Yomigana(tC.keyword)
			t.Log(result)
			if result != tC.expected {
				t.Errorf("expected: %s but result: %s", tC.expected, result)
			}
		})
	}
}

func TestKatakanaYomigana(t *testing.T) {
	testCases := []struct {
		keyword  string
		expected string
	}{
		{keyword: "大器晩成", expected: "タイキバンセイ"},
		{keyword: "スーパーグレートゴルファー", expected: "スーパーグレートゴルファー"},
		{keyword: "すもももももももものうち", expected: "スモモモモモモモモノウチ"},
		{keyword: "防風", expected: "ボウフウ"},
		// {keyword: "築地銀だこ", expected: "ツキジギンダコ"}, // 商標(?) need to use user dictionary: 築(チク) / 地銀(チギン) / だこ(ダコ)
		// {keyword: "錦織圭", expected: "ニシコリケイ""}, // 人名 need to use user dictionary: 錦織(ニシキオリ) / 圭(ケイ)
	}
	for _, tC := range testCases {
		t.Run(tC.keyword, func(t *testing.T) {
			result := KatakanaYomigana(tC.keyword)
			t.Log(result)
			if result != tC.expected {
				t.Errorf("expected: %s but result: %s", tC.expected, result)
			}
		})
	}
}
