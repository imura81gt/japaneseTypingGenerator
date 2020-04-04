package ja

import (
	"github.com/ikawaha/kagome/tokenizer"
)

// Yomigana is return Japanese Yomigana
func Yomigana(keyword string) string {
	katakana := KatakanaYomigana(keyword)
	if katakana == "*" {
		katakana = keyword
	}
	hiragana := KatakanaToHiragana(katakana)
	return hiragana
}

// KatakanaYomigana is return Japanese KatakanaYomigana(カタカナ)
func KatakanaYomigana(keyword string) string {
	var yomigana string

	t := tokenizer.New()
	tokens := t.Analyze(keyword, tokenizer.Search)

	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			// BOS: Begin Of Sentence, EOS: End Of Sentence.
			continue
		}
		features := token.Features()
		yomi := features[len(features)-2]
		if yomi == "*" {
			yomi = token.Surface
		}
		yomigana = yomigana + yomi
	}

	return yomigana
}
