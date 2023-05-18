+++
title = "unicode"
linkTitle = "unicode"
date = 2023-05-17T09:59:21+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# unicode

[https://pkg.go.dev/unicode@go1.20.1](https://pkg.go.dev/unicode@go1.20.1)

​	unicode包提供了测试Unicode码点某些属性的数据和函数。

### Example

​	以 "Is" 开头的函数可以用于检查一个 rune 属于哪个范围的表。请注意，一个 rune 可能适用于多个范围。

``` go 
package main

import (
	"fmt"
	"unicode"
)

func main() {

	// constant with mixed type runes
	const mixed = "\b5Ὂg̀9! ℃ᾭG"
	for _, c := range mixed {
		fmt.Printf("For %q:\n", c)
		if unicode.IsControl(c) {
			fmt.Println("\tis control rune")
		}
		if unicode.IsDigit(c) {
			fmt.Println("\tis digit rune")
		}
		if unicode.IsGraphic(c) {
			fmt.Println("\tis graphic rune")
		}
		if unicode.IsLetter(c) {
			fmt.Println("\tis letter rune")
		}
		if unicode.IsLower(c) {
			fmt.Println("\tis lower case rune")
		}
		if unicode.IsMark(c) {
			fmt.Println("\tis mark rune")
		}
		if unicode.IsNumber(c) {
			fmt.Println("\tis number rune")
		}
		if unicode.IsPrint(c) {
			fmt.Println("\tis printable rune")
		}
		if !unicode.IsPrint(c) {
			fmt.Println("\tis not printable rune")
		}
		if unicode.IsPunct(c) {
			fmt.Println("\tis punct rune")
		}
		if unicode.IsSpace(c) {
			fmt.Println("\tis space rune")
		}
		if unicode.IsSymbol(c) {
			fmt.Println("\tis symbol rune")
		}
		if unicode.IsTitle(c) {
			fmt.Println("\tis title case rune")
		}
		if unicode.IsUpper(c) {
			fmt.Println("\tis upper case rune")
		}
	}

}
Output:

For '\b':
	is control rune
	is not printable rune
For '5':
	is digit rune
	is graphic rune
	is number rune
	is printable rune
For 'Ὂ':
	is graphic rune
	is letter rune
	is printable rune
	is upper case rune
For 'g':
	is graphic rune
	is letter rune
	is lower case rune
	is printable rune
For '̀':
	is graphic rune
	is mark rune
	is printable rune
For '9':
	is digit rune
	is graphic rune
	is number rune
	is printable rune
For '!':
	is graphic rune
	is printable rune
	is punct rune
For ' ':
	is graphic rune
	is printable rune
	is space rune
For '℃':
	is graphic rune
	is printable rune
	is symbol rune
For 'ᾭ':
	is graphic rune
	is letter rune
	is printable rune
	is title case rune
For 'G':
	is graphic rune
	is letter rune
	is printable rune
	is upper case rune
```


## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=9)

``` go 
const (
	MaxRune         = '\U0010FFFF' // Maximum valid Unicode code point. 最大有效Unicode码点。
	ReplacementChar = '\uFFFD'     // Represents invalid code points.表示无效码点。
	MaxASCII        = '\u007F'     // maximum ASCII value. 最大ASCII值。
	MaxLatin1       = '\u00FF'     // maximum Latin-1 value.最大Latin-1值。
)
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=70)

``` go 
const (
	UpperCase = iota
	LowerCase
	TitleCase
	MaxCase
)
```

​	Delta数组的索引，表示用于大小写映射的不同情况。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=82)

``` go 
const (
	UpperLower = MaxRune + 1 // (不能是有效的增量。)
)
```

​	如果CaseRange结构体的Delta字段是UpperLower，则表示此CaseRange结构体表示形式为(例如)Upper Lower Upper Lower的序列。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/tables.go;l=6)

``` go 
const Version = "13.0.0"
```

​	Version是派生表的Unicode版本。

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/tables.go;l=3604)

``` go 
var (
	Cc     = _Cc // Cc是类别为Cc(其他，控制)的Unicode字符集。
	Cf     = _Cf // Cf是类别为Cf(其他，格式)的Unicode字符集。
	Co     = _Co // Co是类别为Co(其他，专用)的Unicode字符集。
	Cs     = _Cs // Cs是类别为Cs(其他，代理)的Unicode字符集。
	Digit  = _Nd // Digit是具有"十进制数"属性的Unicode字符集。
	Nd     = _Nd // Nd是类别为Nd(数字，十进制数)的Unicode字符集。
	Letter = _L  // Letter/L是Unicode字母集，类别L。
	L      = _L
	Lm     = _Lm // Lm是类别为Lm(字母，修饰符)的Unicode字符集。
	Lo     = _Lo // Lo是类别为Lo(字母，其他)的Unicode字符集。
	Lower  = _Ll // Lower是Unicode小写字母集。
	Ll     = _Ll // Ll是类别为Ll(字母，小写)的Unicode字符集。
	Mark   = _M  // Mark/M是Unicode标记字符集，类别M。
	M      = _M
	Mc     = _Mc // Mc是类别为Mc(标记，间距组合)的Unicode字符集。
	Me     = _Me // Me是类别为Me(标记，封闭)的Unicode字符集。
	Mn     = _Mn // Mn是类别为Mn(标记，非间距)的Unicode字符集。
	Nl     = _Nl // Nl是类别为Nl(数字，字母)的Unicode字符集。
	No     = _No // No是类别为No(数字，其他)的Unicode字符集。
	Number = _N  // Number/N是Unicode数字字符集，类别N。
	N      = _N
	Other  = _C //  Other/C是Unicode控制和特殊字符集，类别C。
	C      = _C
	Pc     = _Pc // Pc是类别为Pc(标点符号，连接符)的Unicode字符集。
	Pd     = _Pd // Pd是类别为Pd(标点符号，破折号)的Unicode字符集。
	Pe     = _Pe // Pe是类别为Pe(标点符号，右括号)的Unicode字符集。
	Pf     = _Pf // Pf是类别为Pf(标点符号，结束引号)的Unicode字符集。
	Pi     = _Pi // Pi是类别为Pi(标点符号，开始引号)的Unicode字符集。
	Po     = _Po // Po是类别为Po(标点符号，其他)的Unicode字符集。
	Ps     = _Ps // Ps是类别为Ps(标点符号，左括号)的Unicode字符集。
	Punct  = _P  // Punct/P是Unicode标点字符集，类别P。
	P      = _P
	Sc     = _Sc // Sc是类别为Sc(符号，货币)的Unicode字符集。
	Sk     = _Sk // Sk 是Unicode修饰符号字符集(Symbol, modifier)。
	Sm     = _Sm // Sm 是Unicode数学符号字符集(Symbol, math)。
	So     = _So // So 是Unicode其他符号字符集(Symbol, other)。
	Space  = _Z  // Space/Z 是Unicode空格字符集(Separator, space)。
	Z      = _Z
	Symbol = _S // Symbol/S 是Unicode符号字符集(Symbol)。
	S      = _S
	Title  = _Lt // Title 是Unicode标题大小写字母集。
	Lt     = _Lt // Lt 是Unicode标题大小写字母字符集(Letter, titlecase)。
	Upper  = _Lu // Upper 是Unicode大写字母集。
	Lu     = _Lu // Lu 是Unicode大写字母字符集(Letter, uppercase)。
	Zl     = _Zl // Zl 是Unicode行分隔符字符集(Separator, line)。
	Zp     = _Zp // Zp 是Unicode段落分隔符字符集(Separator, paragraph)。
	Zs     = _Zs // Zs 是Unicode空格分隔符字符集(Separator, space)。
)
```

这些变量的类型为`*RangeTable`。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/tables.go;l=5585)

``` go 
var (
	Adlam                  = _Adlam                  // Adlam 是脚本 Adlam 中的 Unicode 字符集。
	Ahom                   = _Ahom                   // Ahom 是脚本 Ahom 中的 Unicode 字符集。
	Anatolian_Hieroglyphs  = _Anatolian_Hieroglyphs  // Anatolian_Hieroglyphs 是脚本 Anatolian_Hieroglyphs 中的 Unicode 字符集。
	Arabic                 = _Arabic                 // Arabic 是脚本 Arabic 中的 Unicode 字符集。
	Armenian               = _Armenian               // Armenian 是脚本 Armenian 中的 Unicode 字符集。
	Avestan                = _Avestan                // Avestan 是脚本 Avestan 中的 Unicode 字符集。
	Balinese               = _Balinese               // Balinese 是脚本 Balinese 中的 Unicode 字符集。
	Bamum                  = _Bamum                  // Bamum 是脚本 Bamum 中的 Unicode 字符集。
	Bassa_Vah              = _Bassa_Vah              // Bassa_Vah 是脚本 Bassa_Vah 中的 Unicode 字符集。
	Batak                  = _Batak                  // Batak 是脚本 Batak 中的 Unicode 字符集。
	Bengali                = _Bengali                // Bengali 是脚本 Bengali 中的 Unicode 字符集。
	Bhaiksuki              = _Bhaiksuki              // Bhaiksuki 是脚本 Bhaiksuki 中的 Unicode 字符集。
	Bopomofo               = _Bopomofo               // Bopomofo 是脚本 Bopomofo 中的 Unicode 字符集。
	Brahmi                 = _Brahmi                 // Brahmi 是脚本 Brahmi 中的 Unicode 字符集。
	Braille                = _Braille                // Braille 是脚本 Braille 中的 Unicode 字符集。
	Buginese               = _Buginese               // Buginese 是脚本 Buginese 中的 Unicode 字符集。
	Buhid                  = _Buhid                  // Buhid 是脚本 Buhid 中的 Unicode 字符集。
	Canadian_Aboriginal    = _Canadian_Aboriginal    // Canadian_Aboriginal 是脚本Canadian_Aboriginal 中的 Unicode 字符集。
	Carian                 = _Carian                 // Carian 是脚本 Carian 中的 Unicode 字符集。
	Caucasian_Albanian     = _Caucasian_Albanian     // Caucasian_Albanian 是脚本 Caucasian_Albanian 中的 Unicode 字符集。
	Chakma                 = _Chakma                 // Chakma 是脚本 Chakma 中的 Unicode 字符集。
	Cham                   = _Cham                   // Cham 是脚本 Cham 中的 Unicode 字符集。
	Cherokee               = _Cherokee               // Cherokee 是脚本 Cherokee 中的 Unicode 字符集。
	Chorasmian             = _Chorasmian             // Chorasmian 是脚本 Chorasmian 中的 Unicode 字符集。
	Common                 = _Common                 // Common 是脚本 Common 中的 Unicode 字符集。
	Coptic                 = _Coptic                 // Coptic 是脚本 Coptic 中的 Unicode 字符集。
	Cuneiform              = _Cuneiform              // Cuneiform 是 Unicode 中的楔形文字脚本字符集。
	Cypriot                = _Cypriot                // Cypriot 是 Unicode 中的塞浦路斯音节脚本字符集。
	Cyrillic               = _Cyrillic               // Cyrillic 是 Unicode 中的西里尔字母脚本字符集。
	Deseret                = _Deseret                // Deseret 是 Unicode 中的德塞雷特字母脚本字符集。
	Devanagari             = _Devanagari             // Devanagari 是 Unicode 中的天城文脚本字符集。
	Dives_Akuru            = _Dives_Akuru            // Dives_Akuru 是 Unicode 中的迪维希阿库鲁文脚本字符集。
	Dogra                  = _Dogra                  // Dogra 是 Unicode 中的多格拉文脚本字符集。
	Duployan               = _Duployan               // Duployan 是 Unicode 中的杜普洛扬字母脚本字符集。
	Egyptian_Hieroglyphs   = _Egyptian_Hieroglyphs   // Egyptian_Hieroglyphs 是 Unicode 中的古埃及象形文字脚本字符集。
	Elbasan                = _Elbasan                // Elbasan 是 Unicode 中的埃尔巴桑字母脚本字符集。
	Elymaic                = _Elymaic                // Elymaic 是 Unicode 中的埃勒美亚文脚本字符集。
	Ethiopic               = _Ethiopic               // Ethiopic 是 Unicode 中的吉兹文脚本字符集。
	Georgian               = _Georgian               // Georgian 是 Unicode 中的格鲁吉亚字母脚本字符集。
	Glagolitic             = _Glagolitic             // Glagolitic 是 Unicode 中的格拉哥里字母脚本字符集。
	Gothic                 = _Gothic                 // Gothic 是 Unicode 字符集中哥特文脚本的字符集。
	Grantha                = _Grantha                // Grantha 是 Unicode 字符集中格兰塔文脚本的字符集。
	Greek                  = _Greek                  // Greek 是 Unicode 字符集中希腊文脚本的字符集。
	Gujarati               = _Gujarati               // Gujarati 是 Unicode 字符集中古吉拉特文脚本的字符集。
	Gunjala_Gondi          = _Gunjala_Gondi          // Gunjala_Gondi 是 Unicode 字符集中贡杰拉贡迪文脚本的字符集。
	Gurmukhi               = _Gurmukhi               // Gurmukhi 是 Unicode 字符集中古木基文脚本的字符集。
	Han                    = _Han                    // Han 是 Unicode 字符集中汉字的字符集。
	Hangul                 = _Hangul                 // Hangul 是 Unicode 字符集中朝鲜文脚本的字符集。
	Hanifi_Rohingya        = _Hanifi_Rohingya        // Hanifi_Rohingya 是 Unicode 字符集中罕尼孚·罗兴亚文脚本的字符集。
	Hanunoo                = _Hanunoo                // Hanunoo 是 Unicode 字符集中哈努努文脚本的字符集。
	Hatran                 = _Hatran                 // Hatran 是 Unicode 字符集中哈特兰文脚本的字符集。
	Hebrew                 = _Hebrew                 // Hebrew 是 Unicode 字符集中希伯来文脚本的字符集。
	Hiragana               = _Hiragana               // Hiragana 是 Unicode 字符集中日文平假名的字符集。
	Imperial_Aramaic       = _Imperial_Aramaic       // Imperial_Aramaic 是 Unicode 字符集中帝国阿拉米文脚本的字符集。
	Inherited              = _Inherited              // Inherited 是 Unicode 字符集中继承类的字符集。
	Inscriptional_Pahlavi  = _Inscriptional_Pahlavi  // Inscriptional_Pahlavi 是 Unicode 字符集中巴列维文碑铭脚本的字符集。
	Inscriptional_Parthian = _Inscriptional_Parthian // Inscriptional_Parthian 是 Unicode 字符集中帕提亚文碑铭脚本的字符集。
	Javanese               = _Javanese               // Javanese 是 Unicode 字符集中爪哇文脚本的字符集。
	Kaithi                 = _Kaithi                 // Kaithi 是 Unicode 字符集中凯提文脚本的字符集。
	Kannada                = _Kannada                // Kannada 是 Unicode 字符集中卡纳达文脚本的字符集。
	Katakana               = _Katakana               // Katakana 是 Unicode 字符集中日文片假名的字符集。
	Kayah_Li               = _Kayah_Li               // Kayah_Li 是蒲瑤文的 Unicode 字符集。
	Kharoshthi             = _Kharoshthi             // Kharoshthi 是佉盧文的 Unicode 字符集。
	Khitan_Small_Script    = _Khitan_Small_Script    // Khitan_Small_Script 是契丹小字的 Unicode 字符集。
	Khmer                  = _Khmer                  // Khmer 是高棉文的 Unicode 字符集。
	Khojki                 = _Khojki                 // Khojki 是库杰克文的 Unicode 字符集。
	Khudawadi              = _Khudawadi              // Khudawadi 是库达瓦迪文的 Unicode 字符集。
	Lao                    = _Lao                    // Lao 是老挝文的 Unicode 字符集。
	Latin                  = _Latin                  // Latin 是拉丁文的 Unicode 字符集。
	Lepcha                 = _Lepcha                 // Lepcha 是雷布查文的 Unicode 字符集。
	Limbu                  = _Limbu                  // Limbu 是林布文的 Unicode 字符集。
	Linear_A               = _Linear_A               // Linear_A 是线性 A 文字的 Unicode 字符集。
	Linear_B               = _Linear_B               // Linear_B 是线性 B 文字的 Unicode 字符集。
	Lisu                   = _Lisu                   // Lisu 是傈僳文的 Unicode 字符集。
	Lycian                 = _Lycian                 // Lycian是Unicode字符集中Lycian脚本的字符集。
	Lydian                 = _Lydian                 // Lydian是Unicode字符集中Lydian脚本的字符集。
	Mahajani               = _Mahajani               // Mahajani是Unicode字符集中Mahajani脚本的字符集。
	Makasar                = _Makasar                // Makasar是Unicode字符集中Makasar脚本的字符集。
	Malayalam              = _Malayalam              // Malayalam是Unicode字符集中Malayalam脚本的字符集。
	Mandaic                = _Mandaic                // Mandaic是Unicode字符集中Mandaic脚本的字符集。
	Manichaean             = _Manichaean             // Manichaean是Unicode字符集中Manichaean脚本的字符集。
	Marchen                = _Marchen                // Marchen是Unicode字符集中Marchen脚本的字符集。
	Masaram_Gondi          = _Masaram_Gondi          // Masaram_Gondi是Unicode字符集中Masaram_Gondi脚本的字符集。
	Medefaidrin            = _Medefaidrin            // Medefaidrin是Unicode字符集中Medefaidrin脚本的字符集。
	Meetei_Mayek           = _Meetei_Mayek           // Meetei_Mayek是Unicode字符集中Meetei_Mayek脚本的字符集。
	Mende_Kikakui          = _Mende_Kikakui          // Mende_Kikakui是Unicode字符集中Mende_Kikakui脚本的字符集。
	Meroitic_Cursive       = _Meroitic_Cursive       // Meroitic_Cursive是Unicode字符集中Meroitic_Cursive脚本的字符集。
	Meroitic_Hieroglyphs   = _Meroitic_Hieroglyphs   // Meroitic_Hieroglyphs是Unicode字符集中Meroitic_Hieroglyphs脚本的字符集。
	Miao                   = _Miao                   // Miao是Unicode字符集中Miao脚本的字符集。
	Modi                   = _Modi                   // Modi是Unicode字符集中Modi脚本的字符集。
	Mongolian              = _Mongolian              // Mongolian是Unicode字符集中Mongolian脚本的字符集。
	Mro                    = _Mro                    // Mro是莫罗文的Unicode字符集。
	Multani                = _Multani                // Multani是Multani文的Unicode字符集。
	Myanmar                = _Myanmar                // 缅甸语是缅甸语文字符集。
	Nabataean              = _Nabataean              // 纳巴泰文字是纳巴泰文字的Unicode字符集。
	Nandinagari            = _Nandinagari            // Nandinagari是南迪那加里文的Unicode字符集。
	New_Tai_Lue            = _New_Tai_Lue            // 新傣文是新傣文的Unicode字符集。
	Newa                   = _Newa                   // Newa是尼瓦尔文的Unicode字符集。
	Nko                    = _Nko                    // N'Ko是N'Ko文的Unicode字符集。
	Nushu                  = _Nushu                  // Nushu是女书的Unicode字符集。
	Nyiakeng_Puachue_Hmong = _Nyiakeng_Puachue_Hmong // Nyiakeng_Puachue_Hmong是Nyiakeeng Puachue Hmong文的Unicode字符集。
	Ogham                  = _Ogham                  // Ogham是欧甘文的Unicode字符集。
	Ol_Chiki               = _Ol_Chiki               // Ol Chiki是Ol Chiki文的Unicode字符集。
	Old_Hungarian          = _Old_Hungarian          // Old Hungarian是古匈牙利文的Unicode字符集。
	Old_Italic             = _Old_Italic             // Old Italic是古意大利文的Unicode字符集。
	Old_North_Arabian      = _Old_North_Arabian      // Old North Arabian是古北阿拉伯文的Unicode字符集。
	Old_Permic             = _Old_Permic             // Old Permic是古彼尔姆文的Unicode字符集。
	Old_Persian            = _Old_Persian            // Old Persian是古波斯文的Unicode字符集。
	Old_Sogdian            = _Old_Sogdian            // Old Sogdian是古粟特文的Unicode字符集。
	Old_South_Arabian      = _Old_South_Arabian      // Old South Arabian是古南阿拉伯文的Unicode字符集。
	Old_Turkic             = _Old_Turkic             // Old Turkic是古突厥文的Unicode字符集。
	Oriya                  = _Oriya                  // Oriya是奥里亚文的Unicode字符集。
	Osage                  = _Osage                  // Osage是奥萨格文的Unicode字符集。
	Osmanya                = _Osmanya                // Osmanya是奥斯曼亚文的Unicode字符集。
	Pahawh_Hmong           = _Pahawh_Hmong           // Pahawh_Hmong是白苗文的Unicode字符集。
	Palmyrene              = _Palmyrene              // Palmyrene是巴尔米拉文的Unicode字符集。
	Pau_Cin_Hau            = _Pau_Cin_Hau            // Pau_Cin_Hau是保赤话的Unicode字符集。
	Phags_Pa               = _Phags_Pa               // Phags_Pa是八思巴文的Unicode字符集。
	Phoenician             = _Phoenician             // Phoenician是腓尼基文的Unicode字符集。
	Psalter_Pahlavi        = _Psalter_Pahlavi        // Psalter_Pahlavi是赞德文字母书的Unicode字符集。
	Rejang                 = _Rejang                 // Rejang是热洛温的Unicode字符集。
	Runic                  = _Runic                  // Runic是古代北欧文的Unicode字符集。
	Samaritan              = _Samaritan              // Samaritan是撒马利亚文的Unicode字符集。
	Saurashtra             = _Saurashtra             // Saurashtra是索拉什特拉文的Unicode字符集。
	Sharada                = _Sharada                // Sharada是沙拉达文的Unicode字符集。
	Shavian                = _Shavian                // Shavian是肖维安文的Unicode字符集。
	Siddham                = _Siddham                // Siddham是悉昙文的Unicode字符集。
	SignWriting            = _SignWriting            // SignWriting是符号书写系统的Unicode字符集。
	Sinhala                = _Sinhala                // Sinhala是僧伽罗文的Unicode字符集。
	Sogdian                = _Sogdian                // Sogdian是粟特文的Unicode字符集。
	Sora_Sompeng           = _Sora_Sompeng           // Sora_Sompeng是梭罗文的Unicode字符集。
	Soyombo                = _Soyombo                // Soyombo是索永布文的Unicode字符集。
	Sundanese              = _Sundanese              // Sundanese是巽他文的Unicode字符集。
	Syloti_Nagri           = _Syloti_Nagri           // Syloti_Nagri是锡罗蒂镇那格里文的Unicode字符集。
	Syriac                 = _Syriac                 // Syriac是叙利亚文的Unicode字符集。
	Tagalog                = _Tagalog                // Tagalog是塔加洛语的Unicode字符集。
	Tagbanwa               = _Tagbanwa               // Tagbanwa是塔格巴努亚文的Unicode字符集。
	Tai_Le                 = _Tai_Le                 // Tai_Le是傣历文的Unicode字符集。
	Tai_Tham               = _Tai_Tham               // Tai_Tham是傣南文的Unicode字符集。
	Tai_Viet               = _Tai_Viet               // Tai_Viet是傣越文的Unicode字符集。
	Takri                  = _Takri                  // Takri是Takri文字的Unicode字符集。
	Tamil                  = _Tamil                  // Tamil是Tamil文字的Unicode字符集。
	Tangut                 = _Tangut                 // Tangut是Tangut文字的Unicode字符集。
	Telugu                 = _Telugu                 // Telugu是Telugu文字的Unicode字符集。
	Thaana                 = _Thaana                 // Thaana是Thaana文字的Unicode字符集。
	Thai                   = _Thai                   // Thai是Thai文字的Unicode字符集。
	Tibetan                = _Tibetan                // Tibetan是藏文的Unicode字符集。
	Tifinagh               = _Tifinagh               // Tifinagh是Tifinagh文字的Unicode字符集。
	Tirhuta                = _Tirhuta                // Tirhuta是Tirhuta文字的Unicode字符集。
	Ugaritic               = _Ugaritic               // Ugaritic是Ugaritic文字的Unicode字符集。
	Vai                    = _Vai                    // Vai是Vai文字的Unicode字符集。
	Wancho                 = _Wancho                 // Wancho是Wancho文字的Unicode字符集。
	Warang_Citi            = _Warang_Citi            // Warang_Citi是Warang Citi文字的Unicode字符集。
	Yezidi                 = _Yezidi                 // Yezidi是Yezidi文字的Unicode字符集。
	Yi                     = _Yi                     // Yi是彝文的Unicode字符集。
	Zanabazar_Square       = _Zanabazar_Square       // Zanabazar_Square是Zanabazar方块文字的Unicode字符集。
)
```

​	这些变量的类型为`*RangeTable`。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/tables.go;l=6889)

``` go 
var (
	ASCII_Hex_Digit                    = _ASCII_Hex_Digit                    // ASCII_Hex_Digit 是具有 ASCII_Hex_Digit 属性的 Unicode 字符集。
	Bidi_Control                       = _Bidi_Control                       // Bidi_Control 是具有 Bidi_Control 属性的 Unicode 字符集。
	Dash                               = _Dash                               // Dash 是具有 Dash 属性的 Unicode 字符集。
	Deprecated                         = _Deprecated                         // Deprecated 是具有 Deprecated 属性的 Unicode 字符集。
	Diacritic                          = _Diacritic                          // Diacritic 是具有 Diacritic 属性的 Unicode 字符集。
	Extender                           = _Extender                           // Extender 是具有 Extender 属性的 Unicode 字符集。
	Hex_Digit                          = _Hex_Digit                          // Hex_Digit 是具有 Hex_Digit 属性的 Unicode 字符集。
	Hyphen                             = _Hyphen                             // Hyphen 是具有 Hyphen 属性的 Unicode 字符集。
	IDS_Binary_Operator                = _IDS_Binary_Operator                // IDS_Binary_Operator 是具有 IDS_Binary_Operator 属性的 Unicode 字符集。
	IDS_Trinary_Operator               = _IDS_Trinary_Operator               // IDS_Trinary_Operator 是具有 IDS_Trinary_Operator 属性的 Unicode 字符集。
	Ideographic                        = _Ideographic                        // Ideographic 是具有 Ideographic 属性的 Unicode 字符集。
	Join_Control                       = _Join_Control                       // Join_Control 是具有 Join_Control 属性的 Unicode 字符集。
	Logical_Order_Exception            = _Logical_Order_Exception            // Logical_Order_Exception 是具有 Logical_Order_Exception 属性的 Unicode 字符集。
	Noncharacter_Code_Point            = _Noncharacter_Code_Point            // Noncharacter_Code_Point 是具有 Noncharacter_Code_Point 属性的 Unicode 字符集。
	Other_Alphabetic                   = _Other_Alphabetic                   // Other_Alphabetic 是具有 Other_Alphabetic 属性的 Unicode 字符集。
	Other_Default_Ignorable_Code_Point = _Other_Default_Ignorable_Code_Point // Other_Default_Ignorable_Code_Point 是具有 Other_Default_Ignorable_Code_Point 属性的 Unicode 字符集。
	Other_Grapheme_Extend              = _Other_Grapheme_Extend              // Other_Grapheme_Extend 是具有 Other_Grapheme_Extend 属性的 Unicode 字符集。
	Other_ID_Continue                  = _Other_ID_Continue                  // Other_ID_Continue 是具有 Other_ID_Continue 属性的 Unicode 字符集。
	Other_ID_Start                     = _Other_ID_Start                     // Other_ID_Start 是 Unicode 字符集中具有 Other_ID_Start 属性的字符集。
	Other_Lowercase                    = _Other_Lowercase                    // Other_Lowercase 是 Unicode 字符集中具有 Other_Lowercase 属性的字符集。
	Other_Math                         = _Other_Math                         // Other_Math 是 Unicode 字符集中具有 Other_Math 属性的字符集。
	Other_Uppercase                    = _Other_Uppercase                    // Other_Uppercase 是 Unicode 字符集中具有 Other_Uppercase 属性的字符集。
	Pattern_Syntax                     = _Pattern_Syntax                     // Pattern_Syntax 是 Unicode 字符集中具有 Pattern_Syntax 属性的字符集。
	Pattern_White_Space                = _Pattern_White_Space                // Pattern_White_Space 是 Unicode 字符集中具有 Pattern_White_Space 属性的字符集。
	Prepended_Concatenation_Mark       = _Prepended_Concatenation_Mark       // Prepended_Concatenation_Mark 是 Unicode 字符集中具有 Prepended_Concatenation_Mark 属性的字符集。
	Quotation_Mark                     = _Quotation_Mark                     // Quotation_Mark 是 Unicode 字符集中具有 Quotation_Mark 属性的字符集。
	Radical                            = _Radical                            // Radical 是 Unicode 字符集中具有 Radical 属性的字符集。
	Regional_Indicator                 = _Regional_Indicator                 // Regional_Indicator 是 Unicode 字符集中具有 Regional_Indicator 属性的字符集。
	STerm                              = _Sentence_Terminal                  // STerm 是 Sentence_Terminal 的别名。
	Sentence_Terminal                  = _Sentence_Terminal                  // Sentence_Terminal 是 Unicode 字符集中具有 Sentence_Terminal 属性的字符集。
	Soft_Dotted                        = _Soft_Dotted                        // Soft_Dotted 是 Unicode 字符集中具有 Soft_Dotted 属性的字符集。
	Terminal_Punctuation               = _Terminal_Punctuation               // Terminal_Punctuation 是 Unicode 字符集中具有 Terminal_Punctuation 属性的字符集。
	Unified_Ideograph                  = _Unified_Ideograph                  // Unified_Ideograph 是 Unicode 字符集中具有 Unified_Ideograph 属性的字符集。
	Variation_Selector                 = _Variation_Selector                 // Variation_Selector 是 Unicode 字符集中具有 Variation_Selector 属性的字符集。
	White_Space                        = _White_Space                        // White_Space 是 Unicode 字符集中具有 White_Space 属性的字符集。
)
```

​	这些变量的类型为`*RangeTable`。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/tables.go;l=6929)

``` go 
var CaseRanges = _CaseRanges
```

​	CaseRanges是描述所有具有非自身映射的字母的大小写映射的表。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/tables.go;l=9)

``` go 
var Categories = map[string]*RangeTable{
	"C":  C,
	"Cc": Cc,
	"Cf": Cf,
	"Co": Co,
	"Cs": Cs,
	"L":  L,
	"Ll": Ll,
	"Lm": Lm,
	"Lo": Lo,
	"Lt": Lt,
	"Lu": Lu,
	"M":  M,
	"Mc": Mc,
	"Me": Me,
	"Mn": Mn,
	"N":  N,
	"Nd": Nd,
	"Nl": Nl,
	"No": No,
	"P":  P,
	"Pc": Pc,
	"Pd": Pd,
	"Pe": Pe,
	"Pf": Pf,
	"Pi": Pi,
	"Po": Po,
	"Ps": Ps,
	"S":  S,
	"Sc": Sc,
	"Sk": Sk,
	"Sm": Sm,
	"So": So,
	"Z":  Z,
	"Zl": Zl,
	"Zp": Zp,
	"Zs": Zs,
}
```

​	Categories是Unicode类别表的集合。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/tables.go;l=7736)

``` go 
var FoldCategory = map[string]*RangeTable{
	"L":  foldL,
	"Ll": foldLl,
	"Lt": foldLt,
	"Lu": foldLu,
	"M":  foldM,
	"Mn": foldMn,
}
```

​	FoldCategory将类别名称映射到表格，其中类别外的码点在简单大小写折叠下与类别内的码点等效。如果类别名称没有条目，则不存在这样的点。 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/tables.go;l=8026)

``` go 
var FoldScript = map[string]*RangeTable{
	"Common":    foldCommon,
	"Greek":     foldGreek,
	"Inherited": foldInherited,
}
```

​	FoldScript将脚本名称映射到表格，其中脚本外的码点在简单大小写折叠下与脚本内的码点等效。如果脚本名称没有条目，则不存在这样的点。 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/graphic.go;l=23)

``` go 
var GraphicRanges = []*RangeTable{
	L, M, N, P, S, Zs,
}
```

​	GraphicRanges根据Unicode定义了一组图形字符。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/graphic.go;l=29)

``` go 
var PrintRanges = []*RangeTable{
	L, M, N, P, S,
}
```

​	PrintRanges根据Go定义了一组可打印字符。 ASCII空格U+0020单独处理。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/tables.go;l=5745)

``` go 
var Properties = map[string]*RangeTable{
	"ASCII_Hex_Digit":                    ASCII_Hex_Digit,
	"Bidi_Control":                       Bidi_Control,
	"Dash":                               Dash,
	"Deprecated":                         Deprecated,
	"Diacritic":                          Diacritic,
	"Extender":                           Extender,
	"Hex_Digit":                          Hex_Digit,
	"Hyphen":                             Hyphen,
	"IDS_Binary_Operator":                IDS_Binary_Operator,
	"IDS_Trinary_Operator":               IDS_Trinary_Operator,
	"Ideographic":                        Ideographic,
	"Join_Control":                       Join_Control,
	"Logical_Order_Exception":            Logical_Order_Exception,
	"Noncharacter_Code_Point":            Noncharacter_Code_Point,
	"Other_Alphabetic":                   Other_Alphabetic,
	"Other_Default_Ignorable_Code_Point": Other_Default_Ignorable_Code_Point,
	"Other_Grapheme_Extend":              Other_Grapheme_Extend,
	"Other_ID_Continue":                  Other_ID_Continue,
	"Other_ID_Start":                     Other_ID_Start,
	"Other_Lowercase":                    Other_Lowercase,
	"Other_Math":                         Other_Math,
	"Other_Uppercase":                    Other_Uppercase,
	"Pattern_Syntax":                     Pattern_Syntax,
	"Pattern_White_Space":                Pattern_White_Space,
	"Prepended_Concatenation_Mark":       Prepended_Concatenation_Mark,
	"Quotation_Mark":                     Quotation_Mark,
	"Radical":                            Radical,
	"Regional_Indicator":                 Regional_Indicator,
	"Sentence_Terminal":                  Sentence_Terminal,
	"STerm":                              Sentence_Terminal,
	"Soft_Dotted":                        Soft_Dotted,
	"Terminal_Punctuation":               Terminal_Punctuation,
	"Unified_Ideograph":                  Unified_Ideograph,
	"Variation_Selector":                 Variation_Selector,
	"White_Space":                        White_Space,
}
```

​	Properties是Unicode属性表的集合。 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/tables.go;l=3655)

``` go 
var Scripts = map[string]*RangeTable{}/* 156 elements not displayed */
```

​	Scripts是Unicode脚本表的集合。

## 函数

#### func [In](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/graphic.go;l=69)  <- go1.2

``` go 
func In(r rune, ranges ...*RangeTable) bool
```

​	In函数报告符文是否属于其中一个范围。

#### func [Is](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=157) 

``` go 
func Is(rangeTab *RangeTable, r rune) bool
```

​	Is函数报告符文是否在指定的范围表中。

#### func [IsControl](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/graphic.go;l=81) 

``` go 
func IsControl(r rune) bool
```

​	IsControl函数报告符文是否为控制字符。C (Other) Unicode类别包括更多的码点，例如代理项；使用Is(C, r) 进行测试。

#### func [IsDigit](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/digit.go;l=8) 

``` go 
func IsDigit(r rune) bool
```

​	IsDigit函数报告符文是否为十进制数字。

##### IsDigit Example
``` go 
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%t\n", unicode.IsDigit('৩'))
	fmt.Printf("%t\n", unicode.IsDigit('A'))
}
Output:

