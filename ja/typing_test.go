package ja

import (
	"reflect"
	"testing"
)

func TestTyping(t *testing.T) {
	testCases := []struct {
		keyword  string
		expected []string
	}{
		{keyword: "taikibannsei", expected: []string{"ｔ", "た", "たい", "たいｋ", "たいき", "たいきｂ", "たいきば", "たいきばｎ", "たいきばん", "たいきばんｓ", "たいきばんせ", "たいきばんせい"}},
		{keyword: "sukyatto", expected: []string{"ｓ", "す", "すｋ", "すｋｙ", "すきゃ", "すきゃｔ", "すきゃっｔ", "すきゃっと"}},
	}
	for _, tC := range testCases {
		t.Run(tC.keyword, func(t *testing.T) {
			result := Typing(tC.keyword)
			t.Log(result)
			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("\nexpected: %+v but \nresult: %+v", tC.expected, result)
			}
		})
	}
}
