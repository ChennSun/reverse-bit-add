package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var cursor uint32 = 0
	for i := 0; i < 100; i++ {
		cursor, _ = RevBitAdd(cursor)
	}
}

func RevBitAdd(cursor uint32) (uint32, error) {
	var mask = uint32(math.Pow(float64(2), float64(18)) - 1)
	// uint32按位取反后高位会变成1， 将无用位置为0
	cursor |= (^mask & mask)
	revBitStr := reverse(fmt.Sprintf("%018b", cursor))
	i, err := strconv.ParseInt(revBitStr, 2, 32)
	if err != nil {
		return 0, err
	}
	i++
	revBitStr = reverse(fmt.Sprintf("%018b", i))
	fmt.Println(revBitStr)
	i, err = strconv.ParseInt(revBitStr, 2, 32)
	if err != nil {
		return 0, err
	}
	return uint32(i), nil
}

func reverse(str string) string {
	strByte := []byte(str)
	for i, j := 0, len(strByte)-1; i < j; i, j = i+1, j-1 {
		strByte[i], strByte[j] = strByte[j], strByte[i]
	}
	return string(strByte)
}