true
false
```

#### func [IsGraphic](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/graphic.go;l=36) 

``` go 
func IsGraphic(r rune) bool
```

​	IsGraphic函数报告符文是否根据Unicode定义为图形字符。这些字符包括类别L、M、N、P、S、Zs的字母、标记、数字、标点、符号和空格。

#### func [IsLetter](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/graphic.go;l=90) 

``` go 
func IsLetter(r rune) bool
```

​	IsLetter函数报告符文是否为字母(类别L)。

##### IsLetter Example
``` go 
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%t\n", unicode.IsLetter('A'))
	fmt.Printf("%t\n", unicode.IsLetter('7'))
}
Output:

true
false
```

#### func [IsLower](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=193) 

``` go 
func IsLower(r rune) bool
```

​	IsLower函数报告符文是否为小写字母。

##### IsLower Example
``` go 
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%t\n", unicode.IsLower('a'))
	fmt.Printf("%t\n", unicode.IsLower('A'))
}
Output:

true
false
```

#### func [IsMark](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/graphic.go;l=98) 

``` go 
func IsMark(r rune) bool
```

​	IsMark函数报告符文是否为标记字符(类别M)。

#### func [IsNumber](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/graphic.go;l=104) 

``` go 
func IsNumber(r rune) bool
```

​	IsNumber函数报告符文是否为数字(类别N)。

##### IsNumber Example
``` go 
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%t\n", unicode.IsNumber('Ⅷ'))
	fmt.Printf("%t\n", unicode.IsNumber('A'))
}
Output:

