package kanjinum

import "math/big"

// NumToKanji converts int to string of kanji.
func NumToKanji(n *big.Int) string {
	digits := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	su := []string{"一", "十", "百", "千"}
	lu := []string{"一", "万", "億", "兆", "京", "垓", "𥝱", "穣", "溝", "澗", "正", "載", "恒河沙", "阿僧祇", "那由他", "不可思議", "無量大数"}

	sign := "正ノ"
	if n.Sign() < 0 {
		sign = "負ノ"
	}

	zero, ten := big.NewInt(0), big.NewInt(10)
	if n.CmpAbs(new(big.Int).Exp(ten, big.NewInt(68), nil)) >= 0 {
		return sign + lu[len(lu)-1]
	}

	n.Abs(n)

	s := ""
	ndigit := 0

	r := new(big.Int).Mod(n, ten).Int64()
	n.Div(n, ten)
	if r >= 1 {
		s = digits[r] + s
	} else if n.Cmp(zero) == 0 {
		return digits[r]
	}
	ndigit++

	for n.Cmp(zero) > 0 {
		r = new(big.Int).Mod(n, ten).Int64()
		n.Div(n, ten)
		if ndigit%4 == 0 {
			s = lu[ndigit/4] + s
			if r >= 1 {
				s = digits[r] + s
			}
		} else if r > 0 {
			s = su[ndigit%4] + s
			if r >= 2 {
				s = digits[r] + s
			}
		}
		ndigit++
	}

	return sign + s
}
