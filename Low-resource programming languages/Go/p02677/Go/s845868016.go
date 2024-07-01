package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

// 1行読み込み
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// 読み込みをint型へキャスト
func nextInt() int {
	sc.Scan()
	num, err := strconv.Atoi(sc.Text())
	if(err != nil) {
		panic(err)
	}
	return num
}

// 読み込みをfloat型へキャスト
func nextFloat() float64 {
	sc.Scan()
	num, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	// scannerの挙動を改行区切り → 空白区切りに変更
	sc.Split(bufio.ScanWords)
	a := nextFloat()
	b := nextFloat()
	h := nextFloat()
	m := nextFloat()

	a_angle := h * 30 + (m / 60) * 30
	b_angle := m * 6
	angle := a_angle - b_angle
	if angle < 0 {
		angle *= -1
	}
	if angle >= 180 {
		angle -= 180
	}
	fmt.Println(a, b, a_angle, b_angle, angle)
	c_pow := math.Pow(a, 2.0) + math.Pow(b, 2.0) - 2 * a * b * math.Cos(angle / 360 * 2 * math.Pi)
	fmt.Println(math.Sqrt(c_pow))
}