true
false
```

#### func [IsOneOf](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/graphic.go;l=59) 

``` go 
func IsOneOf(ranges []*RangeTable, r rune) bool
```

​	IsOneOf函数报告符文是否属于其中一个范围。函数"In"提供了更好的签名，应优先使用IsOneOf。

#### func [IsPrint](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/graphic.go;l=50) 

``` go 
func IsPrint(r rune) bool
```

​	IsPrint函数报告符文是否根据Go定义为可打印字符。这些字符包括类别L、M、N、P、S和ASCII空格字符。此分类与IsGraphic相同，除了唯一的间距字符是ASCII空格U+0020。

#### func [IsPunct](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/graphic.go;l=113) 

``` go 
func IsPunct(r rune) bool
```

​	IsPunct函数报告符文是否为Unicode标点字符(类别P)。

#### func [IsSpace](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/graphic.go;l=128) 

``` go 
func IsSpace(r rune) bool
```

​	IsSpace函数判断rune是否为Unicode空格字符，包括Latin-1空格字符

```
'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
```

以及其他由类别Z和属性Pattern_White_Space定义的空格字符。

##### IsSpace Example
``` go 
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%t\n", unicode.IsSpace(' '))
	fmt.Printf("%t\n", unicode.IsSpace('\n'))
	fmt.Printf("%t\n", unicode.IsSpace('\t'))
	fmt.Printf("%t\n", unicode.IsSpace('a'))
}
Output:

