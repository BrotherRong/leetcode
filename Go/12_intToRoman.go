package main

import (
	"fmt"
	"math"
	"strconv"
)

//罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。
//字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。
//通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
//给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。

func intToRoman(num int) string {
	M := []string{"", "M", "MM", "MMM"}  //1000，2000，3000
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"} // 100~900
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"} // 10~90
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"} // 1~9
	return M[num / 1000] + C[(num % 1000) / 100] + X[(num % 100) / 10] + I[num % 10]
}

func intToRoman1(num int) string {
	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	intSlice := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	s := ""
	x := len(intSlice) - 1
	for ; num != 0; {
		i := len(strconv.Itoa(num)) - 1
		n := (num / int(math.Pow10(i))) * int(math.Pow10(i))
		for ; x >= 0; x-- {
			if n >= intSlice[x] {
				s += roman[x]
				num -= intSlice[x]
				break
			}
		}
	}
	return s
}


func main() {
	fmt.Println(intToRoman(23))
}