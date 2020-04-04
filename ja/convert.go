package ja

import (
	"strings"
	"unicode"
)

// Hiragana is convert Katakana(カタカナ) to Hiragana(ひらがな)
// https://ikawaha.hateblo.jp/entry/2015/06/26/185523
// https://www.serendip.ws/archives/6307
func KatakanaToHiragana(keyword string) string {
	var hiragana string

	kanaConv := unicode.SpecialCase{

		// ひらがなをカタカナに変換
		{
			0x3041, // Lo: ぁ
			0x3093, // Hi: ん
			[unicode.MaxCase]rune{
				0x30a1 - 0x3041, // UpperCase
				0,               // LowerCase
				0x30a1 - 0x3041, // TitleCase
			},
		},
		// カタカナをひらがなに変換
		unicode.CaseRange{
			0x30a1, // Lo: ァ
			0x30f3, // Hi: ン
			[unicode.MaxCase]rune{
				0,               // UpperCase
				0x3041 - 0x30a1, // LowerCase
				0,               // TitleCase
			},
		},
	}
	// hiragana = strings.ToLowerSpecial(kanaConv, keyword)
	hiragana = strings.ToLowerSpecial(kanaConv, keyword)

	return hiragana
}