true
true
true
false
```

#### func [IsSymbol](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/graphic.go;l=141) 

``` go 
func IsSymbol(r rune) bool
```

​	IsSymbol函数判断rune是否为Unicode符号字符。

#### func [IsTitle](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=202) 

``` go 
func IsTitle(r rune) bool
```

​	IsTitle函数判断rune是否为Unicode标题大小写字母。

##### IsTitle Example
``` go 
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%t\n", unicode.IsTitle('ǅ'))
	fmt.Printf("%t\n", unicode.IsTitle('a'))
}
Output:

true
false
```

#### func [IsUpper](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=184) 

``` go 
func IsUpper(r rune) bool
```

​	IsUpper函数判断rune是否为Unicode大写字母。

##### IsUpper Example
``` go 
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%t\n", unicode.IsUpper('A'))
	fmt.Printf("%t\n", unicode.IsUpper('a'))
}
Output:

true
false
```

#### func [SimpleFold](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=340) 

``` go 
func SimpleFold(r rune) rune
```

​	SimpleFold函数迭代Unicode定义的简单大小写折叠等效的Unicode码点。在等效于rune(包括rune本身)的码点中，如果存在一个最小的大于r的rune，则SimpleFold函数返回该rune，否则返回最小的rune>=0。如果r不是有效的Unicode码点，则SimpleFold(r)返回r。

For example:

```
SimpleFold('A') = 'a'
SimpleFold('a') = 'A'

