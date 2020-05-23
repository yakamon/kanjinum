package kanjinum

import (
	"math/big"
)

// NumToKanji converts int to string of kanji.
func NumToKanji(n *big.Int) string {
	D := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	S := []string{"", "十", "百", "千"}
	L := []string{"", "万", "億", "兆", "京", "垓", "𥝱", "穣", "溝", "澗", "正", "載", "極", "恒河沙", "阿僧祇", "那由他", "不可思議", "無量大数"}

	// 符号
	sign := "正ノ"
	if n.Sign() < 0 {
		sign = "負ノ"
	}
	n.Abs(n)

	zero, ten := big.NewInt(0), big.NewInt(10)
	// 零
	if n.Cmp(zero) == 0 {
		return D[0]
	}
	// 無量大数
	if n.Cmp(new(big.Int).Exp(ten, big.NewInt(68), nil)) >= 0 {
		return sign + L[len(L)-1]
	}

	kanjinum := ""
	part := ""
	si := 0
	li := 0
	for n.Cmp(zero) > 0 {
		remain := new(big.Int).Mod(n, ten).Int64()
		n.Div(n, ten)
		if remain > 0 {
			part = S[si] + part
			if si == 0 || remain > 1 {
				part = D[remain] + part
			}
		}
		si++

		if si == 4 {
			if len(part) > 0 {
				kanjinum = part + L[li] + kanjinum
			}
			part = ""
			si = 0
			li++
		}
	}
	if len(part) > 0 {
		kanjinum = part + L[li] + kanjinum
	}
	return sign + kanjinum
}
