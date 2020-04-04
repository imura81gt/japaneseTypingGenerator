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

```
echo "靴　黄色" | go run main.go
```

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


Not Yet
----------------------------------------

- [ ] Convert When input "助詞"
- [ ] Select blank type (narrow or widen)
- [ ] Use user dictionary