SimpleFold('K') = 'k'
SimpleFold('k') = '\u212A' (Kelvin symbol, K)
SimpleFold('\u212A') = 'K'

SimpleFold('1') = '1'

SimpleFold(-2) = -2
```

##### SimpleFold Example
``` go 
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%#U\n", unicode.SimpleFold('A'))      // 'a'
	fmt.Printf("%#U\n", unicode.SimpleFold('a'))      // 'A'
	fmt.Printf("%#U\n", unicode.SimpleFold('K'))      // 'k'
	fmt.Printf("%#U\n", unicode.SimpleFold('k'))      // '\u212A' (Kelvin symbol, K)
	fmt.Printf("%#U\n", unicode.SimpleFold('\u212A')) // 'K'
	fmt.Printf("%#U\n", unicode.SimpleFold('1'))      // '1'

}
Output:

U+0061 'a'
U+0041 'A'
U+006B 'k'
U+212A 'K'
U+004B 'K'
U+0031 '1'
```

#### func [To](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=248) 

``` go 
func To(_case int, r rune) rune
```

​	To函数将rune映射到指定的大小写：UpperCase，LowerCase或TitleCase。

##### To Example
``` go 
package main

import (
	"fmt"
	"unicode"
)

func main() {
	const lcG = 'g'
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, lcG))
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, lcG))
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, lcG))

	const ucG = 'G'
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, ucG))
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, ucG))
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, ucG))

}
Output:

