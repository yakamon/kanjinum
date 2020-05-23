package kanjinum

import (
	"math/big"
	"testing"
)

func TestKanjiNumber(t *testing.T) {
	testcases := []string{
		"0",
		"1",
		"10",
		"-11",
		"20",
		"1234",
		"12345",
		"123456",
		"-1234567",
		"12345678",
		"100000000000000000000000000000000000000000000000000000000000000000000",
	}
	expected := []string{
		"零",
		"正ノ一",
		"正ノ十",
		"負ノ十一",
		"正ノ二十",
		"正ノ千二百三十四",
		"正ノ一万二千三百四十五",
		"正ノ十二万三千四百五十六",
		"負ノ百二十三万四千五百六十七",
		"正ノ千二百三十四万五千六百七十八",
		"正ノ無量大数",
	}
	for i, s := range testcases {
		n, _ := new(big.Int).SetString(s, 10)
		actual := NumToKanji(n)
		if actual != expected[i] {
			t.Errorf("KanjiNumber() is invalid: got %v, wanted %v", actual, expected[i])
		}
	}
}
