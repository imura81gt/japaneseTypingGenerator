package ja

import (
	"strings"
	"unicode"

	"golang.org/x/text/width"
)

func Do(keyword string) []string {
	var ks = strings.Fields(keyword)
	var tl = []string{}
	var te = ""

	for _, k := range ks {
		var y = Yomigana(k)
		var h = HiraganaToHebon(y)

		if te == "" {
			for _, t := range Typing(h) {
				tl = append(tl, t)
			}
			tl = append(tl, k)
			te = k
		} else {
			for _, t := range Typing(h) {
				tl = append(tl, te+"　"+t)
			}
			tl = append(tl, te+"　"+k)
			te = te + "　" + k

		}
	}
	// var t = Typing(h)
	return tl
}

// Typing is return []string like Typed Japanese
func Typing(hebon string) []string {
	typing := []string{}
	hebonTable := NewHebonTable()

	ch := ""
	for _, c := range hebon {
		bch := ch
		if unicode.IsSpace(c) {
			if len(typing) == 0 {
				typing = append(typing, Widen(string(c)))
			} else {
				typing = append(typing, strings.Replace(typing[len(typing)-1], Widen(bch), "", 1)+Widen(string(c)))
			}
			ch = ""
			continue
		}
		if bch == string(c) && string(c) != "n" {
			if len(typing) == 0 {
				typing = append(typing, "っ"+Widen(string(c)))
			} else {
				typing = append(typing, strings.Replace(typing[len(typing)-1], Widen(bch), "", 1)+"っ"+Widen(string(c)))
			}
			continue
		}
		ch = ch + string(c)

		for _, v := range hebonTable.ThreeCharactors {
			if ch == v.Hebon {
				if len(typing) == 0 {
					typing = append(typing, v.Hiragana)
				} else {
					typing = append(typing, strings.Replace(typing[len(typing)-1], Widen(bch), "", 1)+v.Hiragana)
				}
				ch = ""
				continue
			}
		}

		for _, v := range hebonTable.TwoCharactors {
			if ch == v.Hebon {
				if len(typing) == 0 {
					typing = append(typing, v.Hiragana)
				} else {
					typing = append(typing, strings.Replace(typing[len(typing)-1], Widen(bch), "", 1)+v.Hiragana)
				}
				ch = ""
				continue
			}
		}

		for _, v := range hebonTable.OneCharactor {
			if ch == v.Hebon {
				if len(typing) == 0 {
					typing = append(typing, v.Hiragana)
				} else {
					typing = append(typing, strings.Replace(typing[len(typing)-1], Widen(bch), "", 1)+v.Hiragana)
				}
				ch = ""
				continue
			}
		}

		if ch != "" {
			if len(typing) == 0 {
				typing = append(typing, Widen(string(ch)))
			} else {
				typing = append(typing, strings.Replace(typing[len(typing)-1], Widen(bch), "", 1)+Widen(string(ch)))
			}
		}
	}
	return typing
}

func Widen(keyword string) string {
	return width.Widen.String(keyword)
}