U+0047 'G'
U+0067 'g'
U+0047 'G'
U+0047 'G'
U+0067 'g'
U+0047 'G'
```

#### func [ToLower](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=265) 

``` go 
func ToLower(r rune) rune
```

​	ToLower函数将rune映射为小写字母。

##### ToLower  Example
``` go 
package main

import (
	"fmt"
	"unicode"
)

func main() {
	const ucG = 'G'
	fmt.Printf("%#U\n", unicode.ToLower(ucG))

}
Output:

U+0067 'g'
```

#### func [ToTitle](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=276) 

``` go 
func ToTitle(r rune) rune
```

​	ToTitle函数将rune映射为标题大小写。

##### ToTitle Example
``` go 
package main

import (
	"fmt"
	"unicode"
)

func main() {
	const ucG = 'g'
	fmt.Printf("%#U\n", unicode.ToTitle(ucG))

}
Output:

U+0047 'G'
```

#### func [ToUpper](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=254) 

``` go 
func ToUpper(r rune) rune
```

​	ToUpper函数将rune映射为大写字母。

##### ToUpper Example
``` go 
package main

import (
	"fmt"
	"unicode"
)

func main() {
	const ucG = 'g'
	fmt.Printf("%#U\n", unicode.ToUpper(ucG))

}
Output:

