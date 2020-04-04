Japanese Typing Generator
====================================

Overview
------------------------------------

If you want to create test data that depend on user input and Japanese, You can use this.


Usage
------------------------------------

### e.g. "日本語"

Run command

```
echo "日本語" | go run main.go
```

Result

```
に
にｈ
にほ
にほｎ
にほん
にほんｇ
にほんご
日本語
```


### e.g. "靴　黄色"

Run command

```
echo "靴　黄色" | go run main.go
```

Result

```
ｋ
く
くｔ
くｔｓ
くつ
靴
靴　ｋ
靴　き
靴　きい
靴　きいｒ
靴　きいろ
靴　黄色
```

### e.g. "Tシャツ　赤色　半袖"

```
echo "Tシャツ　赤色　半袖" | go run main.go
```

```
ｔ
ｔｈ
てぃ
てぃー
てぃーｓ
てぃーｓｈ
てぃーしゃ
てぃーしゃｔ
てぃーしゃｔｓ
てぃーしゃつ
Tシャツ
Tシャツ　あ
Tシャツ　あｋ
Tシャツ　あか
Tシャツ　あかい
Tシャツ　あかいｒ
Tシャツ　あかいろ
Tシャツ　赤色
Tシャツ　赤色　ｈ
Tシャツ　赤色　は
Tシャツ　赤色　はｎ
Tシャツ　赤色　はん
Tシャツ　赤色　はんｓ
Tシャツ　赤色　はんそ
Tシャツ　赤色　はんそｄ
Tシャツ　赤色　はんそで
Tシャツ　赤色　半袖
```

Not Yet
----------------------------------------

- [ ] Convert When input "助詞"
- [ ] Select blank type (narrow or widen)
- [ ] Use user dictionary
