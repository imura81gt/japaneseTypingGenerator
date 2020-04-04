package ja

import (
	"strings"
)

// type TypingInformation map[string]Charactors

// type NumberOfCharactors int

// const (
// 	Unknown NumberOfCharactors = iota
// 	OneCharactor
// 	TowCharactors
// 	ThreeCharactors
// )

type HebonTable struct {
	ThreeCharactors []HebonPair
	TwoCharactors   []HebonPair
	OneCharactor    []HebonPair
}

type HebonPair struct {
	Hiragana string
	Hebon    string
}

func NewHebonTable() *HebonTable {
	h := HebonTable{
		ThreeCharactors: []HebonPair{
			{Hiragana: "っきゃ", Hebon: "kkya"},
			{Hiragana: "っきぃ", Hebon: "kkyi"},
			{Hiragana: "っきゅ", Hebon: "kkyu"},
			{Hiragana: "っきぇ", Hebon: "kkye"},
			{Hiragana: "っきょ", Hebon: "kkyo"},
			{Hiragana: "っぎゃ", Hebon: "ggya"},
			{Hiragana: "っぎぃ", Hebon: "ggyi"},
			{Hiragana: "っぎゅ", Hebon: "ggyu"},
			{Hiragana: "っぎぇ", Hebon: "ggye"},
			{Hiragana: "っぎょ", Hebon: "ggyo"},

			{Hiragana: "っしゃ", Hebon: "ssha"},
			{Hiragana: "っしぃ", Hebon: "ssyi"},
			{Hiragana: "っしゅ", Hebon: "sshu"},
			{Hiragana: "っしぇ", Hebon: "sshe"},
			{Hiragana: "っしょ", Hebon: "ssho"},
			{Hiragana: "っじゃ", Hebon: "jja"},
			{Hiragana: "っじぃ", Hebon: "jjyi"},
			{Hiragana: "っじゅ", Hebon: "jju"},
			{Hiragana: "っじぇ", Hebon: "jje"},
			{Hiragana: "っじょ", Hebon: "jjo"},

			{Hiragana: "っちゃ", Hebon: "tcha"},
			{Hiragana: "っちぃ", Hebon: "ttyi"},
			{Hiragana: "っちゅ", Hebon: "tchu"},
			{Hiragana: "っちぇ", Hebon: "tche"},
			{Hiragana: "っちょ", Hebon: "tcho"},
			{Hiragana: "っぢゃ", Hebon: "ddya"},
			{Hiragana: "っぢぃ", Hebon: "ddyi"},
			{Hiragana: "っぢゅ", Hebon: "ddyu"},
			{Hiragana: "っぢぇ", Hebon: "ddye"},
			{Hiragana: "っぢょ", Hebon: "ddyo"},

			{Hiragana: "っひゃ", Hebon: "hhya"},
			{Hiragana: "っひい", Hebon: "hhyi"},
			{Hiragana: "っひゅ", Hebon: "hhyu"},
			{Hiragana: "っひぇ", Hebon: "hhye"},
			{Hiragana: "っひょ", Hebon: "hhyo"},
			{Hiragana: "っびゃ", Hebon: "bbya"},
			{Hiragana: "っびぃ", Hebon: "bbyi"},
			{Hiragana: "っびゅ", Hebon: "bbyu"},
			{Hiragana: "っびぇ", Hebon: "bbye"},
			{Hiragana: "っびょ", Hebon: "bbyo"},
			{Hiragana: "っぴゃ", Hebon: "ppya"},
			{Hiragana: "っぴぃ", Hebon: "ppyi"},
			{Hiragana: "っぴゅ", Hebon: "ppyu"},
			{Hiragana: "っぴぇ", Hebon: "ppye"},
			{Hiragana: "っぴょ", Hebon: "ppyo"},

			{Hiragana: "っふぁ", Hebon: "ffa"},
			{Hiragana: "っふぃ", Hebon: "ffi"},
			{Hiragana: "っふぇ", Hebon: "ffe"},
			{Hiragana: "っふぉ", Hebon: "ffo"},

			{Hiragana: "っみゃ", Hebon: "mmya"},
			{Hiragana: "っみぃ", Hebon: "mmyi"},
			{Hiragana: "っみゅ", Hebon: "mmyu"},
			{Hiragana: "っみぇ", Hebon: "mmye"},
			{Hiragana: "っみょ", Hebon: "mmyo"},

			{Hiragana: "っりゃ", Hebon: "rrya"},
			{Hiragana: "っりぃ", Hebon: "rryi"},
			{Hiragana: "っりゅ", Hebon: "rryu"},
			{Hiragana: "っりぇ", Hebon: "rrye"},
			{Hiragana: "っりょ", Hebon: "rryo"},
		},
		TwoCharactors: []HebonPair{

			{Hiragana: "きゃ", Hebon: "kya"},
			{Hiragana: "きぃ", Hebon: "kyi"},
			{Hiragana: "きゅ", Hebon: "kyu"},
			{Hiragana: "きぇ", Hebon: "kye"},
			{Hiragana: "きょ", Hebon: "kyo"},
			{Hiragana: "ぎゃ", Hebon: "gya"},
			{Hiragana: "ぎぃ", Hebon: "gyi"},
			{Hiragana: "ぎゅ", Hebon: "gyu"},
			{Hiragana: "ぎぇ", Hebon: "gye"},
			{Hiragana: "ぎょ", Hebon: "gyo"},

			{Hiragana: "しゃ", Hebon: "sha"},
			{Hiragana: "しぃ", Hebon: "syi"},
			{Hiragana: "しゅ", Hebon: "shu"},
			{Hiragana: "しぇ", Hebon: "she"},
			{Hiragana: "しょ", Hebon: "sho"},
			{Hiragana: "じゃ", Hebon: "ja"},
			{Hiragana: "じぃ", Hebon: "zyi"},
			{Hiragana: "じゅ", Hebon: "ju"},
			{Hiragana: "じぇ", Hebon: "je"},
			{Hiragana: "じょ", Hebon: "jo"},

			{Hiragana: "ちゃ", Hebon: "cha"},
			{Hiragana: "ちぃ", Hebon: "tyi"},
			{Hiragana: "ちゅ", Hebon: "chu"},
			{Hiragana: "ちぇ", Hebon: "che"},
			{Hiragana: "ちょ", Hebon: "cho"},
			{Hiragana: "ぢゃ", Hebon: "dya"},
			{Hiragana: "ぢぃ", Hebon: "dyi"},
			{Hiragana: "ぢゅ", Hebon: "dyu"},
			{Hiragana: "ぢぇ", Hebon: "dye"},
			{Hiragana: "ぢょ", Hebon: "dyo"},

			{Hiragana: "てゃ", Hebon: "tha"},
			{Hiragana: "てぃ", Hebon: "thi"},
			{Hiragana: "てゅ", Hebon: "thu"},
			{Hiragana: "てぇ", Hebon: "the"},
			{Hiragana: "てょ", Hebon: "tho"},

			{Hiragana: "にゃ", Hebon: "nya"},
			{Hiragana: "にぃ", Hebon: "nyi"},
			{Hiragana: "にゅ", Hebon: "nyu"},
			{Hiragana: "にぇ", Hebon: "nye"},
			{Hiragana: "にょ", Hebon: "nyo"},

			{Hiragana: "ひゃ", Hebon: "hya"},
			{Hiragana: "ひぃ", Hebon: "hyi"},
			{Hiragana: "ひゅ", Hebon: "hyu"},
			{Hiragana: "ひぇ", Hebon: "hye"},
			{Hiragana: "ひょ", Hebon: "hyo"},
			{Hiragana: "びゃ", Hebon: "bya"},
			{Hiragana: "びぃ", Hebon: "byi"},
			{Hiragana: "びゅ", Hebon: "byu"},
			{Hiragana: "びぇ", Hebon: "bye"},
			{Hiragana: "びょ", Hebon: "byo"},
			{Hiragana: "ぴゃ", Hebon: "pya"},
			{Hiragana: "ぴぃ", Hebon: "pyi"},
			{Hiragana: "ぴゅ", Hebon: "pyu"},
			{Hiragana: "ぴぇ", Hebon: "pye"},
			{Hiragana: "ぴょ", Hebon: "pyo"},

			{Hiragana: "みゃ", Hebon: "mya"},
			{Hiragana: "みぃ", Hebon: "myi"},
			{Hiragana: "みゅ", Hebon: "myu"},
			{Hiragana: "みぇ", Hebon: "mye"},
			{Hiragana: "みょ", Hebon: "myo"},

			{Hiragana: "りゃ", Hebon: "rya"},
			{Hiragana: "りぃ", Hebon: "ryi"},
			{Hiragana: "りゅ", Hebon: "ryu"},
			{Hiragana: "りぇ", Hebon: "rye"},
			{Hiragana: "りょ", Hebon: "ryo"},

			{Hiragana: "うぃ", Hebon: "wi"},
			{Hiragana: "うぇ", Hebon: "we"},

			{Hiragana: "くぁ", Hebon: "kwa"},
			{Hiragana: "くぃ", Hebon: "kwi"},
			{Hiragana: "くぅ", Hebon: "kwu"},
			{Hiragana: "くぇ", Hebon: "kwe"},
			{Hiragana: "くぉ", Hebon: "kwo"},
			{Hiragana: "ぐぁ", Hebon: "gwa"},
			{Hiragana: "ぐぃ", Hebon: "gwi"},
			{Hiragana: "ぐぅ", Hebon: "gwu"},
			{Hiragana: "ぐぇ", Hebon: "gwe"},
			{Hiragana: "ぐぉ", Hebon: "gwo"},

			{Hiragana: "とぁ", Hebon: "twa"},
			{Hiragana: "とぃ", Hebon: "twi"},
			{Hiragana: "とぅ", Hebon: "twu"},
			{Hiragana: "とぇ", Hebon: "twe"},
			{Hiragana: "とぉ", Hebon: "two"},
			{Hiragana: "どぁ", Hebon: "dwa"},
			{Hiragana: "どぃ", Hebon: "dwi"},
			{Hiragana: "どぅ", Hebon: "dwu"},
			{Hiragana: "どぇ", Hebon: "dwe"},
			{Hiragana: "どぉ", Hebon: "dwo"},

			{Hiragana: "ふぁ", Hebon: "fa"},
			{Hiragana: "ふぃ", Hebon: "fi"},
			{Hiragana: "ふぇ", Hebon: "fe"},
			{Hiragana: "ふぉ", Hebon: "fo"},

			{Hiragana: "っか", Hebon: "kka"},
			{Hiragana: "っき", Hebon: "kki"},
			{Hiragana: "っく", Hebon: "kku"},
			{Hiragana: "っけ", Hebon: "kke"},
			{Hiragana: "っこ", Hebon: "kko"},
			{Hiragana: "っが", Hebon: "gga"},
			{Hiragana: "っぎ", Hebon: "ggi"},
			{Hiragana: "っぐ", Hebon: "ggu"},
			{Hiragana: "っげ", Hebon: "gge"},
			{Hiragana: "っご", Hebon: "ggo"},

			{Hiragana: "っさ", Hebon: "ssa"},
			{Hiragana: "っし", Hebon: "sshi"},
			{Hiragana: "っす", Hebon: "ssu"},
			{Hiragana: "っせ", Hebon: "sse"},
			{Hiragana: "っそ", Hebon: "sso"},
			{Hiragana: "っざ", Hebon: "zza"},
			{Hiragana: "っじ", Hebon: "jji"},
			{Hiragana: "っず", Hebon: "zzu"},
			{Hiragana: "っぜ", Hebon: "zze"},
			{Hiragana: "っぞ", Hebon: "zzo"},

			{Hiragana: "った", Hebon: "tta"},
			{Hiragana: "っち", Hebon: "tchi"},
			{Hiragana: "っつ", Hebon: "ttsu"},
			{Hiragana: "って", Hebon: "tte"},
			{Hiragana: "っと", Hebon: "tto"},
			{Hiragana: "っだ", Hebon: "dda"},
			{Hiragana: "っぢ", Hebon: "ddi"},
			{Hiragana: "っづ", Hebon: "ddu"},
			{Hiragana: "っで", Hebon: "dde"},
			{Hiragana: "っど", Hebon: "ddo"},

			{Hiragana: "っは", Hebon: "hha"},
			{Hiragana: "っひ", Hebon: "hhi"},
			{Hiragana: "っふ", Hebon: "ffu"},
			{Hiragana: "っへ", Hebon: "hhe"},
			{Hiragana: "っほ", Hebon: "hho"},
			{Hiragana: "っば", Hebon: "bba"},
			{Hiragana: "っび", Hebon: "bbi"},
			{Hiragana: "っぶ", Hebon: "bbu"},
			{Hiragana: "っべ", Hebon: "bbe"},
			{Hiragana: "っぼ", Hebon: "bbo"},
			{Hiragana: "っぱ", Hebon: "ppa"},
			{Hiragana: "っぴ", Hebon: "ppi"},
			{Hiragana: "っぷ", Hebon: "ppu"},
			{Hiragana: "っぺ", Hebon: "ppe"},
			{Hiragana: "っぽ", Hebon: "ppo"},

			{Hiragana: "っま", Hebon: "mma"},
			{Hiragana: "っみ", Hebon: "mmi"},
			{Hiragana: "っむ", Hebon: "mmu"},
			{Hiragana: "っめ", Hebon: "mme"},
			{Hiragana: "っも", Hebon: "mmo"},

			{Hiragana: "っや", Hebon: "yya"},
			{Hiragana: "っゆ", Hebon: "yyu"},
			{Hiragana: "っよ", Hebon: "yyo"},

			{Hiragana: "っら", Hebon: "rra"},
			{Hiragana: "っり", Hebon: "rri"},
			{Hiragana: "っる", Hebon: "rru"},
			{Hiragana: "っれ", Hebon: "rre"},
			{Hiragana: "っろ", Hebon: "rro"},

			{Hiragana: "っわ", Hebon: "wwa"},
			{Hiragana: "っゐ", Hebon: "wwi"},
			{Hiragana: "っゑ", Hebon: "wwe"},
			{Hiragana: "っを", Hebon: "wwo"},
		},
		OneCharactor: []HebonPair{
			{Hiragana: "あ", Hebon: "a"},
			{Hiragana: "い", Hebon: "i"},
			{Hiragana: "う", Hebon: "u"},
			{Hiragana: "え", Hebon: "e"},
			{Hiragana: "お", Hebon: "o"},
			{Hiragana: "ぁ", Hebon: "xa"},
			{Hiragana: "ぃ", Hebon: "xi"},
			{Hiragana: "ぅ", Hebon: "xu"},
			{Hiragana: "ぇ", Hebon: "xe"},
			{Hiragana: "ぉ", Hebon: "xo"},

			{Hiragana: "か", Hebon: "ka"},
			{Hiragana: "き", Hebon: "ki"},
			{Hiragana: "く", Hebon: "ku"},
			{Hiragana: "け", Hebon: "ke"},
			{Hiragana: "こ", Hebon: "ko"},
			{Hiragana: "が", Hebon: "ga"},
			{Hiragana: "ぎ", Hebon: "gi"},
			{Hiragana: "ぐ", Hebon: "gu"},
			{Hiragana: "げ", Hebon: "ge"},
			{Hiragana: "ご", Hebon: "go"},

			{Hiragana: "さ", Hebon: "sa"},
			{Hiragana: "し", Hebon: "shi"},
			{Hiragana: "す", Hebon: "su"},
			{Hiragana: "せ", Hebon: "se"},
			{Hiragana: "そ", Hebon: "so"},
			{Hiragana: "ざ", Hebon: "za"},
			{Hiragana: "じ", Hebon: "ji"},
			{Hiragana: "ず", Hebon: "zu"},
			{Hiragana: "ぜ", Hebon: "ze"},
			{Hiragana: "ぞ", Hebon: "zo"},

			{Hiragana: "た", Hebon: "ta"},
			{Hiragana: "ち", Hebon: "chi"},
			{Hiragana: "つ", Hebon: "tsu"},
			{Hiragana: "て", Hebon: "te"},
			{Hiragana: "と", Hebon: "to"},
			{Hiragana: "っ", Hebon: "xtsu"},
			{Hiragana: "だ", Hebon: "da"},
			{Hiragana: "ぢ", Hebon: "di"},
			{Hiragana: "づ", Hebon: "du"},
			{Hiragana: "で", Hebon: "de"},
			{Hiragana: "ど", Hebon: "do"},

			{Hiragana: "な", Hebon: "na"},
			{Hiragana: "に", Hebon: "ni"},
			{Hiragana: "ぬ", Hebon: "nu"},
			{Hiragana: "ね", Hebon: "ne"},
			{Hiragana: "の", Hebon: "no"},

			{Hiragana: "は", Hebon: "ha"},
			{Hiragana: "ひ", Hebon: "hi"},
			{Hiragana: "ふ", Hebon: "fu"},
			{Hiragana: "へ", Hebon: "he"},
			{Hiragana: "ほ", Hebon: "ho"},
			{Hiragana: "ば", Hebon: "ba"},
			{Hiragana: "び", Hebon: "bi"},
			{Hiragana: "ぶ", Hebon: "bu"},
			{Hiragana: "べ", Hebon: "be"},
			{Hiragana: "ぼ", Hebon: "bo"},
			{Hiragana: "ぱ", Hebon: "pa"},
			{Hiragana: "ぴ", Hebon: "pi"},
			{Hiragana: "ぷ", Hebon: "pu"},
			{Hiragana: "ぺ", Hebon: "pe"},
			{Hiragana: "ぽ", Hebon: "po"},

			{Hiragana: "ま", Hebon: "ma"},
			{Hiragana: "み", Hebon: "mi"},
			{Hiragana: "む", Hebon: "mu"},
			{Hiragana: "め", Hebon: "me"},
			{Hiragana: "も", Hebon: "mo"},

			{Hiragana: "や", Hebon: "ya"},
			{Hiragana: "ゆ", Hebon: "yu"},
			{Hiragana: "よ", Hebon: "yo"},
			{Hiragana: "ゃ", Hebon: "xya"},
			{Hiragana: "ゅ", Hebon: "xyu"},
			{Hiragana: "ょ", Hebon: "xyo"},

			{Hiragana: "ら", Hebon: "ra"},
			{Hiragana: "り", Hebon: "ri"},
			{Hiragana: "る", Hebon: "ru"},
			{Hiragana: "れ", Hebon: "re"},
			{Hiragana: "ろ", Hebon: "ro"},

			{Hiragana: "わ", Hebon: "wa"},
			{Hiragana: "ゐ", Hebon: "wi"},
			{Hiragana: "ゑ", Hebon: "we"},
			{Hiragana: "を", Hebon: "wo"},
			{Hiragana: "ゎ", Hebon: "xwa"},

			{Hiragana: "ん", Hebon: "nn"},

			{Hiragana: "ー", Hebon: "-"},
			{Hiragana: "、", Hebon: ","},
			{Hiragana: "。", Hebon: "."},
		},
	}
	return &h
}

func HiraganaToHebon(keyword string) string {
	hebonTable := NewHebonTable()
	_ = hebonTable
	//TODO: 文字数多い順に
	for _, v := range hebonTable.ThreeCharactors {
		keyword = strings.ReplaceAll(keyword, v.Hiragana, v.Hebon)
	}

	for _, v := range hebonTable.TwoCharactors {
		keyword = strings.ReplaceAll(keyword, v.Hiragana, v.Hebon)
	}

	for _, v := range hebonTable.OneCharactor {
		keyword = strings.ReplaceAll(keyword, v.Hiragana, v.Hebon)
	}

	return keyword
}