U+0047 'G'
```

## 类型

### type [CaseRange](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=56) 

``` go 
type CaseRange struct {
	Lo    uint32
	Hi    uint32
	Delta d
}
```

​	CaseRange结构体表示Unicode码点的一个范围，用于简单(一个码点对一个码点)的大小写转换。范围从Lo到Hi，包括Lo和Hi，步幅为1。Delta是添加到码点的数字，以达到该字符的不同大小写的码点。它们可以是负数。如果是零，则表示该字符的大小写相同。有一种特殊情况，表示交替的对应的大写和小写字符序列。它显示为具有固定Delta的{UpperLower，UpperLower，UpperLower}。常量UpperLower具有无法实现的Delta值。

### type [Range16](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=29) 

``` go 
type Range16 struct {
	Lo     uint16
	Hi     uint16
	Stride uint16
}
```

​	Range16结构体表示16位Unicode码点的一个范围。范围从Lo到Hi，包括Lo和Hi，并具有指定的步幅。

### type [Range32](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=38) 

``` go 
type Range32 struct {
	Lo     uint32
	Hi     uint32
	Stride uint32
}
```

​	Range32结构体表示 Unicode 码点的范围，当一个或多个值无法适应 16 位时使用。该范围从 Lo 到 Hi，包括 Lo 和 Hi，并具有指定的步长。Lo 和 Hi 必须始终 >= 1<<16。

### type [RangeTable](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=21) 

``` go 
type RangeTable struct {
	R16         []Range16
	R32         []Range32
	LatinOffset int // number of entries in R16 with Hi <= MaxLatin1
}
```

​	RangeTable结构体通过列出集合内码点的范围来定义 Unicode 码点的集合。范围在两个切片中列出以节省空间：一个 16 位范围的切片和一个 32 位范围的切片。两个切片必须按排序顺序且不重叠。此外，R32 应仅包含值 >= 0x10000(1<<16)。

### type [SpecialCase](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=64) 

``` go 
type SpecialCase []CaseRange
```

​	SpecialCase表示特定语言的大小写映射，例如土耳其语。SpecialCase 的方法通过自定义(覆盖)标准映射来定制。

##### SpecialCase Example
``` go 
package main

import (
	"fmt"
	"unicode"
)

func main() {
	t := unicode.TurkishCase

	const lci = 'i'
	fmt.Printf("%#U\n", t.ToLower(lci))
	fmt.Printf("%#U\n", t.ToTitle(lci))
	fmt.Printf("%#U\n", t.ToUpper(lci))

	const uci = 'İ'
	fmt.Printf("%#U\n", t.ToLower(uci))
	fmt.Printf("%#U\n", t.ToTitle(uci))
	fmt.Printf("%#U\n", t.ToUpper(uci))

}
Output:

U+0069 'i'
U+0130 'İ'
U+0130 'İ'
U+0069 'i'
U+0130 'İ'
U+0130 'İ'
```

``` go 
var AzeriCase SpecialCase = _TurkishCase
var TurkishCase SpecialCase = _TurkishCase
```

#### (SpecialCase) [ToLower](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=305) 

``` go 
func (special SpecialCase) ToLower(r rune) rune
```

​	ToLower方法将符文映射为小写字母，优先考虑特殊映射。

#### (SpecialCase) [ToTitle](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=296) 

``` go 
func (special SpecialCase) ToTitle(r rune) rune
```

​	ToTitle方法将符文映射为标题大小写，优先考虑特殊映射。

#### (SpecialCase) [ToUpper](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/letter.go;l=287) 

``` go 
func (special SpecialCase) ToUpper(r rune) rune
```

​	ToUpper方法将符文映射为大写字母，优先考虑特殊映射。

## Notes

​	存在全大写或全小写等特殊映射，如土耳其语中的特殊映射。

## Bugs

- 不存在完全的大小写折叠机制，即涉及输入或输出中的多个符文的字符。